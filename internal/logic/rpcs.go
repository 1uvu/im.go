package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"im/internal/logic/dao"
	"im/internal/pkg/logger"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
	"strconv"
	"time"
)

type Stub struct{}

func (stub *Stub) Signup(ctx context.Context, arg *proto.LogicSignupArg, reply *proto.LogicSignupReply) error {
	reply.Code = proto.CodeFailedReply
	userModel := &dao.UserModel{
		UserID:         1, // todo userID
		UserName:       arg.UserName,
		SaltedPassword: common.SaltPassword(arg.Password),
	}

	userID, err := dao.Create(userModel)

	if err != nil {
		return err
	}

	userModel.UserID = userID

	// set session
	randToken := common.CreateToken(32)
	sessionID := common.CreateSessionIDByToken(randToken)
	session := map[string]interface{}{
		"userID":   userID,
		"userName": arg.UserName,
	}

	sessionInstance.Client.Do("MULTI")
	sessionInstance.Client.HMSet(sessionID, session)
	sessionInstance.Client.Expire(sessionID, config.GetConfig().Connect.SessionExpireTime*time.Second)

	if err = sessionInstance.Client.Do("EXEC").Err(); err != nil {
		logger.Infof("register session set failed")
		return err
	}

	reply.Code = proto.CodeSuccessReply
	reply.AuthToken = randToken

	return nil
}

func (stub *Stub) Signin(ctx context.Context, arg *proto.LogicSigninArg, reply *proto.LogicSigninReply) error {
	reply.Code = proto.CodeFailedReply

	userModel, err := dao.ReadByName(arg.UserName)

	if err != nil {
		return err
	}

	if userModel.SaltedPassword != common.SaltPassword(arg.Password) {
		return common.ErrUnmatchedPassword
	}

	// check if signin session has existed?
	signinSessionID := common.CreateSessionIDByUserID(userModel.UserID)

	_token, _ := sessionInstance.Client.Get(signinSessionID).Result()

	if _token != "" {
		// token has exist, so signout firstly
		oldSession := common.CreateSessionIDByToken(_token)

		if err := sessionInstance.Client.Del(oldSession).Err(); err != nil {
			return common.ErrUserSignoutFailed
		}
	}

	// and then update session
	randToken := common.CreateToken(32)
	sessionID := common.CreateSessionIDByToken(randToken)
	session := map[string]interface{}{
		"userID":   userModel.UserID,
		"userName": userModel.UserName,
	}

	sessionInstance.Client.Do("MULTI")
	sessionInstance.Client.HMSet(sessionID, session)
	sessionInstance.Client.Expire(sessionID, config.GetConfig().Connect.SessionExpireTime*time.Second)

	if err = sessionInstance.Client.Do("EXEC").Err(); err != nil {
		logger.Infof("register session set failed")
		return err
	}

	reply.Code = proto.CodeSuccessReply
	reply.AuthToken = randToken

	return nil
}

func (stub *Stub) Signout(ctx context.Context, arg *proto.LogicSignoutArg, reply *proto.LogicSignoutReply) error {
	reply.Code = proto.CodeFailedReply
	sessionID := common.GetSessionIDByToken(arg.AuthToken)

	// get session
	session, err := sessionInstance.Client.HGetAll(sessionID).Result()

	if err != nil {
		return common.ErrUnmatchedAuthToken
	}

	if len(session) == 0 {
		return common.ErrSessionHasExpired
	}

	userID, _ := strconv.ParseUint(session["userID"], 0, 64)

	// del session from signin
	signinSessionID := common.CreateSessionIDByUserID(userID)

	if err := sessionInstance.Client.Del(signinSessionID).Err(); err != nil {
		return common.ErrSessionDeletFailed
	}

	// del serverID about this user
	logic := NewLogic()
	userServerIDKey := logic.getKey(config.GetConfig().Common.Redis.Prefix, fmt.Sprintf("%d", userID))

	if err := sessionInstance.Client.Del(userServerIDKey).Err(); err != nil {
		return common.ErrUserServerQuitFailed
	}

	if err := sessionInstance.Client.Del(sessionID).Err(); err != nil {
		return common.ErrUserSignoutFailed
	}

	reply.Code = proto.CodeSuccessReply

	return nil
}

