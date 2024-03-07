package utils

import (
	"fmt"
	"os"

	redis "github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client {
	err := LoadEnvWithPath()

	if err != nil {
		os.Exit(1)
	}

	password := os.Getenv("REDIS_PASSWORD")

	opt, _ := redis.ParseURL(fmt.Sprintf("rediss://default:%s@us1-apt-marten-41949.upstash.io:41949", password))
	client := redis.NewClient(opt)

	return client
}
