package rpc

import (
	"context"
	"errors"
	"fmt"
	"sync"

	xclient "github.com/smallnest/rpcx/client"

	"im/internal/pkg/logger"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
)

type RPCStub struct {
	// opt update stub.clients when kv changed
	Clients   map[string]*RPCClient
	ClientNum int
}

type RPCClient struct {
	Client xclient.XClient
}

var stubs map[string]*RPCStub
var rwMux sync.RWMutex

func GetStub(serverPath string) *RPCStub {
	rwMux.RLock()

	var (
		stub *RPCStub
		err  error
	)

	stub, ok := stubs[serverPath]

	if !ok {
		rwMux.Lock()

		stub, err = newStub(serverPath)

		if err != nil {
			logger.Panic(err)
		}

		stubs[serverPath] = stub

		rwMux.Unlock()
	}

	rwMux.RUnlock()

	return stub
}

func (stub *RPCStub) CallAll(fcn string, arg proto.IRPCArg, reply proto.IRPCReply, checkf func(reply proto.IRPCReply) bool) bool {
	var ok bool = true
	for serverIDx := range stub.Clients {

		ok = ok && stub.Call(
			serverIDx,
			fcn,
			arg,
			reply,
			checkf,
		)

	}

	return ok
}

func (stub *RPCStub) Call(serverIDx string, fcn string, arg proto.IRPCArg, reply proto.IRPCReply, checkf func(reply proto.IRPCReply) bool) bool {
	rpcclient, ok := stub.Clients[serverIDx]

	if !ok {
		reply.SetErrMsg(fmt.Sprintf("rpc call connect peer push got error: %s", common.ErrServerIDxNotExisted))
		return false
	}

	if arg == nil {
		reply.SetErrMsg(common.ErrNaNRPCArg.Error())
		return false
	}

	err := rpcclient.Client.Call(context.Background(), fcn, arg, reply)

	if err != nil {
		reply.SetErrMsg(fmt.Sprintf("client Call got error: %s", err.Error()))
	}

	return checkf(reply)
}

func newStub(serverPath string) (*RPCStub, error) {
	stub := new(RPCStub)

	zkd, err := xclient.NewZookeeperDiscovery(
		config.GetConfig().Common.ETCD.BasePath,
		serverPath,
		[]string{config.GetConfig().Common.ETCD.Host},
		nil,
	)

	if err != nil {
		logger.Fatal(err.Error())
	}

	stub.ClientNum = len(zkd.GetServices())
	stub.Clients = make(map[string]*RPCClient, stub.ClientNum)

	for _, service := range zkd.GetServices() {
		d, err := xclient.NewPeer2PeerDiscovery(service.Key, "")

		if err != nil {
			logger.Errorf("discover %s servers got error: %s", serverPath, err.Error())
		}

		client := xclient.NewXClient(serverPath, xclient.Failtry, xclient.RandomSelect, d, xclient.DefaultOption)
		stub.Clients[service.Value] = &RPCClient{Client: client}
	}

	if stub == nil {
		return nil, errors.New("init rpc server error")
	}

	return stub, nil
}
