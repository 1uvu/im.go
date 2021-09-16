package connect

import (
	"errors"
	"sync"

	"im/pkg/proto"
)

type Group struct {
	GroupID int
	Count   uint64
	rwMux   sync.RWMutex
	live    bool
	head    *Dialog
}

func NewGroup(gid int) *Group {
	return &Group{
		GroupID: gid,
		live:    true,
	}
}

func (g *Group) Join(d *Dialog) error {
	g.rwMux.Lock()
	defer g.rwMux.Unlock()

	if g.live {
		if g.head != nil {
			g.head.Prev = d
		}

		d.Next = g.head
		d.Prev = nil
		g.head = d
		g.Count++
	} else {
		return errors.New("group is not live")
	}

	return nil
}

func (g *Group) Push(msg *proto.Msg) {
	g.rwMux.RLock()
	defer g.rwMux.RUnlock()

	for d := g.head; d != nil; d = d.Next {
		d.Push(msg)
	}
}

func (g *Group) Remove(d *Dialog) bool {
	g.rwMux.RLock()
	defer g.rwMux.RUnlock()

	if d.Next != nil {
		d.Next.Prev = d.Prev
	}

	if d.Prev != nil {
		d.Prev.Next = d.Next
	} else {
		g.head = d.Next
	}

	g.Count--
	g.live = true

	if g.Count <= 0 {
		g.live = false
	}

	return false
}
