package config

import "time"

type LogicConfig struct {
	ServerID   string    `yaml:"serverID"`
	CPUs       int       `yaml:"cpus"`
	RPCAddress string    `yaml:"rpcaddress"`
	Auth       LogicAuth `yaml:"auth"`
	DB         LogicDB   `yaml:"db"`
}

type LogicAuth struct {
	CertPath string `yaml:"certPath"`
	KeyPath  string `yaml:"keyPath"`
}

type LogicDB struct {
	DBName string   `yaml:"dbName"`
	Sqlite SqliteDB `yaml:"sqlite"`
}

type SqliteDB struct {
	DBPath          string        `yaml:"dbPath"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
	ConnMaxIdletime time.Duration `yaml:"connMaxLifetime"`
}
