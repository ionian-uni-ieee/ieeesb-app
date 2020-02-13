package events

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"

// Service is a business rule controlled entity that uses repositories to perform actions on model entities
type Service struct {
	repositories *repositories.Repositories
}

// MakeService is a Events Service factory
func MakeService(repositories *repositories.Repositories) *Service {
	return &Service{
		repositories: repositories,
	}
}
