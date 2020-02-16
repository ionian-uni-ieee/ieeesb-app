package users

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

func (s *Service) Update(updatedUser models.User) validation.Validation {
	v := s.Validate(updatedUser)
	if v.HasError() {
		return v
	}

	err := s.repositories.Users.UpdateByID(updatedUser.ID.Hex(), updatedUser)

	v.AddError("repositories", err)
	return v
}
