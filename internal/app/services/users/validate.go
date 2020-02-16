package users

import (
	"regexp"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

func (s *Service) Validate(user models.User) validation.Validation {
	v := validation.Validation{}
	if user.GetUsername() == "" {
		v.AddError("username", ErrUsernameEmpty)
	}

	emailRegex := regexp.MustCompile("^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$")
	isEmailValid := emailRegex.MatchString(user.GetEmail())
	if !isEmailValid {
		v.AddError("email", ErrInvalidEmail)
	}

	return v
}
