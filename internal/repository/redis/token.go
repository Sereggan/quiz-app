package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type TokenClient struct {
	c *redis.Client
}

func (t TokenClient) Delete(context context.Context, key string) error {
	return t.c.Del(context, key).Err()
}

func (t TokenClient) Set(context context.Context, key string, value interface{}, ttl time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return t.c.Set(context, key, p, ttl).Err()
}

func (t TokenClient) Get(context context.Context, key string) (value interface{}, err error) {
	return t.c.Get(context, key).Result()
}

func NewTokenCLient(c *redis.Client) *TokenClient {
	return &TokenClient{c}
}
