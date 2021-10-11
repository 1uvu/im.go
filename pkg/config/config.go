package config

import (
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var configName = [7]string{
	"common",
	"logger",
	"website",
	"api",
	"connect",
	"logic",
	"task",
}

type Config struct {
	Common  CommonConfig
	Logger  LoggerConfig
	Website WebsiteConfig
	API     APIConfig
	Connect ConnectConfig
	Logic   LogicConfig
	Task    TaskConfig
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
	absPath, _ := filepath.Abs("./")
	configPath := filepath.Join(
		absPath,
		"configs",
	)

	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.SetConfigName("/common")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	for i := 1; i < len(configName); i++ {
		viper.SetConfigName(configName[i])

		err = viper.MergeInConfig()

		if err != nil {
			panic(err)
		}
	}

	conf = new(Config)

	viper.Unmarshal(&conf.Common)
	viper.Unmarshal(&conf.Logger)
	viper.Unmarshal(&conf.Website)
	viper.Unmarshal(&conf.API)
	viper.Unmarshal(&conf.Connect)
	viper.Unmarshal(&conf.Logic)
	viper.Unmarshal(&conf.Task)
}
