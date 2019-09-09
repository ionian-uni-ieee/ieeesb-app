package tickets

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"
)

type Controller struct {
	repositories *repositories.Repositories
}

func MakeController(r *repositories.Repositories) *Controller {
	return &Controller{
		r,
	}
}
