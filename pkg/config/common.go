package config

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
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
