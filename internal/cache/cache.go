package cache

import (
	"context"
	"errors"
	"time"

	"github.com/bytedance/sonic"
	"github.com/nordew/go-errx"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func New(client *redis.Client) *Cache {
	return &Cache{client: client}
}

func (c *Cache) Set(
	ctx context.Context,
	key string,
	value any,
	expiration time.Duration,
) error {
	data, err := sonic.Marshal(value)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to marshal value for key", err)
	}

	if err := c.client.Set(
		ctx,
		key,
		data,
		expiration,
	).Err(); err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to set value for key", err)
	}

	return nil
}

func (c *Cache) Get(
	ctx context.Context,
	key string,
	dest any,
) error {
	data, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return errx.NewNotFound().WithDescription("cache: key not found")
		}

		return errx.NewInternal().WithDescriptionAndCause("failed to get value for key", err)
	}

	if err := sonic.Unmarshal([]byte(data), dest); err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to unmarshal value for key", err)
	}

	return nil
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	if err := c.client.Del(ctx, key).Err(); err != nil {
		return errx.NewInternal().WithDescriptionAndCause("failed to delete key", err)
	}

	return nil
}
