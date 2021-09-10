package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"im/internal/pkg/rpc"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
)

type FormSignin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Signin(c *gin.Context) {
	var form FormSignin
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	reply := new(proto.SigninReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"Signin",
		&proto.SigninArg{
			UserName: form.Username,
			Password: common.SHA1(form.Password),
		},
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.SigninReply)
			return _reply.Code != proto.CodeFailed && _reply.AuthToken != ""
		},
	)

	if !ok {
		Failed(c, NewResponse(reply.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}

type FormSignup struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var form FormSignup
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	reply := new(proto.SignupReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"Signup",
		&proto.SignupArg{
			UserName: form.Username,
			Password: common.SHA1(form.Password),
		},
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.SignupReply)
			return _reply.Code != proto.CodeFailed && _reply.AuthToken != ""
		},
	)

	if !ok {
		Failed(c, NewResponse(reply.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}

type FormSignout struct {
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

func Signout(c *gin.Context) {
	var form FormSignout
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	reply := new(proto.SignoutReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"Signout",
		&proto.SignoutArg{
			AuthToken: form.AuthToken,
		},
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.SignoutReply)
			return _reply.Code != proto.CodeFailed
		},
	)

	if !ok {
		Failed(c, NewResponse(reply.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}

type FormAuthCheck struct {
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

func AuthCheck(c *gin.Context) {
	var form FormAuthCheck
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	reply := new(proto.AuthCheckReply)
	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"AuthCheck",
		&proto.AuthCheckArg{
			AuthToken: form.AuthToken,
		},
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.AuthCheckReply)
			return _reply.Code != proto.CodeFailed && _reply.UserID >= 0 && _reply.UserName != ""
		},
	)

	if !ok {
		ResponseWithCode(c, proto.CodeSessionError, NewResponse(reply.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}
