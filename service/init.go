package service

import (
	"github.com/go-redis/redis"
	"github.com/terachanple/redis-ranking-demo/cache"
	"github.com/terachanple/redis-ranking-demo/config"
)

var redisClient *redis.Client

func Initialize(config config.Config) {
	redisClient = cache.NewClient(config)
}
