package user

import (
	"context"

	domainUser "lk_sut/internal/domain/user"
)

type Interactor struct {
	userBehavior *domainUser.Behavior
}

func NewInteractor(userBehavior *domainUser.Behavior) *Interactor {
	return &Interactor{
		userBehavior: userBehavior,
	}
}

func (i *Interactor) AddUser(ctx context.Context, user domainUser.User) error {
	if err := i.validateCreateUser(ctx, user); err != nil {
		return err
	}

	if err := i.userBehavior.Authorize(ctx, user); err != nil {
		return err
	}

	return i.userBehavior.Repo().AddUser(ctx, user)
}

func (i *Interactor) UpdateUser(ctx context.Context, user domainUser.UpdateUser) error {
	if err := i.validateUpdateUser(ctx, user); err != nil {
		return err
	}

	if user.OldPassword == user.NewPassword {
		return nil
	}

	newUser := domainUser.User{
		Login:    user.Login,
		Password: user.NewPassword,
	}

	if err := i.userBehavior.Authorize(ctx, newUser); err != nil {
		return err
	}

	return i.userBehavior.Repo().AddUser(ctx, newUser)
}

func (i *Interactor) DeleteUser(ctx context.Context, user domainUser.User) error {
	if err := i.validateDeleteUser(ctx, user); err != nil {
		return err
	}

	return i.userBehavior.Repo().DeleteUser(ctx, user)
}
