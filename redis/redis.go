package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	Client *redis.Client
)

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "pass",
		DB:       0,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}