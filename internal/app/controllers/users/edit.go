package users

import (
	"errors"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
)

func (c *Controller) Edit(userID string, edit models.User) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	return nil
}
