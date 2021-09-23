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
	// todo update stub.clients when kv changed
	clients []xclient.XClient
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

func (stub *RPCStub) Call(fcn string, arg proto.IRPCArg, reply proto.IRPCReply, checkf func(reply proto.IRPCReply) bool) bool {
	if arg == nil {
		reply.SetErrMsg(common.ErrNaNRPCArg.Error())
		return false
	}

	var errs []error = make([]error, 0)

	for _, client := range stub.clients {
		err := client.Call(context.Background(), fcn, arg, reply)
		if err != nil {
			errs = append(errs, fmt.Errorf("client Call got error: %s", err.Error()))
		}
	}

	if errs[0] != nil {
		reply.SetErrMsg(errs[0].Error())
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

	stub.clients = make([]xclient.XClient, 0, len(zkd.GetServices()))

	for _, service := range zkd.GetServices() {
		d, err := xclient.NewPeer2PeerDiscovery(service.Key, "")

		if err != nil {
			logger.Errorf("discover %s servers got error: %s", serverPath, err.Error())
		}

		client := xclient.NewXClient(serverPath, xclient.Failtry, xclient.RandomSelect, d, xclient.DefaultOption)
		stub.clients = append(stub.clients, client)
	}

	if stub == nil {
		return nil, errors.New("init rpc server error")
	}

	return stub, nil
}
