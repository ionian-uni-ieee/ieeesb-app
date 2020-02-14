package users

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"

// Service is a business rule controlled entity that uses repositories to perform actions on model entities
type Service struct {
	repositories *repositories.Repositories
}

// MakeService is a Users Service factory
func MakeService(repositories *repositories.Repositories) *Service {
	return &Service{
		repositories: repositories,
	}
}

// SortBy defines by which field to sort the user values by
type SortBy int

const (
	// None for no sorting
	None SortBy = iota
	// Username sorting
	Username
	// Fullname sorting
	Fullname
	// Email sorting
	Email
)
