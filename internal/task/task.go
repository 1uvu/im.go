package task

import (
	"im/internal/pkg/logger"
	"im/pkg/config"
	"runtime"
)

type Task struct {
}

func NewTask() *Task {
	return &Task{}
}

func (task *Task) Run() {
	runtime.GOMAXPROCS(config.GetConfig().Task.CPUs)

	if err := task.InitRedisQueueInstance(); err != nil {
		logger.Panicf("task queue client initialize got error: %s", err.Error())
	}

	task.DoPush()
}
