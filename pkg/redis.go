package pkg

import (
	"github.com/go-redis/redis/v8"
)

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
