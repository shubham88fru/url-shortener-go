package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)


var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNo,
	})
}