package config

type TaskConfig struct {
	CPUs  int       `yaml:"cpus"`
	Redis TaskRedis `yaml:"redis"`
}

type TaskRedis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}
