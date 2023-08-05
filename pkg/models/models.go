package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/arpithku/URL_Shortener_1/pkg/config"
	"github.com/redis/go-redis/v9"
)

var client_redis *redis.Client

func init() {
	config.Connect()
	client_redis = config.GetRedisClient()
}

func SetURL(ctx context.Context, key string, value string) (string, string) {
	err := client_redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	return key, value
}

func GetURL(ctx context.Context, key string) string {
	val, err := client_redis.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("--VALUE--", val)
	return val
}

func DeleteURL(ctx context.Context, key string) {
	val := client_redis.Del(ctx, key)
	fmt.Println(val)
}

func UpdateURL(ctx context.Context, key string, value string) error {
	val, err := client_redis.Get(ctx, key).Result()
	if val == "" {
		fmt.Println("Key does not exist")
		return errors.New("Key does not exist")
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = client_redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
