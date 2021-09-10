package config

type LogicConfig struct {
	ServerID string    `yaml:"serverID"`
	CPUs     int       `yaml:"cpus"`
	Address  string    `yaml:"address"`
	Auth     LogicAuth `yaml:"auth"`
}

type LogicAuth struct {
	CertPath string `yaml:"certPath"`
	KeyPath  string `yaml:"keyPath"`
}
