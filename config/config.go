package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort    int    `envconfig:"server_port" default:"8080"`
	RedisAddr     string `envconfig:"redis_addr" default:"localhost:6379"`
	RedisPassword string `envconfig:"redis_password" default:""`
	RedisDB       int    `envconfig:"redis_db" default:"0"`
}

func New(prefix string) Config {
	config := Config{}
	envconfig.MustProcess(prefix, &config)

	return config
}
