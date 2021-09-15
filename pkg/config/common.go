package config

import "time"

type CommonConfig struct {
	ETCD  CommonETCD  `yaml:"etcd"`
	Redis CommonRedis `yaml:"redis"`
}

type CommonETCD struct {
	BasePath          string `yaml:"basePath"`
	Host              string `yaml:"host"`
	ServerPathConnect string `yaml:"serverPathConnect"`
	ServerPathLogic   string `yaml:"serverPathLogic"`
}

type CommonRedis struct {
	QueueName       string        `yaml:"queueName"`
	BaseValidTime   int           `yaml:"baseValidTime"`
	Prefix          string        `yaml:"prefix"`
	GroupPrefix     string        `yaml:"groupPrefix"`
	GroupLivePrefix string        `yaml:"groupLivePrefix"`
	Address         string        `yaml:"address"`
	Password        string        `yaml:"password"`
	DBidx           int           `yaml:"dbidx"`
	MaxConnAge      time.Duration `yaml:"maxConnAge"`
}
