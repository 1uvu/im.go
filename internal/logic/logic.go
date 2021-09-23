package logic

import (
	"fmt"
	"im/internal/pkg/logger"
	"im/internal/pkg/rpc"
	"im/pkg/config"
	"runtime"

	"github.com/google/uuid"
)

type Logic struct {
	ServerID string
}

func NewLogic() *Logic {
	return &Logic{}
}

func (logic *Logic) Run() {
	runtime.GOMAXPROCS(config.GetConfig().Logic.CPUs)
	logic.ServerID = fmt.Sprintf("logic-%s", uuid.New().String())

	if err := logic.InitPublishInstance(); err != nil {
		logger.Panicf("logic publish client initialize got error: %s", err.Error())
	}

	rpcServer := rpc.RPCServer{}

	if err := rpcServer.Run(config.GetConfig().Logic.RPCAddress); err != nil {
		logger.Panicf("logic rpc server initialize got error: %s", err.Error())
	}
}
