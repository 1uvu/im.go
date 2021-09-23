package logic

import (
	"im/internal/pkg/logger"
	"im/internal/pkg/rpc"
	"im/pkg/config"
	"runtime"
)

type Logic struct {
}

func NewLogic() *Logic {
	return &Logic{}
}

func (logic *Logic) Run() {
	runtime.GOMAXPROCS(config.GetConfig().Logic.CPUs)

	if err := logic.InitPublishInstance(); err != nil {
		logger.Panicf("logic publish client initialize got error: %s", err.Error())
	}

	if err := rpc.RunRPCServer(config.GetConfig().Common.ETCD.ServerPathLogic, config.GetConfig().Logic.RPCAddress, new(LogicRPCServer)); err != nil {
		logger.Panicf("logic rpc server initialize got error: %s", err.Error())
	}
}
