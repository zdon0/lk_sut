package user

type User struct {
	Login    string
	Password string
}

type UpdateUser struct {
	Login       string
	OldPassword string
	NewPassword string
}
