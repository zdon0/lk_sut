package domain

import "errors"

var ErrNotFound = errors.New("entity not found")
var ErrBadUser = errors.New("wrong login or password")
var ErrUserExists = errors.New("user with same login exists")
