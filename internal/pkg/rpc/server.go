package rpc

import (
	"fmt"
	"im/internal/pkg/logger"
	"im/pkg/common"
	"im/pkg/config"
	"strings"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

type IRPCServer interface {
	MustEmbedDefaultRPCServer()
}

type DefaultRPCServer struct {
}

func (s *DefaultRPCServer) MustEmbedDefaultRPCServer() {}

func RunRPCServer(serverPath, address string, rcvr IRPCServer) error {

	rpcAddressList := strings.Split(address, ",")

	// opt update stub.clients when servers changed
	for idx, bind := range rpcAddressList {
		if network, addr, err := common.ParseNetworkAddr(bind); err != nil {
			logger.Panicf("init rpc server got error: %s", err.Error())
			return err
		} else {
			logger.Infof("rpc server start at: %s:%s", network, addr)

			// server idx format with: serverID-idx
			go createRPCServer(serverPath, fmt.Sprintf("%s-%d", serverPath, idx), network, addr, rcvr)
		}
	}

	return nil
}

func createRPCServer(serverPath, serverIDx, network, addr string, rcvr IRPCServer) {
	s := server.NewServer()

	addRegisterPlugin(s, network, addr)

	s.RegisterName(serverPath, rcvr, serverIDx)

	s.RegisterOnShutdown(func(s *server.Server) {
		s.UnregisterAll()
	})

	s.Serve(network, addr)
}

func addRegisterPlugin(s *server.Server, network, addr string) {
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
