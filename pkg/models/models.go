package models

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/arpithku/URL_Shortener_1/pkg/config"
	"github.com/redis/go-redis/v9"
)

type RequestResponseBody struct {
	URL           string `json:"url"`
	Shortened_URL string `json:"shortened_url"`
}

var client_redis *redis.Client

const stringset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = 62 //As there are only a total of 62 characters in the above set
func init() {
	config.Connect()
	client_redis = config.GetRedisClient()
	err := client_redis.Set(context.Background(), "uuid_count", 0, 0).Err() //Setting the incremental id static intially
	if err != nil {
		panic(err)
	}
}

func SetURL(ctx context.Context, key string) string {

	uuid_count, err := client_redis.Get(ctx, "uuid_count").Result()
	if err != nil {
		fmt.Println("Error while retreiving static string")
		panic(err)
	}
	if uuid_count == "" {
		panic("uuid_count LOST")
	}
	uuid_count_int, err := strconv.Atoi(uuid_count)
	if err != nil {
		panic(err)
	}
	uuid_val := uuid_count_int + 1 //Value incremented to store later in Redis
	uuid_count_int = uuid_count_int + 1
	var uuid string
	for uuid_count_int > 0 {
		uuid = uuid + string(stringset[uuid_count_int%base])
		uuid_count_int = uuid_count_int / base
	}
	err = client_redis.MSet(ctx, "uuid_count", uuid_val, uuid, key).Err()
	if err != nil {
		panic(err)
	}
	//fmt.Println(uuid)

	return uuid
}

func GetURL(ctx context.Context, key string) string {
	val, err := client_redis.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	//fmt.Println("--VALUE--", val)
	return val
}

func DeleteURL(ctx context.Context, key string) error {
	val, err := client_redis.Get(ctx, key).Result()
	if val == "" {
		fmt.Println("URL does not exist")
		return errors.New("url does not exist")
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	client_redis.Del(ctx, key)
	//fmt.Println(output)
	return nil
}

func UpdateURL(ctx context.Context, key string, value string) error {
	val, err := client_redis.Get(ctx, key).Result()
	if val == "" {
		fmt.Println("URL does not exist")
		return errors.New("url does not exist")
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
