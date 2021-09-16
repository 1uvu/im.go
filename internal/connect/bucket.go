package connect

import (
	"sync"
	"sync/atomic"

	"im/pkg/proto"
)

type Bucket struct {
	BucketID  uint32
	rwMux     sync.RWMutex
	option    BucketOption
	dialogs   map[uint64]*Dialog
	groups    map[int]*Group
	reqs      []chan *proto.GroupPushRequest
	reqsNum   uint64
	broadcast chan []byte
}

type BucketOption struct {
	DialogNum uint64
	GroupNum  int
	ArgAmount uint64
	ArgSize   uint64
}

func NewBucket(bid uint32, option BucketOption) *Bucket {
	b := new(Bucket)

	b.BucketID = bid
	b.option = option
	b.dialogs = make(map[uint64]*Dialog, option.DialogNum)
	b.groups = make(map[int]*Group, option.GroupNum)
	b.reqs = make([]chan *proto.GroupPushRequest, option.ArgAmount)
	b.broadcast = make(chan []byte)

	for i := range b.reqs {
		arg := make(chan *proto.GroupPushRequest, option.ArgSize)
		b.reqs[i] = arg

		go b.GroupPush(arg)
	}

	return b
}

func (b *Bucket) GroupPush(argCh chan *proto.GroupPushRequest) {
	arg := <-argCh
	if group, ok := b.GetGroup(arg.GroupID); ok {
		group.Push(&arg.Msg)
	}

}

func (b *Bucket) GetGroup(gid int) (*Group, bool) {
	b.rwMux.RLock()
	group, ok := b.groups[gid]
	b.rwMux.RUnlock()
	return group, ok
}

func (b *Bucket) PutUserIntoGroup(uid uint64, gid int, d *Dialog) error {
	b.rwMux.Lock()
	defer b.rwMux.Unlock()

	group, ok := b.groups[gid]
	if !ok {
		group = NewGroup(gid)
		b.groups[gid] = group
	}

	d.Group = group
	d.UserID = uid

	b.dialogs[uid] = d

	return group.Join(d)
}

func (b *Bucket) DeleteDialog(d *Dialog) bool {
	b.rwMux.RLock()
	defer b.rwMux.RUnlock()

	_, ok := b.dialogs[d.UserID]

	if ok {
		group, _ := b.GetGroup(d.Group.GroupID)
		group.Remove(d)

		delete(b.dialogs, d.UserID)
		if !group.live {
			delete(b.groups, group.GroupID)
		}
	}

	return ok
}

func (b *Bucket) GetDialog(uid uint64) (*Dialog, bool) {
	b.rwMux.RLock()
	dialog, ok := b.dialogs[uid]
	b.rwMux.RUnlock()
	return dialog, ok
}

func (b *Bucket) Broadcast(arg *proto.GroupPushRequest) {
	reqsNum := atomic.AddUint64(&b.reqsNum, 1) % b.option.ArgAmount
	b.reqs[reqsNum] <- arg
}
