package sponsors

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

//Add adds a new sponsor to the sponsor repository
func (s *Service) Add(sponsor models.Sponsor) (string, *validation.Validation) {
	v := s.Validate(sponsor)
	sponsorID, err := s.repositories.Sponsors.InsertOne(sponsor)

	v.AddError("repositories", err)
	return sponsorID, v
}