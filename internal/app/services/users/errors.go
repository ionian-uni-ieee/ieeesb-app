package users

import "errors"

// ErrErrUsernameEmpty username string is empty
var ErrUsernameEmpty = errors.New("User username can't be empty")

// ErrIErrInvalidEmail email does not match a proper email form
var ErrInvalidEmail = errors.New("Invalid email")

// ErrUserExists user with same username, email or fullname already exists
var ErrUserExists = errors.New("User with same username, email or fullname already exists")
