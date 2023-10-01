package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"lk_sut/internal/config"
)

func NewClient(cfg config.Redis) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    cfg.Password,
		DB:          cfg.DB,
		DialTimeout: cfg.Timeout,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
