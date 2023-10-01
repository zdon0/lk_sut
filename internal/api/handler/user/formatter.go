package user

import (
	"lk_sut/internal/domain"
	"lk_sut/pkg/dto"
)

func makeUserDomain(user dto.User) domain.User {
	return domain.User{
		Login:    user.Login,
		Password: user.Password,
	}
}

func makeUpdateUserDomain(updateUser dto.UpdateUser) domain.UpdateUser {
	return domain.UpdateUser{
		Login:       updateUser.Login,
		OldPassword: updateUser.OldPassword,
		NewPassword: updateUser.NewPassword,
	}
}
