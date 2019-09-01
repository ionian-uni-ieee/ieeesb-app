package users_test

import (
	"time"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/users"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Mock passwords
var mockUserPass = "joepass"
var mockUserPass2 = "johndoepass"
var mockUserPass3 = "nickpass"

var mockUserHash, _ = bcrypt.GenerateFromPassword([]byte(mockUserPass), bcrypt.MinCost)
var mockUserHash2, _ = bcrypt.GenerateFromPassword([]byte(mockUserPass2), bcrypt.MinCost)
var mockUserHash3, _ = bcrypt.GenerateFromPassword([]byte(mockUserPass3), bcrypt.MinCost)

// Test users
var mockUser = models.User{
	ID:       primitive.NewObjectID(),
	Username: "joe",
	Password: string(mockUserHash),
	Email:    "joe@mail.com",
	Fullname: "Joe Smith",
	Permissions: models.Permissions{
		Users:    true,
		Events:   false,
		Tickets:  false,
		Sponsors: false,
	},
}

var mockUsers = []models.User{
	mockUser,
	models.User{
		ID:       primitive.NewObjectID(),
		Username: "johndoe",
		Password: string(mockUserHash2),
		Email:    "johndoe@mail.com",
		Fullname: "John Doe",
		Permissions: models.Permissions{
			Users:    false,
			Events:   false,
			Tickets:  false,
			Sponsors: true,
		},
	},
	models.User{
		ID:       primitive.NewObjectID(),
		Username: "nick",
		Password: string(mockUserHash3),
		Email:    "nick@mail.com",
		Fullname: "Nick Brian",
		Permissions: models.Permissions{
			Users:    false,
			Events:   true,
			Tickets:  false,
			Sponsors: false,
		},
	},
}

// Mock Sessions
var mockSession = models.Session{
	ID:       primitive.NewObjectID(),
	Username: mockUser.Username,
	Expires:  time.Now().Unix() + 30*60*1000,
}

func makeController() (*testingDatabase.DatabaseSession, *users.Controller) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	reps := repositories.MakeRepositories(database)
	controller := users.MakeController(reps)

	return database, controller
}
