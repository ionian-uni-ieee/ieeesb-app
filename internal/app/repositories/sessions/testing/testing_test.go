package sessions_test

import (
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	sessions "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sessions/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() *sessions.Repository {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	sessionsRepository := sessions.MakeRepository(database)

	return sessionsRepository
}

// Clears the collection's data
func resetCollection(repository *sessions.Repository) {
	for key, _ := range repository.Collection.Columns {
		repository.Collection.Columns[key] = []interface{}{}
	}
}

// setupData resets the collection and inserts an array of data in it
func setupData(repository *sessions.Repository, data ...models.Session) {
	resetCollection(repository)

	sessionFieldNames, err := reflections.GetFieldNames(&models.Session{})
	if err != nil {
		panic(err)
	}

	for _, item := range data {
		for _, fieldName := range sessionFieldNames {
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
	sessionsRepository := makeRepository()

	// Regular example
	session := models.Session{
		ID:       primitive.NewObjectID(),
		Username: "bugs",
		Expires:  1,
	}
	setupData(sessionsRepository, session)

	sessionFound, err := sessionsRepository.FindByID(session.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if sessionFound == nil {
		t.Error("Expected result to be an session object, got nil instead")
	}

	if sessionFound != nil && sessionFound.ID != session.ID {
		t.Error("Expected session's id to be", session.ID.Hex(), "but is", sessionFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	session := models.Session{
		ID:       primitive.NewObjectID(),
		Username: "bugs",
		Expires:  1,
	}
	setupData(sessionsRepository, session)

	err := sessionsRepository.UpdateByID(session.ID.Hex(), map[string]interface{}{
		"Username": "not bugs",
	})

	if err != nil {
		t.Error(err)
	}

	if name := sessionsRepository.Collection.Columns["Username"][0]; name != "not bugs" {
		t.Error("Expected username to be 'not bugs', but instead got", name)
	}
}

func TestDeleteByID(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	session := models.Session{
		ID:       primitive.NewObjectID(),
		Username: "bugs",
		Expires:  1,
	}
	setupData(sessionsRepository, session)

	err := sessionsRepository.DeleteByID(session.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if ids := sessionsRepository.Collection.Columns["ID"]; len(ids) == 0 {
		t.Error("Expected session id to have length of 0, but instead got", len(ids))
	}
}

func TestFind(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	sessions := []models.Session{
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "bugs",
			Expires:  1,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "jake",
			Expires:  2,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "jake",
			Expires:  3,
		},
	}
	setupData(sessionsRepository, sessions...)

	sessionsFound, err := sessionsRepository.Find(map[string]interface{}{
		"Username": "jake",
	})

	if err != nil {
		t.Error(err)
	}

	if len(sessionsFound) != 2 {
		t.Error("Expected len(sessions) to be 2, instead got", len(sessionsFound))
	}

	if sessionsFound[0].Username != sessionsFound[1].Username {
		t.Error("Expected username to equal to each other, instead got",
			sessionsFound[0].Username,
			sessionsFound[1].Username)
	}
}

func TestFindOne(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	sessions := []models.Session{
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "username",
			Expires:  1,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "username2",
			Expires:  2,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "jake",
			Expires:  3,
		},
	}
	setupData(sessionsRepository, sessions...)

	sessionFound, err := sessionsRepository.FindOne(map[string]interface{}{
		"Username": "jake",
	})

	if err != nil {
		t.Error(err)
	}

	if sessionFound.Username != "jake" {
		t.Error("Expected session description to equal 'desc3', instead got", sessionFound.Username)
	}
}

func TestUpdateMany(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	sessions := []models.Session{
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "bugs",
			Expires:  1,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "jake",
			Expires:  2,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "jake",
			Expires:  3,
		},
	}
	setupData(sessionsRepository, sessions...)
}

func TestDeleteMany(t *testing.T) {

	sessionsRepository := makeRepository()

	// Regular example
	sessions := []models.Session{
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "bugs",
			Expires:  1,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "jake",
			Expires:  2,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "jake",
			Expires:  3,
		},
	}
	setupData(sessionsRepository, sessions...)
}

func TestInsertOne(t *testing.T) {

	sessionsRepository := makeRepository()

	// Regular example
	resetCollection(sessionsRepository)

	newSession := models.Session{
		ID:       primitive.NewObjectID(),
		Username: "jake",
		Expires:  1,
	}
	insertedID, err := sessionsRepository.InsertOne(newSession)

	if err != nil {
		t.Error(err)
	}

	if insertedID != newSession.ID.Hex() {
		t.Error("Expected inserted id to be ", newSession.ID.Hex(), "but instead got", insertedID)
	}
}

func TestInsertMany(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	resetCollection(sessionsRepository)

	newSessions := []models.Session{
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "jake",
			Expires:  1,
		},
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "bugs",
			Expires:  2,
		},
	}

	insertedIDs, err := sessionsRepository.InsertMany(newSessions)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != newSessions[0].ID.Hex() ||
		insertedIDs[1] != newSessions[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", newSessions[0].ID.Hex(), newSessions[1].ID.Hex(), "but instead got", insertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	sessionsRepository := makeRepository()

	// Name is duplicate
	sessions := []models.Session{
		models.Session{
			ID:       primitive.NewObjectID(),
			Username: "bugs",
			Expires:  2,
		},
	}
	setupData(sessionsRepository, sessions...)

	isDuplicate := sessionsRepository.IsDuplicate("bugs")

	if !isDuplicate {
		t.Error("Expected name to be duplicate")
	}
}
