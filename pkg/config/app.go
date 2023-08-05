package config

import (
	"github.com/redis/go-redis/v9"
)

var client_redis *redis.Client

func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	client_redis = client

}

func GetRedisClient() *redis.Client {
	return client_redis
}
