package users

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
)

type Controller struct {
	repositories *repositories.Repositories
}

// MakeController creates a new user use case caller
func MakeController(r *repositories.Repositories) *Controller {
	return &Controller{
		r,
	}
}
