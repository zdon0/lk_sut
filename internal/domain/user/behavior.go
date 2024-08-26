package user

import (
	"context"
	"errors"

	"lk_sut/internal/sutclient"
)

type Behavior struct {
	repo Repository
	auth Authorization
}

func NewBehavior(repo Repository, auth Authorization) *Behavior {
	return &Behavior{
		repo: repo,
		auth: auth,
	}
}

func (b *Behavior) Repo() Repository {
	return b.repo
}

func (b *Behavior) Authorize(ctx context.Context, user User) error {
	err := b.auth.AuthorizeUser(ctx, user.Login, user.Password)
	if errors.Is(err, sutclient.ErrBadUser) {
		err = ErrBadUser
	}

	return err
}
