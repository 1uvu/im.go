package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"im/internal/pkg/rpc"
	"im/pkg/config"
	"im/pkg/proto"
)

type FormPeerChat struct {
	Msg       string `form:"msg" json:"msg" binding:"required"`
	ToUserID  string `form:"toUserID" json:"toUserID" binding:"required"`
	ToGroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

func PeerChat(c *gin.Context) {
	var form FormPeerChat

	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	toUserID, _ := strconv.Atoi(form.ToUserID)

	replyUserInfoQuery := new(proto.UserInfoQueryReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"UserInfoQuery",
		&proto.UserInfoQueryArg{
			UserID: toUserID,
		},
		replyUserInfoQuery,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.UserInfoQueryReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, NewResponse(replyUserInfoQuery.GetErrMsg(), nil))
		return
	}

	toUserName := replyUserInfoQuery.UserName

	replyAuthCheck := new(proto.AuthCheckReply)

	ok = rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"AuthCheck",
		&proto.AuthCheckArg{
			AuthToken: form.AuthToken,
		},
		replyAuthCheck,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.AuthCheckReply)
			return _reply.Code != proto.CodeFailedReply && _reply.UserID >= 0 && _reply.UserName != ""
		},
	)

	if !ok {
		Failed(c, NewResponse(replyAuthCheck.GetErrMsg(), nil))
		return
	}

	fromUserID := replyAuthCheck.UserID
	fromUserName := replyAuthCheck.UserName

	toGroupID, _ := strconv.Atoi(form.ToGroupID)

	replyPeerChat := new(proto.OpReply)

	ok = rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"PeerChat",
		&proto.OpArg{
			Msg:          form.Msg,
			FromUserId:   fromUserID,
			FromUserName: fromUserName,
			ToUserId:     toUserID,
			ToUserName:   toUserName,
			GroupId:      toGroupID,
			Op:           proto.OpPeerChat,
		},
		replyPeerChat,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.OpReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, NewResponse(replyPeerChat.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}

type FormGroupChat struct {
	Msg       string `form:"msg" json:"msg" binding:"required"`
	ToGroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

func GroupChat(c *gin.Context) {
	var form FormGroupChat

	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	replyAuthCheck := new(proto.AuthCheckReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"AuthCheck",
		&proto.AuthCheckArg{
			AuthToken: form.AuthToken,
		},
		replyAuthCheck,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.AuthCheckReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		ResponseWithCode(c, proto.CodeSessionError, NewResponse(replyAuthCheck.GetErrMsg(), nil))
		return
	}

	fromUserID := replyAuthCheck.UserID
	fromUserName := replyAuthCheck.UserName

	toGroupID, _ := strconv.Atoi(form.ToGroupID)

	replyGroupChat := new(proto.OpReply)

	ok = rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"GroupChat",
		&proto.OpArg{
			Msg:          form.Msg,
			FromUserId:   fromUserID,
			FromUserName: fromUserName,
			GroupId:      toGroupID,
			Op:           proto.OpGroupChat,
		},
		replyGroupChat,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.OpReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, NewResponse(replyGroupChat.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))

}

type FormGroupCount struct {
	GroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
}

func GroupCount(c *gin.Context) {
	var form FormGroupCount

	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	groupID, _ := strconv.Atoi(form.GroupID)

	reply := new(proto.OpReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"GroupCount",
		&proto.OpArg{
			GroupId: groupID,
			Op:      proto.OpGroupCount,
		},
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.OpReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, NewResponse(reply.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}

type FormGroupInfo struct {
	GroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
}

func GroupInfo(c *gin.Context) {
	var form FormGroupInfo

	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	groupID, _ := strconv.Atoi(form.GroupID)

	reply := new(proto.OpReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"GroupInfo",
		&proto.OpArg{
			GroupId: groupID,
			Op:      proto.OpGroupInfo,
		},
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.OpReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, NewResponse(reply.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}