func (stub *Stub) AuthCheck(ctx context.Context, arg *proto.LogicAuthCheckArg, reply *proto.LogicAuthCheckReply) error {
	reply.Code = proto.CodeFailedReply
	sessionID := common.GetSessionIDByToken(arg.AuthToken)

	// get session
	session, err := sessionInstance.Client.HGetAll(sessionID).Result()

	if err != nil {
		return common.ErrUnmatchedAuthToken
	}

	if len(session) == 0 {
		return common.ErrSessionHasExpired
	}

	userID, _ := strconv.ParseUint(session["userID"], 0, 64)
	userName := session["userName"]

	reply.UserID = userID
	reply.UserName = userName
	reply.Code = proto.CodeSuccessReply

	return nil
}

func (stub *Stub) UserInfoQuery(ctx context.Context, arg *proto.LogicUserInfoQueryArg, reply *proto.LogicUserInfoQueryReply) error {
	reply.Code = proto.CodeFailedReply

	userModel, err := dao.Read(arg.UserID)

	if err != nil {
		return common.ErrUserNotExisted
	}

	reply.UserID = arg.UserID
	reply.UserName = userModel.UserName
	reply.Code = proto.CodeSuccessReply

	return nil
}

func (stub *Stub) PeerPush(ctx context.Context, arg *proto.LogicPeerPushArg, reply *proto.LogicPeerPushReply) error {
	reply.Code = proto.CodeFailedReply

	argAsBytes, err := json.Marshal(arg)

	if err != nil {
		return common.ErrMarshalPushArgFailed
	}
	logic := NewLogic()

	userServerIDKey := logic.getKey(config.GetConfig().Common.Redis.Prefix, fmt.Sprintf("%d", arg.ToUserId))
	serverID := sessionInstance.Client.Get(userServerIDKey).Val()

	err = logic.Publish(proto.TaskPeerPushParam{
		Op:       proto.OpPeerPush,
		ServerID: serverID,
		UserID:   arg.ToUserId,
		Msg:      argAsBytes,
	})

	if err != nil {
		return common.ErrPublishFailed
	}

	reply.Code = proto.CodeSuccessReply

	return nil
}

func (stub *Stub) GroupPush(ctx context.Context, arg *proto.LogicGroupPushArg, reply *proto.LogicGroupPushReply) error {
	reply.Code = proto.CodeFailedReply

	argAsBytes, err := json.Marshal(arg)

	if err != nil {
		return common.ErrMarshalPushArgFailed
	}
	logic := NewLogic()

	groupUsersKey := logic.getKey(config.GetConfig().Common.Redis.GroupPrefix, fmt.Sprintf("%d", arg.GroupId))
	groupUserInfos, err := publishInstance.Client.HGetAll(groupUsersKey).Result()

	if err != nil {
		// todo: 在类似的 err return 前加上 logger
		return common.ErrGetGroupUsersFailed
	}

	if len(groupUserInfos) == 0 {
		return common.ErrGroupIsNotLive
	}

	err = logic.Publish(proto.TaskGroupPushParam{
		Op:             proto.OpGroupPush,
		GroupID:        arg.GroupId,
		Count:          len(groupUserInfos),
		Msg:            argAsBytes,
		GroupUserInfos: groupUserInfos,
	})

	if err != nil {
		return common.ErrPublishFailed
	}

	reply.Code = proto.CodeSuccessReply

	return nil
}

func (stub *Stub) GroupCount(ctx context.Context, arg *proto.LogicGroupCountArg, reply *proto.LogicGroupCountReply) error {
	reply.Code = proto.CodeFailedReply

	logic := NewLogic()

	groupCountKey := logic.getKey(config.GetConfig().Common.Redis.GroupCountPrefix, fmt.Sprintf("%d", arg.GroupId))
	groupCount, err := sessionInstance.Client.Get(groupCountKey).Int()

	if err != nil {
		return common.ErrGetGroupCountFailed
	}

	err = logic.Publish(proto.TaskGroupCountParam{
		Op:      proto.OpGroupCount,
		GroupID: arg.GroupId,
		Count:   groupCount,
	})

	if err != nil {
		return common.ErrPublishFailed
	}

	reply.Code = proto.CodeSuccessReply

	return nil
}

