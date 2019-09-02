package users

import (
	"errors"
	"reflect"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
)

// Edit updates a user to the model provided
func (c *Controller) Edit(userID string, update map[string]interface{}) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	isUpdateMapValid := reflections.AreMapFieldsValid(
		reflect.TypeOf(models.User{}),
		update,
	)

	if !isUpdateMapValid {
		return errors.New("Update's field(s) are not valid for User model")
	}

	err := c.repositories.UsersRepository.UpdateByID(userID, update)

	return err
}
