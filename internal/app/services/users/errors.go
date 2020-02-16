package users

import "errors"

var ErrUsernameEmpty = errors.New("User username can't be empty")
var ErrInvalidEmail = errors.New("Invalid email")
