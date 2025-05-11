package configuration

import (
	"github.com/go-redis/redis/v8"
)

func ConnectionRedis(config *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})
	return rdb
}
