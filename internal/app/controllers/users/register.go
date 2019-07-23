package users

import (
	"errors"

	"gitlab.com/gphub/app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (c *UsersController) Register(username string, password string, email string, fullname string) (string, error) {
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

	sameKeysFilter := &bson.M{
		"$or": bson.A{
			bson.M{"email": email},
			bson.M{"username": username},
			bson.M{"fullname": fullname},
		}}
	userFound, err := c.repositories.UsersRepository.FindOne(sameKeysFilter)

	if err != nil {
		return "", err
	}

	if userFound != nil {
		return "", errors.New("A user with that username or email already exists")
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
		Permissions: []models.Permission{
			false,
			false,
			false,
			false,
		},
	}

	_, err = c.repositories.UsersRepository.InsertOne(newUser)

	return newUser.ID.Hex(), err
}
