package rpc

import (
	"im/internal/pkg/logger"
	"im/pkg/common"
	"im/pkg/config"
	"strings"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

type RPCServer struct {
}

func (s *RPCServer) Run(address string) error {

	rpcAddressList := strings.Split(address, ",")

	for _, bind := range rpcAddressList {
		if network, addr, err := common.ParseNetworkAddr(bind); err != nil {
			logger.Panicf("init rpc server got error: %s", err.Error())
			return err
		} else {
			logger.Infof("rpc server start at: %s:%s", network, addr)

			go createRPCServer(network, addr)
		}
	}

	return nil
}

func createRPCServer(network, addr string) {
	s := server.NewServer()
	addRegisterPlugin(s, network, addr)
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
