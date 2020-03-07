package events

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

// Validates that an event object is valid
func (s *Service) Validate(event models.Event) *validation.Validation {
	v := &validation.Validation{}

	if event.GetName() == "" {
		v.AddError("name", ErrNameEmpty)
	}

	if event.GetDate() < 946728000 { // 01/01/2000
		v.AddError("date", ErrInvalidDate)
	}

	return v
}
