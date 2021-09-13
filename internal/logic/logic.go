package logic

import (
	"fmt"
	"im/internal/pkg/logger"
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

	if err := logic.RunRedisPublishClient(); err != nil {
		logger.Panicf("logic redis publish client initialize got error: %s", err.Error())
	}

	if err := logic.RunRPCClient(); err != nil {
		logger.Panicf("logic rpc client initialize got error: %s", err.Error())
	}
}
