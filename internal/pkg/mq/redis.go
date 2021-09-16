package mq

import (
	"encoding/json"
	"im/internal/pkg/logger"
	"im/pkg/config"
	"im/pkg/proto"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type RedisInstance struct {
	*DefaultInstance

	Client *redis.Client
}

type RedisOption struct {
	*DefaultOption

	Address    string
	Password   string
	DBidx      int
	MaxConnAge time.Duration
}

var (
	redisInstances map[string]*RedisInstance
	rwMux          sync.RWMutex
)

func GetRedisInstance(option RedisOption) *RedisInstance {
	rwMux.RLock()
	var instance *RedisInstance

	instance, ok := redisInstances[option.Address]

	if !ok {
		rwMux.Lock()
		instance = &RedisInstance{
			Client: redis.NewClient(&redis.Options{
				Addr:       option.Address,
				Password:   option.Password,
				DB:         option.DBidx,
				MaxConnAge: option.MaxConnAge * time.Second,
			}),
		}

		redisInstances[option.Address] = instance
		rwMux.Unlock()
	}
	rwMux.RUnlock()

	return instance
}

func (instance *RedisInstance) Push(publishArg proto.PublishArg) error {
	publishArgAsBytes, err := json.Marshal(publishArg)

	if err != nil {
		logger.Errorf("logic publish peer marshal got error: %s", err.Error())
		return err
	}

	publishQueueName := config.GetConfig().Common.Redis.QueueName

	if err := instance.Client.LPush(publishQueueName, publishArgAsBytes).Err(); err != nil {
		logger.Errorf("logic publish redis lpush got error: %s", err.Error())
		return err
	}

	return nil
}
