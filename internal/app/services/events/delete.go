package events

import "github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"

// Delete deletes a event from the event repository
func (s *Service) Delete(eventID string) validation.Validation {
	v := validation.Validation{}

	err := s.repositories.Events.DeleteByID(eventID)

	v.AddError("repositories", err)
	return v
}
