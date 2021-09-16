package connect

import (
	"net"

	"github.com/gorilla/websocket"

	"im/pkg/proto"
)

type Dialog struct {
	UserID    uint64
	Group     *Group
	Next      *Dialog
	Prev      *Dialog
	broadcast chan *proto.Msg
	connTCP   *net.TCPConn
	connWS    *websocket.Conn
}

func NewDialog(uid uint64, size int) *Dialog {
	return &Dialog{
		UserID:    uid,
		broadcast: make(chan *proto.Msg, size),
	}
}

func (d *Dialog) Push(msg *proto.Msg) error {
	select {
	case d.broadcast <- msg:
	default:
	}

	return nil
}
