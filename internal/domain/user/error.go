package user

import "errors"

var ErrBadUser = errors.New("wrong login or password")
var ErrUserExists = errors.New("user with same login exists")
