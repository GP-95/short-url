package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type cache struct {
	client *redis.Client
}

const cache_time = time.Minute * 30

var REDIS *cache

func New() (*cache, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	REDIS = &cache{client: r}

	err := REDIS.isAlive(context.Background())
	if err != nil {
		return nil, err
	}

	return REDIS, nil
}

func (c *cache) Close() error {
	err := c.client.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *cache) SaveCodeAndUrl(ctx context.Context, code string, url string) error {
	_, err := c.client.Set(ctx, code, url, cache_time).Result()
	// Ignore nil errors
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}

func (c *cache) GetUrlByCode(ctx context.Context, code string) (string, error) {
	res, err := c.client.Get(ctx, code).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}

	return res, nil
}

func (c *cache) isAlive(ctx context.Context) error {
	_, err := c.client.Ping(ctx).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}

func (c *cache) ResetCodeTTL(ctx context.Context, code string) error {
	_, err := c.client.Expire(ctx, code, cache_time).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}
