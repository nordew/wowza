package dragonfly

import (
	"context"
	"fmt"
	"log"
	"wowza/internal/config"

	"github.com/redis/go-redis/v9"
)

func DragonflyConnect(ctx context.Context, cfg config.Dragonfly) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping dragonfly: %w", err)
	}

	log.Println("Successfully connected to DragonflyDB")

	return rdb, nil
}
