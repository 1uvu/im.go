package connect

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	"im/internal/pkg/logger"
	"im/internal/pkg/rpc"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
)

type Server struct {
	ServerIDx string
	Buckets   []*Bucket
	Option    ServerOption
	bucketNum uint32
	operator  Operator
}

type ServerOption struct {
	WriteWait      time.Duration
	PongWait       time.Duration
	PingPeriod     time.Duration
	MaxMessageSize uint64
	RBufferSize    int
	WBufferSize    int
	BroadcastSize  int
}

func NewServer(bs []*Bucket, op Operator, option ServerOption) *Server {
	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)

	return &Server{
		// server 与 rpc server 相对应
		// 则此处在原本的 rpc server idx 基础上自增 1
		ServerIDx: common.NewServerIDx(config.GetConfig().Common.ETCD.ServerPathConnect, stub.ClientNum),
		Buckets:   bs,
		Option:    option,
		bucketNum: uint32(len(bs)),
		operator:  op,
	}
}

func (s *Server) GetBucket(uid uint64) *Bucket {
	uidStr := fmt.Sprintf("%d", uid)
	bid := common.CityHash32([]byte(uidStr), uint32(len(uidStr))) % s.bucketNum

	return s.Buckets[bid]
}

func (s *Server) writePump(d *Dialog, c *Connect) {
	ticker := time.NewTicker(s.Option.PingPeriod)
	defer func() {
		ticker.Stop()
		d.connWS.Close()
	}()

	for {
		select {
		case msg, ok := <-d.broadcast:
			d.connWS.SetWriteDeadline(time.Now().Add(s.Option.WriteWait))
			if !ok {
				logger.Warn("set write deadline for server failed")
				d.connWS.WriteMessage(websocket.CloseMessage, []byte{})
			}

			w, err := d.connWS.NextWriter(websocket.TextMessage)

			if err != nil {
				logger.Warn("get dialog.connWS.NextWriter got error: %s", err.Error())
				return
			}

			w.Write(msg.Body)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			d.connWS.SetWriteDeadline(time.Now().Add(s.Option.WriteWait))

			if err := d.connWS.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (s *Server) readPump(d *Dialog, c *Connect) {
	defer func() {
		if d.Group == nil || d.UserID == 0 {
			d.connWS.Close()
			return
		}

		disconnArg := &proto.LogicDisconnectArg{
			GroupID: d.Group.GroupID,
			UserID:  d.UserID,
		}

		s.GetBucket(d.UserID).DeleteDialog(d)

		if disconnResp, err := s.operator.Disconnect(disconnArg); err != nil {
			logger.Warn("disconnect got error: %s, with code: %s", err.Error(), disconnResp.Code)
		}

		d.connWS.Close()
	}()

	d.connWS.SetReadLimit(int64(s.Option.MaxMessageSize))
	d.connWS.SetReadDeadline(time.Now().Add(s.Option.PongWait))
	d.connWS.SetPongHandler(func(appData string) error {
		d.connWS.SetReadDeadline(time.Now().Add(s.Option.PongWait))

		return nil
	})

	for {
		_, msg, err := d.connWS.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				logger.Errorf("read pump by calling dialog.connWS.ReadMessage got error: %s", err.Error())
				return
			}
		}

		connReq, err := readConnReq(msg)

		if err != nil {
			logger.Errorf("read conn request got error: %s", err.Error())
			return
		}

		connReq.ServerIDx = s.ServerIDx
		connResp, err := s.operator.Connnect(connReq)

		if err != nil {
			logger.Errorf("connect got error: %s", err.Error())
			return
		}

		bucket := s.GetBucket(connResp.UserID)

		err = bucket.PutUserIntoGroup(connResp.UserID, connReq.GroupID, d)

		if err != nil {
			logger.Errorf("user join into group got error: %s", err.Error())
			d.connWS.Close()
		}

	}

}

func readConnReq(msg []byte) (*proto.LogicConnectArg, error) {
	if msg == nil {
		return nil, errors.New("msg body for conn req is nil")
	}

	connReq := new(proto.LogicConnectArg)

	if err := json.Unmarshal([]byte(msg), &connReq); err != nil {
		return nil, err
	}

	return connReq, nil
}
