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
	groups    map[uint64]*Group
	args      []chan *proto.GroupChatArg
	argsNum   uint64
	broadcast chan []byte
}

type BucketOption struct {
	DialogNum uint64
	GroupNum  uint64
	ArgAmount uint64
	ArgSize   uint64
}

func NewBucket(bid uint32, option BucketOption) *Bucket {
	b := new(Bucket)

	b.BucketID = bid
	b.option = option
	b.dialogs = make(map[uint64]*Dialog, option.DialogNum)
	b.groups = make(map[uint64]*Group, option.GroupNum)
	b.args = make([]chan *proto.GroupChatArg, option.ArgAmount)
	b.broadcast = make(chan []byte)

	for i := range b.args {
		arg := make(chan *proto.GroupChatArg, option.ArgSize)
		b.args[i] = arg

		go b.GroupChat(arg)
	}

	return b
}

func (b *Bucket) GroupChat(argCh chan *proto.GroupChatArg) {
	arg := <-argCh
	if group, ok := b.GetGroup(arg.GroupID); ok {
		group.Chat(&arg.Msg)
	}

}

func (b *Bucket) GetGroup(gid uint64) (*Group, bool) {
	b.rwMux.RLock()
	group, ok := b.groups[gid]
	b.rwMux.RUnlock()
	return group, ok
}

func (b *Bucket) PutUserIntoGroup(uid, gid uint64, d *Dialog) error {
	b.rwMux.Lock()
	defer b.rwMux.Unlock()

	group, ok := b.groups[gid]
	if !ok {
		group = NewGroup(gid)
	}

	d.Group = group
	d.UserID = uid

	b.groups[gid] = group
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

func (b *Bucket) Broadcast(arg *proto.GroupChatArg) {
	argsNum := atomic.AddUint64(&b.argsNum, 1) % b.option.ArgAmount
	b.args[argsNum] <- arg
}
