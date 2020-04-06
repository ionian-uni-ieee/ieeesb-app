package sponsors

import "errors"

// ErrTitleEmpty sponsors name is empty
var ErrNameEmpty = errors.New("Sponsor name can't be empty")

// ErrIErrInvalidEmails emails does not match a proper email form
var ErrInvalidEmails = errors.New("Invalid email")

// ErrIErrInvalidPhones phones does not match a proper phone form
var ErrInvalidPhones = errors.New("Invalid phone")
