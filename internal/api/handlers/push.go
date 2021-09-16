package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"im/internal/pkg/rpc"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
)

type FormPeerPush struct {
	Msg       string `form:"msg" json:"msg" binding:"required"`
	ToUserID  string `form:"toUserID" json:"toUserID" binding:"required"`
	ToGroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

func PeerPush(c *gin.Context) {
	var form FormPeerPush

	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		Failed(c, Response{})
		return
	}

	toUserID, _ := strconv.ParseUint(form.ToUserID, 0, 64)

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
			return _reply.Code != proto.CodeFailedReply && _reply.UserID >= uint64(0) && _reply.UserName != ""
		},
	)

	if !ok {
		Failed(c, NewResponse(replyAuthCheck.GetErrMsg(), nil))
		return
	}

	fromUserID := replyAuthCheck.UserID
	fromUserName := replyAuthCheck.UserName

	toGroupID, _ := strconv.Atoi(form.ToGroupID)

	replyPeerPush := new(proto.PushReply)

	ok = rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"PeerPush",
		&proto.PushArg{
			Msg:          form.Msg,
			FromUserId:   fromUserID,
			FromUserName: fromUserName,
			ToUserId:     toUserID,
			ToUserName:   toUserName,
			GroupId:      toGroupID,
			Op:           proto.OpPeerPush,
			Timestamp:    common.CreateTimestamp(),
		},
		replyPeerPush,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.PushReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, NewResponse(replyPeerPush.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}

type FormGroupPush struct {
	Msg       string `form:"msg" json:"msg" binding:"required"`
	ToGroupID string `form:"toGroupID" json:"toGroupID" binding:"required"`
	AuthToken string `form:"authToken" json:"authToken" binding:"required"`
}

func GroupPush(c *gin.Context) {
	var form FormGroupPush

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

	replyGroupPush := new(proto.PushReply)

	ok = rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"GroupPush",
		&proto.PushArg{
			Msg:          form.Msg,
			FromUserId:   fromUserID,
			FromUserName: fromUserName,
			GroupId:      toGroupID,
			Op:           proto.OpGroupPush,
			Timestamp:    common.CreateTimestamp(),
		},
		replyGroupPush,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.PushReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, NewResponse(replyGroupPush.GetErrMsg(), nil))
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

	reply := new(proto.PushReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"GroupCount",
		&proto.PushArg{
			GroupId:   groupID,
			Op:        proto.OpGroupCount,
			Timestamp: common.CreateTimestamp(),
		},
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.PushReply)
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

	reply := new(proto.PushReply)

	ok := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic).Call(
		"GroupInfo",
		&proto.PushArg{
			GroupId:   groupID,
			Op:        proto.OpGroupInfo,
			Timestamp: common.CreateTimestamp(),
		},
		reply,
		func(reply proto.ILogicReply) bool {
			_reply := reply.(*proto.PushReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, NewResponse(reply.GetErrMsg(), nil))
		return
	}

	Success(c, NewResponse("ok", nil))
}
