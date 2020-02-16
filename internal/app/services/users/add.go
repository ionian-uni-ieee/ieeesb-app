package users

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

// Add adds a new user to the user repository
func (s *Service) Add(user models.User) (string, validation.Validation) {
	v := s.Validate(user)
	userID, err := s.repositories.Users.InsertOne(user)

	v.AddError("repositories", err)
	return userID, v
}
