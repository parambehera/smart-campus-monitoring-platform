package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisClient *redis.Client

func ConnectRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_HOST") + ":" + GetEnv("REDIS_PORT"),
	})

	_, err := RedisClient.Ping(Ctx).Result()

	if err != nil {
		panic(err)
	}

	println("✅ Connected to Redis")
}