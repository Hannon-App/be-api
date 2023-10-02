package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func InitRedis(cfg *AppConfig) *redis.Client {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.IP_Public_Redis,
		Password: cfg.Pass_Redis,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}
