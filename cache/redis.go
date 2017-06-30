package cache

import (
	"github.com/go-redis/redis"
	"github.com/terachanple/redis-ranking-demo/config"
)

func NewClient(config config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
}
