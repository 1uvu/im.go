package config

type APIConfig struct {
	RunMode    string `yaml:"runMode"`
	ListenPort string `yaml:"listenPort"`
	CORSFlag   bool   `yaml:"corsFlag"`
}
