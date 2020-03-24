package sponsors

import "github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"

//Delete deletes a sponsor from the sponsor repository
func(s *Service) Delete(sponsorID string) validation.Validation {
	v := validation.Validation{}

	err := s.repositories.Sponsors.DeleteByID(sponsorID)

	v.AddError("repositories", err)
	return v
}