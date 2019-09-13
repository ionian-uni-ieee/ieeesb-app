package users

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"

func (c *Controller) GetUsers(skip int64, limit int64) ([]models.User, error) {
	users, err := c.repositories.UsersRepository.Find(nil, skip, limit)

	return users, err
}
