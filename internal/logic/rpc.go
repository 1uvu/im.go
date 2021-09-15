package logic

import (
	"context"
	"im/internal/logic/dao"
	"im/internal/pkg/logger"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
	"strings"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

type Stub struct{}

// func (logic *Logic) RunRPCClient() error {

// 	return nil
// }

func (logic *Logic) RunRPCServer() error {
	rpcAddressList := strings.Split(config.GetConfig().Logic.RPCAddress, ",")

	for _, bind := range rpcAddressList {
		if network, addr, err := common.ParseNetworkAddr(bind); err != nil {
			logger.Panicf("init logic rpc server got error: %s", err.Error())
			return err
		} else {
			logger.Infof("logic rpc server start at: %s:%s", network, addr)

			go logic.createRPCServer(network, addr)
		}
	}

	return nil
}

func (stub *Stub) Signup(ctx context.Context, arg *proto.SignupArg, reply *proto.SignupReply) error {
	reply.Code = proto.CodeFailedReply
	userModel := &dao.UserModel{
		UserID:         1,
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

	publishSessionInstance.Client.Do("MULTI")
	publishSessionInstance.Client.HMSet(sessionID, session)
	publishSessionInstance.Client.Expire(sessionID, config.GetConfig().Connect.SessionExpireTime*time.Second)

	if err = publishSessionInstance.Client.Do("EXEC").Err(); err != nil {
		logger.Infof("register session set failed")
		return err
	}

	reply.Code = proto.CodeSuccessReply
	reply.AuthToken = randToken

	return nil
}

func (stub *Stub) Signin(ctx context.Context, arg *proto.SigninArg, reply *proto.SigninReply) error {
	reply.Code = proto.CodeFailedReply

	userModel, err := dao.ReadByName(arg.UserName)

	if err != nil {
		return err
	}

	if userModel.SaltedPassword != common.SaltPassword(arg.Password) {
		return common.UnmatchedPasswordError
	}

	// check if session has existed?
	signinSessionID := common.CreateSessionIDByUserID(userModel.UserID)

	_token, _ := publishSessionInstance.Client.Get(signinSessionID).Result()

	if _token != "" {
		// token has exist, so signout firstly
		oldSession := common.CreateSessionIDByToken(_token)

		if err := publishSessionInstance.Client.Del(oldSession).Err(); err != nil {
			return common.SignoutFailedError
		}
	}

	// and then update session
	randToken := common.CreateToken(32)
	sessionID := common.CreateSessionIDByToken(randToken)
	session := map[string]interface{}{
		"userID":   userModel.UserID,
		"userName": userModel.UserName,
	}

	publishSessionInstance.Client.Do("MULTI")
	publishSessionInstance.Client.HMSet(sessionID, session)
	publishSessionInstance.Client.Expire(sessionID, config.GetConfig().Connect.SessionExpireTime*time.Second)

	if err = publishSessionInstance.Client.Do("EXEC").Err(); err != nil {
		logger.Infof("register session set failed")
		return err
	}

	reply.Code = proto.CodeSuccessReply
	reply.AuthToken = randToken

	return nil
}

func (stub *Stub) Signout(ctx context.Context, arg *proto.SignoutArg, reply *proto.SignoutReply) error {
	return nil
}

func (stub *Stub) AuthCheck(ctx context.Context, arg *proto.AuthCheckArg, reply *proto.AuthCheckReply) error {
	return nil
}

func (stub *Stub) UserInfoQuery(ctx context.Context, arg *proto.UserInfoQueryArg, reply *proto.UserInfoQueryReply) error {
	return nil
}

func (stub *Stub) PeerChat(ctx context.Context, arg *proto.OpArg, reply *proto.OpReply) error {
	return nil
}

func (stub *Stub) GroupChat(ctx context.Context, arg *proto.OpArg, reply *proto.OpReply) error {
	return nil
}

func (stub *Stub) GroupCount(ctx context.Context, arg *proto.OpArg, reply *proto.OpReply) error {
	return nil
}

func (stub *Stub) GroupInfo(ctx context.Context, arg *proto.OpArg, reply *proto.OpReply) error {
	return nil
}

func (stub *Stub) Connect(ctx context.Context, arg *proto.ConnectArg, reply *proto.ConnectReply) error {
	return nil
}

func (stub *Stub) Disconnect(ctx context.Context, arg *proto.DisconnectArg, reply *proto.DisconnectReply) error {
	return nil
}

func (logic *Logic) createRPCServer(network, addr string) {
	s := server.NewServer()
	logic.addRegisterPlugin(s, network, addr)
}

func (logic *Logic) addRegisterPlugin(s *server.Server, network, addr string) {
	p := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   strings.Join([]string{network, addr}, common.NetworkSplitSign),
		ZooKeeperServers: []string{config.GetConfig().Common.ETCD.Host},
		BasePath:         config.GetConfig().Common.ETCD.BasePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}

	if err := p.Start(); err != nil {
		logger.Fatal(err)
	}

	s.Plugins.Add(p)
}
