package users_test

import (
	"reflect"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/users"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
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

func makeController() (*testingDatabase.DatabaseSession, *users.Controller) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	reps := repositories.MakeRepositories(database)
	controller := users.MakeController(reps)

	return database, controller
}

func isUserEqualToDbRow(db *testingDatabase.DatabaseSession, user models.User, row int) bool {

	fieldNames, err := reflections.GetFieldNames(&user)

	if err != nil {
		panic(err)
	}

	users := db.GetCollection("users").(*testingDatabase.Collection)

	for _, fieldName := range fieldNames {
		field, err := reflections.GetField(&user, fieldName)

		if err != nil {
			panic(err)
		}

		if !reflect.DeepEqual(users.Columns[fieldName][row], field) {
			return false
		}
	}

	return true
}

func getUserAtRow(db *testingDatabase.DatabaseSession, row int) (user *models.User) {
	users := db.GetCollection("users").(*testingDatabase.Collection)

	if isCollectionEmpty(db, "users") {
		return nil
	}

	fieldNames, err := reflections.GetFieldNames(&models.User{})

	if err != nil {
		panic(err)
	}

	user = &models.User{}

	for _, fieldName := range fieldNames {
		field := users.Columns[fieldName][row]
		err := reflections.SetField(user, fieldName, field)
		if err != nil {
			panic(err)
		}
	}

	return user
}

func isCollectionEmpty(db *testingDatabase.DatabaseSession, collectionName string) bool {
	users := db.GetCollection(collectionName).(*testingDatabase.Collection)

	if users == nil ||
		users.Columns == nil ||
		len(users.Columns) == 0 {
		return true
	}

	for _, column := range users.Columns {
		if len(column) != 0 {
			return false
		}
	}

	return true
}

// Clears the collection's data
func resetCollection(db *testingDatabase.DatabaseSession, collectionName string) {
	collection := db.GetCollection(collectionName).(*testingDatabase.Collection)
	for key, _ := range collection.Columns {
		collection.Columns[key] = []interface{}{}
	}
}

// setupData resets the collection and inserts an array of data in it
func setupData(db *testingDatabase.DatabaseSession, collectionName string, data ...models.User) {
	resetCollection(db, collectionName)

	ticketFieldNames, err := reflections.GetFieldNames(&models.User{})
	if err != nil {
		panic(err)
	}

	collection := db.GetCollection(collectionName).(*testingDatabase.Collection)
	for _, item := range data {
		for _, fieldName := range ticketFieldNames {
			fieldValue, err := reflections.GetField(&item, fieldName)
			if err != nil {
				panic(err)
			}

			collection.Columns[fieldName] = append(
				collection.Columns[fieldName],
				fieldValue)
		}
	}
}
