package connect

import (
	"im/internal/pkg/logger"
	"im/internal/pkg/rpc"
	"im/pkg/proto"
)

type ConnectRPCServer struct {
	*rpc.DefaultRPCServer
}

func (server *ConnectRPCServer) Peer(arg *proto.ConnectPeerArg, reply *proto.ConnectReply) {
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

func (server *ConnectRPCServer) Group(arg *proto.ConnectGroupArg, reply *proto.ConnectReply) {
	reply.Code = proto.CodeSuccessReply

	for _, bucket := range DefaultServer.Buckets {
		bucket.Broadcast(arg)
	}
}
