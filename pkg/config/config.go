package config

import "sync"

type Config struct {
	Common  CommonConfig
	API     APIConfig
	Connect ConnectConfig
	Logic   LogicConfig
	Task    TaskConfig
	Logger  LoggerConfig
	Website WebsiteConfig
}

var (
	conf *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		if conf == nil {
			initConf()
		}
	})

	return conf
}

func initConf() {
	// todo 1 从 yaml 读取配置写入 conf
}
