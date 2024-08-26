package worker

import (
	"context"
	"time"

	"lk_sut/internal/domain/user"
)

type LkSutCommitter interface {
	AuthorizeUser(ctx context.Context, login, password string) error
	CommitLesson(ctx context.Context) error
}

type Repository interface {
	GetAllUsers(ctx context.Context) ([]user.User, error)
	GetUserLastLogin(ctx context.Context, login string) (time.Time, error)
	SetUserLastLogin(ctx context.Context, login string, timeToSet time.Time) error
	FlushLastLogin(ctx context.Context) error
}
