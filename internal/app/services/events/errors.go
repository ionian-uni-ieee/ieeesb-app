package events

import "errors"

// ErrTitleEmpty events name is empty
var ErrNameEmpty = errors.New("Event name can't be empty")

// ErrInvalidDate date does not match a proper date
var ErrInvalidDate = errors.New("Event date can't be before 2000")
