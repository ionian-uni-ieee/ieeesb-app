package users

import "github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"

// Delete deletes a user from the user repository
func (s *Service) Delete(userID string) validation.Validation {
	v := validation.Validation{}

	err := s.repositories.Users.DeleteByID(userID)

	v.AddError("repositories", err)
	return v
}
