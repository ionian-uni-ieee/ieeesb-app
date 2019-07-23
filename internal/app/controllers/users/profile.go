package users

import (
	"errors"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Profile returns the user's stored model
func (c *Controller) Profile(sessionID string) (*models.User, error) {
	if sessionID == "" {
		return nil, errors.New("SessionID is empty string")
	}

	if _, err := primitive.ObjectIDFromHex(sessionID); err != nil {
		return nil, errors.New("SessionID is not a valid ObjectID")
	}

	session, err := c.repositories.SessionsRepository.FindByID(sessionID)

	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, errors.New("No session with this ID exists")
	}

	user, err := c.repositories.UsersRepository.FindOne(&bson.M{"username": session.Username})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("No user with this ID exists")
	}

	return user, nil
}
