package users

import (
	"errors"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Register creates and stores a new user to the repository
func (c *Controller) Register(username string, password string, email string, fullname string) (string, error) {
	if username == "" {
		return "", errors.New("Username is empty string")
	}

	if password == "" {
		return "", errors.New("Password is empty string")
	}

	if email == "" {
		return "", errors.New("Email is empty string")
	}

	if fullname == "" {
		return "", errors.New("Fullname is empty string")
	}

	isDuplicate := c.repositories.UsersRepository.IsDuplicate(email, username, fullname)

	if isDuplicate {
		return "", errors.New("A user with that username, fullname or email already exists")
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Password: string(passwordEncrypted),
		Email:    email,
		Fullname: fullname,
		Permissions: models.Permissions{
			Users:    false,
			Events:   false,
			Tickets:  false,
			Sponsors: false,
		},
	}

	_, err = c.repositories.UsersRepository.InsertOne(newUser)

	return newUser.ID.Hex(), err
}
