package users

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete removes a user from the database
func (c *UsersController) Delete(userID string) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		return errors.New("Invalid ObjectID")
	}

	err := c.repositories.UsersRepository.DeleteByID(userID)

	return err
}
