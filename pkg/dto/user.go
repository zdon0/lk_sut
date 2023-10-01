package dto

type User struct {
	Login    string `json:"login" example:"example@mail.com"`
	Password string `json:"password" example:"Password123"`
}

type UpdateUser struct {
	Login       string `json:"login" example:"example@mail.com"`
	OldPassword string `json:"old_password" example:"Password123"`
	NewPassword string `json:"new_password" example:"Password321"`
}
