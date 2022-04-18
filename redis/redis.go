package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type RedisContext struct {
	RedisClient *redis.Client
}

func ConnectRedis(config *RedisConfig) (redisContext *RedisContext, err string) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		// DB:       config.DB,
	})

	ctx := context.TODO()
	_, error := redisClient.Ping(ctx).Result()
	if error != nil {
		log.Printf("Failed to connect to Redis")
		return nil, "Failed to connect to Redis"
	}
	return &RedisContext{redisClient}, ""
}
