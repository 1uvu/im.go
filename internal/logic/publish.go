package logic

import (
	"bytes"
	"im/internal/pkg/logger"
	"im/internal/pkg/mq"
	"im/pkg/config"
	"im/pkg/proto"
)

var (
	publishInstance        *mq.RedisInstance
	publishSessionInstance *mq.RedisInstance
)

func (logic *Logic) RunPublishInstance() error {
	err := logic.runRedisPublishInstance()

	return err
}

func (logic *Logic) Publish(publishArg proto.PublishArg) error {

	err := publishInstance.Push(publishArg)

	return err
}

func (logic *Logic) Puah(publishArg proto.PublishArg) error {

	err := publishInstance.Push(publishArg)

	return err
}

func (logic *Logic) runRedisPublishInstance() error {
	option := mq.RedisOption{
		Address:    config.GetConfig().Common.Redis.Address,
		Password:   config.GetConfig().Common.Redis.Password,
		DBidx:      config.GetConfig().Common.Redis.DBidx,
		MaxConnAge: config.GetConfig().Common.Redis.MaxConnAge,
	}

	publishInstance = mq.GetRedisInstance(option)

	pong, err := publishInstance.Client.Ping().Result()

	if err != nil {
		logger.Errorf("redis instance ping result pong: %s, err: %s", pong, err.Error())
	}

	publishSessionInstance = publishInstance

	return err
}

func (logic *Logic) getKey(prefix string, authKey string) string {
	var key bytes.Buffer
	key.WriteString(prefix)
	key.WriteString(authKey)
	return key.String()
}
