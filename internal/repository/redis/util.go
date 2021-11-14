package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func GetClient(address string) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr: address,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
