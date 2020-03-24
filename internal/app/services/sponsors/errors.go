package sponsors

import "errors"

// ErrTitleEmpty sponsors name is empty
var ErrNameEmpty = errors.New("Sponsor name can't be empty")

// ErrEmailsEmpty sponsors emails are empty
var ErrEmailsEmpty = errors.New("Sponsor emails can't be empty")

// ErrIErrInvalidEmails emails does not match a proper email form
var ErrInvalidEmails = errors.New("Invalid email")

// ErrPhonesEmpty sponsors phones are empty
var ErrPhonesEmpty = errors.New("Sponsor phones can't be empty")

// ErrIErrInvalidPhones phones does not match a proper phone form
var ErrInvalidPhones = errors.New("Invalid phone")