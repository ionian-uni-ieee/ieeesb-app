package sponsors

type Controller struct {
	repositories *repositories.Repositories
}

func MakeNewController(r *repositories.Repository) *Controller {
	return &Controller{
		r,
	}
}
