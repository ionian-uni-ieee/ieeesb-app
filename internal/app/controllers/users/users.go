package users

type Controller struct {
	repositories *repositories.Repositories
}

// MakeNewController creates a new user use case caller
func MakeNewController(r *repositories.Repository) *Controller {
	return &Controller{
		r,
	}
}
