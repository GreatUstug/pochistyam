package cache

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		return nil
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return RedisClient.Ping(context.Background()).Err()
}
