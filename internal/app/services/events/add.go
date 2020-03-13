package events

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

// Add adds a new event to the event repository
func (s *Service) Add(event models.Event) (string, *validation.Validation) {
	v := s.Validate(event)
	eventID, err := s.repositories.Events.InsertOne(event)

	v.AddError("repositories", err)
	return eventID, v
}
