package sponsors

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"
)

// Controller keeps the necessary structure for all controllers to function correctly
type Controller struct {
	repositories *repositories.Repositories
}

// MakeController creates a new sponsor use case caller
func MakeController(r *repositories.Repositories) *Controller {
	return &Controller{
		r,
	}
}
