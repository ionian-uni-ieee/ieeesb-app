package sponsors

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
)

type Controller struct {
	repositories *repositories.Repositories
}

func MakeNewController(r *repositories.Repositories) *Controller {
	return &Controller{
		r,
	}
}
