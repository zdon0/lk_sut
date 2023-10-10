package user

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"lk_sut/internal/domain"
)

const minPasswordLength = 6

func (i *Interactor) validateUser(user domain.User) error {
	return validation.ValidateStruct(
		&user,
		validation.Field(&user.Login, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(minPasswordLength, 100)),
	)
}

func (i *Interactor) validateUpdateRequest(user domain.UpdateUser) error {
	return validation.ValidateStruct(
		&user,
		validation.Field(&user.Login, validation.Required, is.Email),
		validation.Field(&user.OldPassword, validation.Required, validation.Length(minPasswordLength, 100)),
		validation.Field(&user.NewPassword, validation.Required, validation.Length(minPasswordLength, 100), validation.NotIn(user.OldPassword)),
	)
}
