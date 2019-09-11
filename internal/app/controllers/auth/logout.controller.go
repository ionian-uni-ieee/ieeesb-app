package auth

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Logout deletes the given user's session
func (c *Controller) Logout(sessionID string) error {
	if sessionID == "" {
		return errors.New("SessionID is empty string")
	}

	if _, err := primitive.ObjectIDFromHex(sessionID); err != nil {
		return errors.New("SessionID is not a valid ObjectID")
	}

	s, err := c.repositories.SessionsRepository.FindByID(sessionID)

	if err != nil {
		return err
	}

	if s == nil {
		return errors.New("No session with this ID exists")
	}

	err = c.repositories.SessionsRepository.DeleteByID(sessionID)

	return err
}
