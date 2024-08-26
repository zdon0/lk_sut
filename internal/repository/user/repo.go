package user

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"lk_sut/internal/config"
	"lk_sut/internal/domain"
	"lk_sut/internal/domain/user"
)

type Repo struct {
	rdb                 *redis.Client
	userDataHTable      string
	userLastLoginHTable string
}

func NewRepo(rdb *redis.Client, cfg *config.Config) *Repo {
	return &Repo{
		rdb:                 rdb,
		userDataHTable:      cfg.Redis.UserDataHTable,
		userLastLoginHTable: cfg.Redis.UserLastLoginHTable,
	}
}

func (r *Repo) GetAllUsers(ctx context.Context) ([]user.User, error) {
	res, err := r.rdb.HGetAll(ctx, r.userDataHTable).Result()
	if err != nil {
		return nil, err
	}

	result := make([]user.User, 0, len(res))

	for key, value := range res {
		result = append(result, user.User{
			Login:    key,
			Password: value,
		})
	}

	return result, nil
}

func (r *Repo) GetUser(ctx context.Context, login string) (user.User, error) {
	res, err := r.rdb.HGet(ctx, r.userDataHTable, login).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return user.User{}, domain.ErrNotFound
		}

		return user.User{}, err
	}

	result := user.User{
		Login:    login,
		Password: res,
	}

	return result, nil
}

func (r *Repo) AddUser(ctx context.Context, user user.User) error {
	return r.rdb.HSet(ctx, r.userDataHTable, user.Login, user.Password).Err()
}

func (r *Repo) DeleteUser(ctx context.Context, user user.User) error {
	return r.rdb.HDel(ctx, r.userDataHTable, user.Login).Err()
}

func (r *Repo) GetUserLastLogin(ctx context.Context, login string) (time.Time, error) {
	res, err := r.rdb.HGet(ctx, r.userLastLoginHTable, login).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return time.Time{}, domain.ErrNotFound
		}

		return time.Time{}, err
	}

	return time.Unix(0, res), nil
}

func (r *Repo) SetUserLastLogin(ctx context.Context, login string, timeToSet time.Time) error {
	return r.rdb.HSet(ctx, r.userLastLoginHTable, login, timeToSet.UnixNano()).Err()
}

func (r *Repo) DeleteUserLastLogin(ctx context.Context, login string) error {
	return r.rdb.HDel(ctx, r.userLastLoginHTable, login).Err()
}

func (r *Repo) FlushLastLogin(ctx context.Context) error {
	return r.rdb.Del(ctx, r.userLastLoginHTable).Err()
}
