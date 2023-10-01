package domain

import "context"

type User struct {
	Login    string
	Password string
}

type UpdateUser struct {
	Login       string
	OldPassword string
	NewPassword string
}

type Repository interface {
	GetUser(ctx context.Context, login string) (User, error)
	AddUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, user User) error
	DeleteUserLastLogin(ctx context.Context, login string) error
}

type Authorization interface {
	AuthorizeUser(ctx context.Context, user User) error
}
