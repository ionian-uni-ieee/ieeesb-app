package users_test

import (
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

	userFieldNames, err := reflections.GetFieldNames(&models.User{})
	if err != nil {
		panic(err)
	}

	for _, item := range data {
		for _, fieldName := range userFieldNames {
			fieldValue, err := reflections.GetField(&item, fieldName)
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
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(usersRepository, user)

	userFound, err := usersRepository.FindByID(event.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if userFound == nil {
		t.Error("Expected result to be an user object, got nil instead")
	}

	if userFound != nil && userFound.ID != event.ID {
		t.Error("Expected user's id to be", user.ID.Hex(), "but is", eventFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	user := models.User{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(usersRepository, user)

	err := usersRepository.UpdateByID(user.ID.Hex(), map[string]interface{}{
		"Name": "new name",
	})

	if err != nil {
		t.Error(err)
	}

	if name := usersRepository.Collection.Columns["Name"][0]; name != "new name" {
		t.Error("Expected user name to be 'new name', but instead got", name)
	}
}

func TestDeleteByID(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	user := models.User{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(usersRepository, user)

	err := usersRepository.DeleteByID(user.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if ids := usersRepository.Collection.Columns["ID"]; len(ids) == 0 {
		t.Error("Expected user id to have length of 0, but instead got", len(ids))
	}
}

func TestFind(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name3",
			Description: "desc3",
			Tags:        []string{"tag3"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(usersRepository, users...)

	usersFound, err := usersRepository.Find(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if len(usersFound) != 2 {
		t.Error("Expected len(users) to be 2, instead got", len(usersFound))
	}

	if usersFound[0].Description != usersFound[1].Description {
		t.Error("Expected users' description to equal to each other, instead got",
			usersFound[0].Description,
			usersFound[1].Description)
	}
}

func TestFindOne(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name3",
			Description: "desc3",
			Tags:        []string{"tag3"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(usersRepository, users...)

	userFound, err := usersRepository.FindOne(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if userFound.Description != "desc3" {
		t.Error("Expected user description to equal 'desc3', instead got", userFound.Description)
	}
}

func TestUpdateMany(t *testing.T) {
	usersRepository := makeRepository()

	// Regular example
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name3",
			Description: "desc3",
			Tags:        []string{"tag3"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
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
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name3",
			Description: "desc3",
			Tags:        []string{"tag3"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(usersRepository, users...)
}

func TestInsertOne(t *testing.T) {

	usersRepository := makeRepository()

	// Regular example
	resetCollection(usersRepository)

	newUser := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "name3",
		Description: "desc3",
		Tags:        []string{"tag3"},
		Type:        "workshop",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
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

	newUsers := []models.Event{
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}

	insertedIDs, err := usersRepository.InsertMany(newUsers)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != newUsers[0].ID.Hex() ||
		insertedIDs[1] != newUsers[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", newUsers[0].ID.Hex(), newEvents[1].ID.Hex(), "but instead got", insertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	usersRepository := makeRepository()

	// Name is duplicate
	users := []models.User{
		models.User{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(usersRepository, users...)

	isDuplicate := usersRepository.IsDuplicate("name2")

	if !isDuplicate {
		t.Error("Expected name to be duplicate")
	}
}
