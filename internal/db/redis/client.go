package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"

	"lk_sut/internal/config"
)

func NewClient(cfg *config.Config, lc fx.Lifecycle) *redis.Client {
	redisCfg := cfg.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", redisCfg.Addr, redisCfg.Port),
		Password:    redisCfg.Password,
		DB:          redisCfg.DB,
		DialTimeout: redisCfg.Timeout,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return rdb.Ping(ctx).Err()
		},
		OnStop: func(_ context.Context) error {
			return rdb.Close()
		},
	})

	return rdb
}
