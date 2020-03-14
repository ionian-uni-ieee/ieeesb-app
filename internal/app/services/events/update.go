package events

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

// Update updates a event with another object
func (s *Service) Update(updatedEvent models.Event) *validation.Validation {
	v := s.Validate(updatedEvent)
	if v.HasError() {
		return v
	}

	err := s.repositories.Users.UpdateByID(updatedEvent.ID.Hex(), updatedEvent)

	v.AddError("repositories", err)
	return v
}
