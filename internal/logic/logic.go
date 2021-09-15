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

	if err := logic.RunPublishInstance(); err != nil {
		logger.Panicf("logic redis publish client initialize got error: %s", err.Error())
	}

	if err := logic.RunRPCServer(); err != nil {
		logger.Panicf("logic rpc server initialize got error: %s", err.Error())
	}
}
