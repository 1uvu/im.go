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

// rpcc

func PeerPush(c *gin.Context) {
	var req proto.APIPeerPushRequest

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		Failed(c, "", nil)
		return
	}

	toUserID, _ := strconv.ParseUint(req.ToUserID, 0, 64)

	replyUserInfoQuery := new(proto.LogicUserInfoQueryReply)

	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)
	serverIDx := common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok := stub.Call(
		serverIDx,
		"UserInfoQuery",
		&proto.LogicUserInfoQueryArg{
			UserID: toUserID,
		},
		replyUserInfoQuery,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicUserInfoQueryReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, replyUserInfoQuery.GetErrMsg(), nil)
		return
	}

	toUserName := replyUserInfoQuery.UserName

	replyAuthCheck := new(proto.LogicAuthCheckReply)

	serverIDx = common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok = stub.Call(
		serverIDx,
		"AuthCheck",
		&proto.LogicAuthCheckArg{
			AuthToken: req.AuthToken,
		},
		replyAuthCheck,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicAuthCheckReply)
			return _reply.Code != proto.CodeFailedReply && _reply.UserID >= uint64(0) && _reply.UserName != ""
		},
	)

	if !ok {
		Failed(c, replyAuthCheck.GetErrMsg(), nil)
		return
	}

	fromUserID := replyAuthCheck.UserID
	fromUserName := replyAuthCheck.UserName

	toGroupID, _ := strconv.Atoi(req.ToGroupID)

	replyPeerPush := new(proto.LogicPeerPushReply)

	serverIDx = common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok = stub.Call(
		serverIDx,
		"PeerPush",
		&proto.LogicPeerPushArg{
			Msg:          req.Msg,
			FromUserId:   fromUserID,
			FromUserName: fromUserName,
			ToUserId:     toUserID,
			ToUserName:   toUserName,
			GroupId:      toGroupID,
			Op:           proto.OpPeerPush,
			Timestamp:    common.CreateTimestamp(),
		},
		replyPeerPush,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicPeerPushReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, replyPeerPush.GetErrMsg(), nil)
		return
	}

	Success(c, "ok", nil)
}

func GroupPush(c *gin.Context) {
	var req proto.APIGroupPushRequest

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		Failed(c, "", nil)
		return
	}

	replyAuthCheck := new(proto.LogicAuthCheckReply)

	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)
	serverIDx := common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok := stub.Call(
		serverIDx,
		"AuthCheck",
		&proto.LogicAuthCheckArg{
			AuthToken: req.AuthToken,
		},
		replyAuthCheck,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicAuthCheckReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		ResponseWithCode(c, proto.CodeSessionError, proto.APIResponse{
			Msg:  replyAuthCheck.GetErrMsg(),
			Data: nil,
		})
		return
	}

	fromUserID := replyAuthCheck.UserID
	fromUserName := replyAuthCheck.UserName
	toGroupID, _ := strconv.Atoi(req.ToGroupID)
	replyGroupPush := new(proto.LogicGroupPushReply)

	serverIDx = common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok = stub.Call(
		serverIDx,
		"GroupPush",
		&proto.LogicGroupPushArg{
			Msg:          req.Msg,
			FromUserId:   fromUserID,
			FromUserName: fromUserName,
			GroupId:      toGroupID,
			Op:           proto.OpGroupPush,
			Timestamp:    common.CreateTimestamp(),
		},
		replyGroupPush,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicGroupPushReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, replyGroupPush.GetErrMsg(), nil)
		return
	}

	Success(c, "ok", nil)

}

func GroupCount(c *gin.Context) {
	var req proto.APIGroupCountRequest

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		Failed(c, "", nil)
		return
	}

	groupID, _ := strconv.Atoi(req.GroupID)

	reply := new(proto.LogicGroupCountReply)

	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)
	serverIDx := common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok := stub.Call(
		serverIDx,
		"GroupCount",
		&proto.LogicGroupCountArg{
			GroupId:   groupID,
			Op:        proto.OpGroupCount,
			Timestamp: common.CreateTimestamp(),
		},
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicGroupCountReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, reply.GetErrMsg(), nil)
		return
	}

	Success(c, "ok", nil)
}

func GroupInfo(c *gin.Context) {
	var req proto.APIGroupInfoRequest

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		Failed(c, "", nil)
		return
	}

	groupID, _ := strconv.Atoi(req.GroupID)

	reply := new(proto.LogicGroupInfoReply)

	stub := rpc.GetStub(config.GetConfig().Common.ETCD.ServerPathLogic)
	serverIDx := common.GetServerIDx(config.GetConfig().Common.ETCD.ServerPathLogic, common.RandInt(stub.ClientNum))

	ok := stub.Call(
		serverIDx,
		"GroupInfo",
		&proto.LogicGroupInfoArg{
			GroupId:   groupID,
			Op:        proto.OpGroupInfo,
			Timestamp: common.CreateTimestamp(),
		},
		reply,
		func(reply proto.IRPCReply) bool {
			_reply := reply.(*proto.LogicGroupInfoReply)
			return _reply.Code != proto.CodeFailedReply
		},
	)

	if !ok {
		Failed(c, reply.GetErrMsg(), nil)
		return
	}

	Success(c, "ok", nil)
}