func (stub *Stub) GroupInfo(ctx context.Context, arg *proto.LogicGroupInfoArg, reply *proto.LogicGroupInfoReply) error {
	reply.Code = proto.CodeFailedReply

	logic := NewLogic()

	groupUsersKey := logic.getKey(config.GetConfig().Common.Redis.GroupPrefix, fmt.Sprintf("%d", arg.GroupId))
	groupUserInfos, err := publishInstance.Client.HGetAll(groupUsersKey).Result()

	if err != nil {
		// todo: 在类似的 err return 前加上 logger
		return common.ErrGetGroupUsersFailed
	}

	if len(groupUserInfos) == 0 {
		return common.ErrGroupIsNotLive
	}

	err = logic.Publish(proto.TaskGroupInfoParam{
		Op:             proto.OpGroupInfo,
		GroupID:        arg.GroupId,
		Count:          len(groupUserInfos),
		GroupUserInfos: groupUserInfos,
	})

	if err != nil {
		return common.ErrPublishFailed
	}

	reply.Code = proto.CodeSuccessReply

	return nil
}

func (stub *Stub) Connect(ctx context.Context, arg *proto.LogicConnectArg, reply *proto.LogicConnectReply) error {
	reply.Code = proto.CodeFailedReply

	sessionID := common.CreateSessionIDByToken(arg.AuthToken)

	session, err := sessionInstance.Client.HGetAll(sessionID).Result()

	if err != nil {
		return common.ErrUnmatchedAuthToken
	}

	if len(session) == 0 {
		reply.UserID = 0
		return common.ErrSessionHasExpired
	}

	userID, _ := strconv.ParseUint(session["userID"], 0, 64)
	reply.UserID = userID

	logic := NewLogic()
	groupUsersKey := logic.getKey(config.GetConfig().Common.Redis.GroupPrefix, fmt.Sprintf("%d", arg.GroupID))

	if reply.UserID != 0 {
		userServerIDKey := logic.getKey(config.GetConfig().Common.Redis.Prefix, fmt.Sprintf("%d", reply.UserID))
		validTime := config.GetConfig().Common.Redis.BaseValidTime * time.Second
		err := publishInstance.Client.Set(userServerIDKey, arg.ServerID, validTime)

		if err != nil {
			return common.ErrConnectFailed
		}

		if publishInstance.Client.HGet(groupUsersKey, fmt.Sprintf("%d", reply.UserID)).Val() == "" {
			publishInstance.Client.HSet(groupUsersKey, fmt.Sprintf("%d", reply.UserID), session["userName"])
			publishInstance.Client.Incr(logic.getKey(config.GetConfig().Common.Redis.GroupCountPrefix, fmt.Sprintf("%d", arg.GroupID)))
		}
	}

	reply.Code = proto.CodeSuccessReply

	return nil
}

func (stub *Stub) Disconnect(ctx context.Context, arg *proto.LogicDisconnectArg, reply *proto.LogicDisconnectReply) error {
	reply.Code = proto.CodeFailedReply

	logic := NewLogic()
	groupUsersKey := logic.getKey(config.GetConfig().Common.Redis.GroupPrefix, fmt.Sprintf("%d", arg.GroupID))

	if arg.GroupID > 0 {
		groupCount, _ := sessionInstance.Client.Get(logic.getKey(config.GetConfig().Common.Redis.GroupCountPrefix, fmt.Sprintf("%d", arg.GroupID))).Int()

		if groupCount > 0 {
			publishInstance.Client.Decr(logic.getKey(config.GetConfig().Common.Redis.GroupCountPrefix, fmt.Sprintf("%d", arg.GroupID))).Result()
		}
	}

	if arg.UserID > 0 {
		if err := publishInstance.Client.HDel(groupUsersKey, fmt.Sprintf("%d", arg.UserID)).Err(); err != nil {
			return common.ErrDisconnectFailed
		}
	}

	groupUserInfos, err := publishInstance.Client.HGetAll(groupUsersKey).Result()

	if err != nil {
		return common.ErrDisconnectFailed
	}

	if err := logic.Publish(proto.TaskGroupPushParam{
		Op:             proto.OpGroupPush,
		GroupID:        arg.GroupID,
		Count:          len(groupUserInfos),
		Msg:            nil,
		GroupUserInfos: groupUserInfos,
	}); err != nil {
		return common.ErrPublishFailed
	}

	return nil
}
