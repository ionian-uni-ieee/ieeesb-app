package users

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
)

type Controller struct {
	repositories *repositories.Repositories
}

// MakeNewController creates a new user use case caller
func MakeNewController(r *repositories.Repositories) *Controller {
	return &Controller{
		r,
	}
}