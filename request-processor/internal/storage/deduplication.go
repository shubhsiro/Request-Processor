package storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisClient{client: client}
}

// IsUnique checks if an ID is unique
func IsUnique(redisClient *RedisClient, id string) (bool, error) {
	ttl := 60 * time.Second

	exists, err := redisClient.client.SetNX(ctx, id, "1", ttl).Result()
	if err != nil {
		return false, err
	}

	return exists, nil
}
