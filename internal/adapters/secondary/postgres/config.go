package postgres

import (
	"payment-gw/pkg/conf"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	PoolSize int
}

func LoadConfig() *Config {
	return &Config{
		Host:     conf.GetString("database", "host"),
		Port:     conf.GetInt("database", "port"),
		User:     conf.GetString("database", "user"),
		Password: conf.GetString("database", "password"),
		Database: conf.GetString("database", "database"),
		PoolSize: conf.GetIntDefault("database", "pool_size", 10),
	}
}
