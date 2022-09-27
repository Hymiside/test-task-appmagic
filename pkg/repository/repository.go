package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

type ConfigRepository struct {
	Host string
	Port string
}

type Repository struct {
	redis *redis.Client
}

func NewRepository(ctx context.Context, c ConfigRepository) (*Repository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
		Password: "",
		DB:       0,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		rdb.Close()
	}(ctx)

	return &Repository{redis: rdb}, nil
}
