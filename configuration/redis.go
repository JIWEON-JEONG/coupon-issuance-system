package configuration

import (
	"github.com/go-redis/redis/v8"
)

func ConnectionRedis(config *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: "password",
		DB:       0,
	})
	return rdb
}
