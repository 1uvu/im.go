package config

import "time"

type TaskConfig struct {
	CPUs          int           `yaml:"cpus"`
	PushPolling   time.Duration `yam:"pushPolling"`
	PushChanCap   int           `yaml:"pushChanCap"`
	PushParamsCap int           `yaml:"pushParamsCap"`
}
