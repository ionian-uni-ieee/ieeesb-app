package testing_test

import (
	"testing"
	"time"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	sessions "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sessions/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() (*testingDatabase.DatabaseSession, *sessions.Repository) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	sessionsRepository := sessions.MakeRepository(database)

	return database, sessionsRepository
}

var testSession1 = models.Session{
	ID:       primitive.NewObjectID(),
	Username: "joejordinson",
	Expires:  time.Now().Unix() + 30*60*1000,
}

var testSession2 = models.Session{
	ID:       primitive.NewObjectID(),
	Username: "billsmith",
	Expires:  time.Now().Unix() + 30*60*1000,
}

var testSession3 = models.Session{
	ID:       primitive.NewObjectID(),
	Username: "johndoe",
	Expires:  time.Now().Unix() + 30*60*1000,
}

func TestFindByID(t *testing.T) {
	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sessions", testSession1)

	sessionFound, err := sessionsRepository.FindByID(testSession1.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if sessionFound == nil {
		t.Error("Expected result to be an session object, got nil instead")
	}

	if sessionFound != nil && sessionFound.ID != testSession1.ID {
		t.Error("Expected session's id to be", testSession1.ID.Hex(), "but is", sessionFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sessions", testSession1)

	newUsername := "New Username"
	err := sessionsRepository.UpdateByID(testSession1.ID.Hex(), map[string]interface{}{
		"Username": newUsername,
	})

	if err != nil {
		t.Error(err)
	}

	storedUsername := sessionsRepository.Collection.Columns["Username"][0]
	nameChanged := storedUsername != newUsername
	if nameChanged {
		t.Error("Expected name to be '"+newUsername+"', but instead got", storedUsername)
	}
}

func TestDeleteByID(t *testing.T) {
	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sessions", testSession1)

	err := sessionsRepository.DeleteByID(testSession1.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	for key, column := range sessionsRepository.Collection.Columns {
		if len(column) > 0 {
			t.Error("Expected column", key, "to have length of 0, but instead got", len(column))
		}
	}
}

func TestFind(t *testing.T) {
	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sessions", testSession1, testSession1)

	sessionsFound, err := sessionsRepository.Find(map[string]interface{}{
		"Username": testSession1.Username,
	})

	if err != nil {
		t.Error(err)
	}

	if len(sessionsFound) != 2 {
		t.Error("Expected len(sessions) to be 2, instead got", len(sessionsFound))
	}

	if sessionsFound[0].Username != sessionsFound[1].Username {
		t.Error("Expected sessionname to equal to each other, instead got",
			sessionsFound[0].Username,
			sessionsFound[1].Username)
	}
}

func TestFindOne(t *testing.T) {
	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sessions", testSession1, testSession2)

	sessionFound, err := sessionsRepository.FindOne(map[string]interface{}{
		"Username": testSession1.Username,
	})

	if err != nil {
		t.Error(err)
	}

	if sessionFound.Username != testSession1.Username {
		t.Error("Expected sessionname to equal '" + testSession1.Username + "', instead got " + sessionFound.Username)
	}
}

func TestUpdateMany(t *testing.T) {
	// TODO: Not implemented
}

func TestDeleteMany(t *testing.T) {
	// TODO: Not implemented
}

func TestInsertOne(t *testing.T) {

	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.ResetCollection(db, "sessions")

	insertedID, err := sessionsRepository.InsertOne(testSession1)

	if err != nil {
		t.Error(err)
	}

	if insertedID != testSession1.ID.Hex() {
		t.Error("Expected inserted id to be ", testSession1.ID.Hex(), "but instead got", insertedID)
	}
}

func TestInsertMany(t *testing.T) {
	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.ResetCollection(db, "sessions")

	sessions := []models.Session{
		testSession1,
		testSession2,
		testSession3,
	}

	insertedIDs, err := sessionsRepository.InsertMany(sessions)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != sessions[0].ID.Hex() ||
		insertedIDs[1] != sessions[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", sessions[0].ID.Hex(), sessions[1].ID.Hex(), "but instead got", insertedIDs)
	}
}
