package config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisContext = context.Background()
var RedisClient *redis.Client

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6397",
		Password: "",
		DB:       0,
	})
	return client
}
