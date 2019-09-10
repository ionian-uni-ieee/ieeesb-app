package users

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// ChangePassword updates the user's password only if the old password is correct
func (c *Controller) ChangePassword(userID string, oldPassword string, newPassword string) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		return errors.New("Invalid ObjectID")
	}

	user, err := c.repositories.UsersRepository.FindByID(userID)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("No user with this ID exists")
	}

	if oldPassword == "" {
		return errors.New("Old password is empty string")
	}

	if newPassword == "" {
		return errors.New("New password is empty string")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))

	if err != nil {
		return errors.New("Old password is not valid")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)

	if err != nil {
		return err
	}

	user.Password = string(hash)

	err = c.repositories.UsersRepository.UpdateByID(userID, *user)

	return err
}
