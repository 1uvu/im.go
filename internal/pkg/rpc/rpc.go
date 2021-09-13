package rpc

import (
	"context"
	"errors"
	"sync"

	xclient "github.com/smallnest/rpcx/client"

	"im/internal/pkg/logger"
	"im/pkg/config"
	"im/pkg/proto"
)

type Stub struct {
	client xclient.XClient
}

var stubs map[string]*Stub
var rwMux sync.RWMutex

func GetStub(serverPath string) *Stub {
	rwMux.RLock()

	var (
		stub *Stub
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

func (stub *Stub) Call(fcn string, arg proto.ILogicArg, reply proto.ILogicReply, checkf func(reply proto.ILogicReply) bool) bool {

	err := stub.client.Call(context.Background(), fcn, arg, reply)

	if err != nil {
		reply.SetErrMsg(err.Error())
	}

	return checkf(reply)
}

func newStub(serverPath string) (*Stub, error) {
	stub := new(Stub)

	zkd, err := xclient.NewZookeeperDiscovery(
		config.GetConfig().Common.ETCD.BasePath,
		serverPath,
		[]string{config.GetConfig().Common.ETCD.Host},
		nil,
	)

	if err != nil {
		logger.Fatal(err.Error())
	}

	stub.client = xclient.NewXClient(serverPath, xclient.Failtry, xclient.RandomSelect, zkd, xclient.DefaultOption)

	if stub == nil {
		return nil, errors.New("init rpc stub error")
	}

	return stub, nil
}
