package config

import "time"

type LogicConfig struct {
	ServerID string    `yaml:"serverID"`
	CPUs     int       `yaml:"cpus"`
	Address  string    `yaml:"address"`
	Auth     LogicAuth `yaml:"auth"`
	DB       LogicDB   `yaml:"db"`
}

type LogicAuth struct {
	CertPath string `yaml:"certPath"`
	KeyPath  string `yaml:"keyPath"`
}

type LogicDB struct {
	DBName string   `yaml:"dbName"`
	Sqlite SqliteDB `yaml:"sqlite"`
	Redis  RedisDB  `yaml:"redis"`
}

type SqliteDB struct {
	DBPath          string        `yaml:"dbPath"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
	ConnMaxIdletime time.Duration `yaml:"connMaxLifetime"`
}

type RedisDB struct {
}
