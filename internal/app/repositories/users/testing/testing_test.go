package users_test

import (
	"reflect"
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	users "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/users/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() *users.Repository {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	usersRepository := users.MakeRepository(database)

	return usersRepository
}

// Clears the collection's data
func resetCollection(repository *users.Repository) {
	for key, _ := range repository.Collection.Columns {
		repository.Collection.Columns[key] = []interface{}{}
	}
}

// setupData resets the collection and inserts an array of data in it
func setupData(repository *users.Repository, data ...models.User) {
	resetCollection(repository)

	userFieldNames, err := reflections.GetFieldNames(reflect.TypeOf(models.User{}))
	if err != nil {
		panic(err)
	}

	for _, item := range data {
		for _, fieldName := range userFieldNames {
			fieldValue, err := reflections.GetFieldValue(item, fieldName)
			if err != nil {
				panic(err)
			}

			repository.Collection.Columns[fieldName] = append(
				repository.Collection.Columns[fieldName],
				fieldValue)
		}
	}
}

func TestFindByID(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	user := models.User{
		ID:          primitive.NewObjectID(),
		Username:    "username",
		Password:    "password",
		Email:       "email",
		Fullname:    "fullname",
		Permissions: models.Permissions{},
	}
	setupData(usersRepository, user)

	userFound, err := usersRepository.FindByID(user.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if userFound == nil {
		t.Error("Expected result to be an user object, got nil instead")
	}

	if userFound != nil && userFound.ID != user.ID {
		t.Error("Expected user's id to be", user.ID.Hex(), "but is", userFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	user := models.User{
		ID:          primitive.NewObjectID(),
		Username:    "username",
		Password:    "password",
		Email:       "email",
		Fullname:    "fullname",
		Permissions: models.Permissions{},
	}
	setupData(usersRepository, user)

	err := usersRepository.UpdateByID(user.ID.Hex(), map[string]interface{}{
		"Username": "new username",
	})

	if err != nil {
		t.Error(err)
	}

	if username := usersRepository.Collection.Columns["Username"][0]; username != "new username" {
		t.Error("Expected username to be 'new name', but instead got", username)
	}
}

func TestDeleteByID(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	user := models.User{
		ID:          primitive.NewObjectID(),
		Username:    "username",
		Password:    "password",
		Email:       "email",
		Fullname:    "fullname",
		Permissions: models.Permissions{},
	}
	setupData(usersRepository, user)

	err := usersRepository.DeleteByID(user.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	for key, column := range usersRepository.Collection.Columns {
		if len(column) > 0 {
			t.Error("Expected column", key, "to have length of 0, but instead got", len(column))
		}
	}
}

func TestFind(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username",
			Password:    "password",
			Email:       "email",
			Fullname:    "fullname",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username2",
			Password:    "password2",
			Email:       "email2",
			Fullname:    "fullname2",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username2",
			Password:    "password3",
			Email:       "email3",
			Fullname:    "fullname3",
			Permissions: models.Permissions{},
		},
	}
	setupData(usersRepository, users...)

	usersFound, err := usersRepository.Find(map[string]interface{}{
		"Username": "username2",
	})

	if err != nil {
		t.Error(err)
	}

	if len(usersFound) != 2 {
		t.Error("Expected len(users) to be 2, instead got", len(usersFound))
	}

	if usersFound[0].Username != usersFound[1].Username {
		t.Error("Expected username to equal to each other, instead got",
			usersFound[0].Username,
			usersFound[1].Username)
	}
}

func TestFindOne(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username",
			Password:    "password",
			Email:       "email",
			Fullname:    "fullname",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username2",
			Password:    "password2",
			Email:       "email2",
			Fullname:    "fullname2",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username3",
			Password:    "password3",
			Email:       "email3",
			Fullname:    "fullname3",
			Permissions: models.Permissions{},
		},
	}
	setupData(usersRepository, users...)

	userFound, err := usersRepository.FindOne(map[string]interface{}{
		"Username": "username2",
	})

	if err != nil {
		t.Error(err)
	}

	if userFound.Username != "username2" {
		t.Error("Expected username to equal 'username2', instead got", userFound.Username)
	}
}

func TestUpdateMany(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username",
			Password:    "password",
			Email:       "email",
			Fullname:    "fullname",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username2",
			Password:    "password2",
			Email:       "email2",
			Fullname:    "fullname2",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username3",
			Password:    "password3",
			Email:       "email3",
			Fullname:    "fullname3",
			Permissions: models.Permissions{},
		},
	}
	setupData(usersRepository, users...)
}

func TestDeleteMany(t *testing.T) {

	usersRepository := makeRepository()

	// Regular example
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username",
			Password:    "password",
			Email:       "email",
			Fullname:    "fullname",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username2",
			Password:    "password2",
			Email:       "email2",
			Fullname:    "fullname2",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username3",
			Password:    "password3",
			Email:       "email3",
			Fullname:    "fullname3",
			Permissions: models.Permissions{},
		},
	}
	setupData(usersRepository, users...)
}

func TestInsertOne(t *testing.T) {

	usersRepository := makeRepository()

	// Regular example
	resetCollection(usersRepository)

	newUser := models.User{
		ID:          primitive.NewObjectID(),
		Username:    "username",
		Password:    "password",
		Email:       "email",
		Fullname:    "fullname",
		Permissions: models.Permissions{},
	}
	insertedID, err := usersRepository.InsertOne(newUser)

	if err != nil {
		t.Error(err)
	}

	if insertedID != newUser.ID.Hex() {
		t.Error("Expected inserted id to be ", newUser.ID.Hex(), "but instead got", insertedID)
	}
}

func TestInsertMany(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	resetCollection(usersRepository)

	newUsers := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username",
			Password:    "password",
			Email:       "email",
			Fullname:    "fullname",
			Permissions: models.Permissions{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username2",
			Password:    "password2",
			Email:       "email2",
			Fullname:    "fullname2",
			Permissions: models.Permissions{},
		},
	}

	insertedIDs, err := usersRepository.InsertMany(newUsers)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != newUsers[0].ID.Hex() ||
		insertedIDs[1] != newUsers[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", newUsers[0].ID.Hex(), newUsers[1].ID.Hex(), "but instead got", insertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	usersRepository := makeRepository()

	// Name is duplicate
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Username:    "username2",
			Password:    "password2",
			Email:       "email2",
			Fullname:    "fullname2",
			Permissions: models.Permissions{},
		},
	}
	setupData(usersRepository, users...)

	isDuplicate := usersRepository.IsDuplicate("email2", "username2", "fullname2")

	if !isDuplicate {
		t.Error("Expected name to be duplicate")
	}
}
