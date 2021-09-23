package task

import (
	"im/internal/pkg/logger"
	"im/internal/pkg/mq"
	"im/pkg/config"
	"time"
)

var (
	queueInstance *mq.RedisInstance
)

func (task *Task) InitRedisQueueInstance() error {
	err := initRedisQueueInstance()

	go func() {
		for {
			result, err := queueInstance.Client.BRPop(config.GetConfig().Task.PushPolling*time.Second, config.GetConfig().Common.Redis.QueueName).Result()

			if err != nil {
				logger.Infof("task queue pop timeout, got error: %s", err.Error())
			}

			if len(result) >= 2 {
				task.Push(result[1])
			}
		}
	}()

	return err
}

func initRedisQueueInstance() error {
	option := mq.RedisOption{
		Address:    config.GetConfig().Common.Redis.Address,
		Password:   config.GetConfig().Common.Redis.Password,
		DBidx:      config.GetConfig().Common.Redis.DBidx,
		MaxConnAge: config.GetConfig().Common.Redis.MaxConnAge,
	}

	queueInstance = mq.GetRedisInstance(option)

	pong, err := queueInstance.Client.Ping().Result()

	if err != nil {
		logger.Errorf("redis queue instance ping result pong: %s, err: %s", pong, err.Error())
	}

	return nil
}
