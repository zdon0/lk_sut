package user

import (
	domainUser "lk_sut/internal/domain/user"
	"lk_sut/pkg/dto"
)

func makeUserDomain(user dto.User) domainUser.User {
	return domainUser.User{
		Login:    user.Login,
		Password: user.Password,
	}
}

func makeUpdateUserDomain(updateUser dto.UpdateUser) domainUser.UpdateUser {
	return domainUser.UpdateUser{
		Login:       updateUser.Login,
		OldPassword: updateUser.OldPassword,
		NewPassword: updateUser.NewPassword,
	}
}
