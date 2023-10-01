package env

import "lk_sut/internal/api/handler/user"

type Env struct {
	User *user.Handler
}

func NewEnv(user *user.Handler) *Env {
	return &Env{
		User: user,
	}
}
