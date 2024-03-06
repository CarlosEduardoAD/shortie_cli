package redis

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type ShortUrlRedis struct {
	redisClient *redis.Client
}

func NewShortUrlRedis(redisClient *redis.Client) *ShortUrlRedis {
	return &ShortUrlRedis{
		redisClient: redisClient,
	}
}

func (sr *ShortUrlRedis) CreateCache(code, originalUrl string) error {
	ctx := context.Background()
	ttl := 60 * time.Second

	status := sr.redisClient.Set(ctx, code, originalUrl, ttl)

	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func (sr *ShortUrlRedis) FindCache(code string) (string, error) {
	ctx := context.Background()

	originalUrl, err := sr.redisClient.Get(ctx, code).Result()

	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
