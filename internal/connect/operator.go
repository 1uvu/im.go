package connect

import (
	"errors"
	"fmt"

	"im/internal/pkg/rpc"
	"im/pkg/config"
	"im/pkg/proto"
)

type Operator interface {
	Connnect(*proto.LogicConnectArg) (*proto.LogicConnectReply, error)
	Disconnect(*proto.LogicDisconnectArg) (*proto.LogicDisconnectReply, error)
}
type DefaultOperator struct {
}

// rpcc

func (op *DefaultOperator) Connnect(arg *proto.LogicConnectArg) (*proto.LogicConnectReply, error) {

	if arg.AuthToken == "" {
		return nil, errors.New(("auth token of conn req is nil"))
	}

	reply := new(proto.LogicConnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"Connect",
		arg,
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicConnectReply)
			return _reply.Code != proto.CodeFailedReply && _reply.UserID != 0
		},
	)

	if !ok {
		return nil, fmt.Errorf("auth token of conn req is invalid with error: %s", reply.GetErrMsg())
	}

	return reply, nil
}

func (op *DefaultOperator) Disconnect(arg *proto.LogicDisconnectArg) (*proto.LogicDisconnectReply, error) {

	reply := new(proto.LogicDisconnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"Disconnect",
		arg,
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicDisconnectReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		return nil, fmt.Errorf("disconn req get error: %s", reply.GetErrMsg())
	}

	return reply, nil
}
