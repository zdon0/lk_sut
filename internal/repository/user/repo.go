package user

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"lk_sut/internal/config"
	"lk_sut/internal/domain"
	"time"
)

const (
	timeLayout = time.RFC3339
)

type Repo struct {
	rdb                 *redis.Client
	userDataHTable      string
	userLastLoginHTable string
}

func NewRepo(rdb *redis.Client, cfg config.Redis) *Repo {
	return &Repo{
		rdb:                 rdb,
		userDataHTable:      cfg.UserDataHTable,
		userLastLoginHTable: cfg.UserLastLoginHTable,
	}
}

func (r *Repo) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	res, err := r.rdb.HGetAll(ctx, r.userDataHTable).Result()
	if err != nil {
		return nil, err
	}

	result := make([]domain.User, 0, len(res))

	for key, value := range res {
		user := domain.User{
			Login:    key,
			Password: value,
		}

		result = append(result, user)
	}

	return result, nil
}

func (r *Repo) GetUser(ctx context.Context, login string) (domain.User, error) {
	res, err := r.rdb.HGet(ctx, r.userDataHTable, login).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return domain.User{}, domain.ErrNotFound
		}

		return domain.User{}, err
	}

	result := domain.User{
		Login:    login,
		Password: res,
	}

	return result, nil
}

func (r *Repo) AddUser(ctx context.Context, user domain.User) error {
	return r.rdb.HSet(ctx, r.userDataHTable, user.Login, user.Password).Err()
}

func (r *Repo) DeleteUser(ctx context.Context, user domain.User) error {
	return r.rdb.HDel(ctx, r.userDataHTable, user.Login).Err()
}

func (r *Repo) GetUserLastLogin(ctx context.Context, login string) (time.Time, error) {
	res, err := r.rdb.HGet(ctx, r.userLastLoginHTable, login).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return time.Time{}, domain.ErrNotFound
		}

		return time.Time{}, err
	}

	parsedTime, err := time.Parse(timeLayout, res)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func (r *Repo) SetUserLastLogin(ctx context.Context, login string, timeToSet time.Time) error {
	return r.rdb.HSet(ctx, r.userLastLoginHTable, login, timeToSet.Format(timeLayout)).Err()
}

func (r *Repo) DeleteUserLastLogin(ctx context.Context, login string) error {
	return r.rdb.HDel(ctx, r.userLastLoginHTable, login).Err()
}

func (r *Repo) FlushLastLogin(ctx context.Context) error {
	return r.rdb.Del(ctx, r.userLastLoginHTable).Err()
}
