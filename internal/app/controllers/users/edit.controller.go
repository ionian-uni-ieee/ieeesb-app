package users

import (
	"errors"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
)

// Edit updates a user to the model provided
func (c *Controller) Edit(userID string, edit models.User) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	return nil
}
