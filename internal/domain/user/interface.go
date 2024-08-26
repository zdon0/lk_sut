package user

import "context"

type Repository interface {
	GetUser(ctx context.Context, login string) (User, error)
	AddUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, user User) error
	DeleteUserLastLogin(ctx context.Context, login string) error
}

type Authorization interface {
	AuthorizeUser(ctx context.Context, login, password string) error
}
