package user

import (
	"context"
	"errors"
	userdomain "lk_sut/internal/domain"
	"lk_sut/internal/sutclient"
)

type Interactor struct {
	repo userdomain.Repository
	auth userdomain.Authorization
}

func NewInteractor(repo userdomain.Repository, auth userdomain.Authorization) *Interactor {
	return &Interactor{
		repo: repo,
		auth: auth,
	}
}

func (i *Interactor) GetUserByLogin(ctx context.Context, login string) (userdomain.User, error) {
	return i.repo.GetUser(ctx, login)
}

func (i *Interactor) AddUser(ctx context.Context, user userdomain.User) error {
	if err := i.validateUser(user); err != nil {
		return err
	}

	_, err := i.repo.GetUser(ctx, user.Login)
	if err == nil {
		return userdomain.ErrUserExists
	}

	if !errors.Is(err, userdomain.ErrNotFound) {
		return err
	}

	if err := i.auth.AuthorizeUser(ctx, user); err != nil {
		if errors.Is(err, sutclient.ErrBadUser) {
			return userdomain.ErrBadUser
		}

		return err
	}

	return i.repo.AddUser(ctx, user)
}

func (i *Interactor) UpdateUser(ctx context.Context, user userdomain.UpdateUser) error {
	if err := i.validateUpdateRequest(user); err != nil {
		return err
	}

	dbUser, err := i.GetUserByLogin(ctx, user.Login)
	if err != nil {
		return err
	}

	if dbUser.Login != user.Login || dbUser.Password != user.OldPassword {
		return userdomain.ErrNotFound
	}

	if dbUser.Password == user.NewPassword {
		return nil
	}

	newUser := userdomain.User{
		Login:    user.Login,
		Password: user.NewPassword,
	}

	if err := i.auth.AuthorizeUser(ctx, newUser); err != nil {
		if errors.Is(err, sutclient.ErrBadUser) {
			return userdomain.ErrBadUser
		}

		return err
	}

	return i.repo.AddUser(ctx, newUser)
}

func (i *Interactor) DeleteUser(ctx context.Context, user userdomain.User) error {
	if err := i.validateUser(user); err != nil {
		return err
	}

	dbUser, err := i.GetUserByLogin(ctx, user.Login)
	if err != nil {
		return err
	}

	if dbUser.Login != user.Login || dbUser.Password != user.Password {
		return userdomain.ErrNotFound
	}

	if err := i.repo.DeleteUser(ctx, user); err != nil {
		return err
	}

	return i.repo.DeleteUserLastLogin(ctx, user.Login)
}
