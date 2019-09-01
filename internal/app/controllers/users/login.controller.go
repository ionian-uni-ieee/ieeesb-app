package users

import (
	"errors"
	"strconv"
	"time"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Login creates a new stored session
func (c *Controller) Login(username string, password string) (string, error) {
	if username == "" {
		return "", errors.New("Username is empty string")
	}

	if password == "" {
		return "", errors.New("Password is empty string")
	}

	user, err := c.repositories.UsersRepository.FindOne(map[string]interface{}{"Username": username})

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("No user with this username exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("Password not verified")
	}

	duration, _ := time.ParseDuration(strconv.Itoa(24*30) + "h")
	newSession := models.Session{
		ID:       primitive.NewObjectID(),
		Username: username,
		Expires:  time.Now().Add(duration).Unix(),
	}

	_, err = c.repositories.SessionsRepository.InsertOne(newSession)

	return newSession.ID.Hex(), err
}
