package connect

import (
	"errors"
	"fmt"

	"im/internal/pkg/rpc"
	"im/pkg/config"
	"im/pkg/proto"
)

type Operator interface {
	Connnect(*proto.ConnectArg) (*proto.ConnectReply, error)
	Disconnect(*proto.DisconnectArg) (*proto.DisconnectReply, error)
}
type DefaultOperator struct {
}

func (op *DefaultOperator) Connnect(arg *proto.ConnectArg) (*proto.ConnectReply, error) {

	if arg.AuthToken == "" {
		return nil, errors.New(("auth token of conn req is nil"))
	}

	reply := new(proto.ConnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"Connect",
		arg,
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.ConnectReply)
			return _reply.Code != proto.CodeFailedReply && _reply.UserID != 0
		},
	)

	if !ok {
		return nil, fmt.Errorf("auth token of conn req is invalid with error: %s", reply.GetErrMsg())
	}

	return reply, nil
}

func (op *DefaultOperator) Disconnect(arg *proto.DisconnectArg) (*proto.DisconnectReply, error) {

	reply := new(proto.DisconnectReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"Disconnect",
		arg,
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.DisconnectReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		return nil, fmt.Errorf("disconn req get error: %s", reply.GetErrMsg())
	}

	return reply, nil
}
