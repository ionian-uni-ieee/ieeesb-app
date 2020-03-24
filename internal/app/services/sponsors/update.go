package sponsors

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

// Update updates a sponsor with another object
func (s *Service) Update(updatedSponsor models.Sponsor) *validation.Validation {
	v := s.Validate(updatedSponsor)
	if v.HasError() {
		return v
	}

	err := s.repositories.Sponsors.UpdateByID(updatedSponsor.ID.Hex(), updatedSponsor)

	v.AddError("repositories", err)
	return v
}