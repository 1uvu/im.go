package connect

import (
	"im/internal/pkg/logger"
	"im/pkg/proto"
)

type Stub struct {
}

func (stub *Stub) Peer(arg *proto.ConnectPeerArg, reply *proto.ConnectReply) {
	reply.Code = proto.CodeFailedReply

	bucket := DefaultServer.GetBucket(arg.UserID)
	dialog, ok := bucket.GetDialog(arg.UserID)

	if !ok {
		err := dialog.Push(&arg.Msg)
		logger.Error(err)
		reply.SetErrMsg(err.Error())
		return
	}

	reply.Code = proto.CodeSuccessReply
}

func (stub *Stub) Group(arg *proto.ConnectGroupArg, reply *proto.ConnectReply) {
	reply.Code = proto.CodeSuccessReply

	for _, bucket := range DefaultServer.Buckets {
		bucket.Broadcast(arg)
	}
}
