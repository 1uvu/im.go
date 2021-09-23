package task

import (
	"im/internal/pkg/logger"
	"im/internal/pkg/rpc"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
)

func (task *Task) peerPush(pushArg *PushArg) {
	arg := &proto.ConnectPeerArg{
		UserID: pushArg.UserID,
		Msg: proto.Msg{
			Ver:       proto.GetVersion(),
			Operation: proto.OpGroupPush,
			SeqID:     common.GetSnowflakeID(1),
			Body:      pushArg.Msg,
		},
	}

	reply := new(proto.ConnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathConnect).Call(
		"Peer",
		arg,
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.ConnectReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		logger.Errorf("rpc call connect peer push got error: %s", reply.GetErrMsg())
	}

}

func (task *Task) groupPush(pushArg *PushArg) {
	arg := &proto.ConnectGroupArg{
		GroupID: pushArg.GroupID,
		Count:   pushArg.Count,
		Msg: proto.Msg{
			Ver:       proto.GetVersion(),
			Operation: proto.OpGroupPush,
			SeqID:     common.GetSnowflakeID(1),
			Body:      pushArg.Msg,
		},
	}

	reply := new(proto.ConnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathConnect).Call(
		"Group",
		arg,
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.ConnectReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		logger.Errorf("rpc call connect group push got error: %s", reply.GetErrMsg())
	}
}
