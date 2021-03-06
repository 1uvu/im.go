package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"im/internal/pkg/rpc"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
)

// rpcc

func Signup(c *gin.Context) {
	var req proto.APISignupRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		Failed(c, "", nil)
		return
	}

	reply := new(proto.LogicSignupReply)

	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)
	serverIDx := common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok := stub.Call(
		serverIDx,
		"Signup",
		&proto.LogicSignupArg{
			UserName: req.Username,
			Password: common.SHA1(req.Password),
		},
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicSignupReply)
			return _reply.Code != proto.CodeFailedReply && _reply.AuthToken != ""
		},
	)

	if !ok {
		Failed(c, reply.GetErrMsg(), nil)
		return
	}

	Success(c, "ok", nil)
}

func Signin(c *gin.Context) {
	var req proto.APISigninRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		Failed(c, "", nil)
		return
	}

	reply := new(proto.LogicSigninReply)

	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)
	serverIDx := common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok := stub.Call(
		serverIDx,
		"Signin",
		&proto.LogicSigninArg{
			UserName: req.Username,
			Password: common.SHA1(req.Password),
		},
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicSigninReply)
			return _reply.Code != proto.CodeFailedReply && _reply.AuthToken != ""
		},
	)

	if !ok {
		Failed(c, reply.GetErrMsg(), nil)
		return
	}

	Success(c, "ok", nil)
}

func Signout(c *gin.Context) {
	var req proto.APISignoutRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		Failed(c, "", nil)
		return
	}

	reply := new(proto.LogicSignoutReply)

	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)
	serverIDx := common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok := stub.Call(
		serverIDx,
		"Signout",
		&proto.LogicSignoutArg{
			AuthToken: req.AuthToken,
		},
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicSignoutReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, reply.GetErrMsg(), nil)
		return
	}

	Success(c, "ok", nil)
}

func AuthCheck(c *gin.Context) {
	var req proto.APIAuthCheckRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		Failed(c, "", nil)
		return
	}

	reply := new(proto.LogicAuthCheckReply)

	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)
	serverIDx := common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok := stub.Call(
		serverIDx,
		"AuthCheck",
		&proto.LogicAuthCheckArg{
			AuthToken: req.AuthToken,
		},
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicAuthCheckReply)
			return _reply.Code != proto.CodeFailedReply && _reply.UserID > 0 && _reply.UserName != ""
		},
	)

	if !ok {
		ResponseWithCode(c, proto.CodeSessionError, proto.APIResponse{
			Msg:  reply.GetErrMsg(),
			Data: nil,
		})
		return
	}

	Success(c, "ok", nil)
}
