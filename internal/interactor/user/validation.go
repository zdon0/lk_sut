package user

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"lk_sut/internal/domain"
	domainUser "lk_sut/internal/domain/user"
)

const minPasswordLength = 6

func (i *Interactor) validateCreateUser(ctx context.Context, user domainUser.User) error {
	err := validation.ValidateStruct(
		&user,
		validation.Field(&user.Login, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(minPasswordLength, 100)),
	)
	if err != nil {
		return err
	}

	_, err = i.userBehavior.Repo().GetUser(ctx, user.Login)
	if err == nil {
		return domainUser.ErrUserExists
	}

	if !errors.Is(err, domain.ErrNotFound) {
		return err
	}

	return nil
}

func (i *Interactor) validateUpdateUser(ctx context.Context, newUser domainUser.UpdateUser) error {
	err := validation.ValidateStruct(
		&newUser,
		validation.Field(&newUser.Login, validation.Required, is.Email),
		validation.Field(&newUser.OldPassword, validation.Required, validation.Length(minPasswordLength, 100)),
		validation.Field(&newUser.NewPassword, validation.Required, validation.Length(minPasswordLength, 100)),
	)
	if err != nil {
		return err
	}

	dbUser, err := i.userBehavior.Repo().GetUser(ctx, newUser.Login)
	if err != nil {
		return err
	}

	if dbUser.Login != newUser.Login || dbUser.Password != newUser.OldPassword {
		return domain.ErrNotFound
	}

	return nil
}

func (i *Interactor) validateDeleteUser(ctx context.Context, user domainUser.User) error {
	err := validation.ValidateStruct(
		&user,
		validation.Field(&user.Login, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(minPasswordLength, 100)),
	)
	if err != nil {
		return err
	}

	dbUser, err := i.userBehavior.Repo().GetUser(ctx, user.Login)
	if err != nil {
		return err
	}

	if dbUser.Login != user.Login || dbUser.Password != user.Password {
		return domain.ErrNotFound
	}

	return nil
}
