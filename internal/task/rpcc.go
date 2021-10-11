package task

import (
	"encoding/json"
	"im/internal/pkg/logger"
	"im/internal/pkg/rpc"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
)

func (task *Task) peerPush(pushParam *PushParam) {
	arg := &proto.ConnectPeerArg{
		UserID: pushParam.UserID,
		Msg: proto.Msg{
			Ver:       proto.GetVersion(),
			Operation: proto.OpGroupPush,
			SeqID:     common.GetSnowflakeID(1),
			Body:      pushParam.Msg,
		},
	}

	reply := new(proto.ConnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathConnect).Call(
		pushParam.ServerIDx,
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

func (task *Task) groupPush(param *proto.TaskGroupPushParam) {
	arg := &proto.ConnectGroupArg{
		GroupID: param.GroupID,
		Msg: proto.Msg{
			Ver:       proto.GetVersion(),
			Operation: param.Op,
			SeqID:     common.GetSnowflakeID(1),
			Body:      param.Msg,
		},
	}

	reply := new(proto.ConnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathConnect).CallAll(
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

func (task *Task) groupCount(param *proto.TaskGroupCountParam) {

	paramAsBytes, _ := json.Marshal(param)

	arg := &proto.ConnectGroupArg{
		GroupID: param.GroupID,
		Msg: proto.Msg{
			Ver:       proto.GetVersion(),
			Operation: param.Op,
			SeqID:     common.GetSnowflakeID(1),
			Body:      paramAsBytes,
		},
	}

	reply := new(proto.ConnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathConnect).CallAll(
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

func (task *Task) groupInfo(param *proto.TaskGroupInfoParam) {

	paramAsBytes, _ := json.Marshal(param)

	arg := &proto.ConnectGroupArg{
		GroupID: param.GroupID,
		Msg: proto.Msg{
			Ver:       proto.GetVersion(),
			Operation: param.Op,
			SeqID:     common.GetSnowflakeID(1),
			Body:      paramAsBytes,
		},
	}

	reply := new(proto.ConnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathConnect).CallAll(
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
