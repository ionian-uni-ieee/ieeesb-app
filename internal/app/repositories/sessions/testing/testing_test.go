package testing_test

import (
	"testing"
	"time"

	testingDatabase "github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	sessions "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/sessions/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
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

	gotSession, gotErr := sessionsRepository.FindByID(testSession1.ID.Hex())

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotSession == nil {
		t.Error("Expected result to be an session object, got nil instead")
	}

	if gotSession != nil && gotSession.ID != testSession1.ID {
		t.Error("Expected session's id to be", testSession1.ID.Hex(), "but is", gotSession.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sessions", testSession1)

	newUsername := "New Username"
	gotErr := sessionsRepository.UpdateByID(testSession1.ID.Hex(), map[string]interface{}{
		"Username": newUsername,
	})

	if gotErr != nil {
		t.Error(gotErr)
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

	gotErr := sessionsRepository.DeleteByID(testSession1.ID.Hex())

	if gotErr != nil {
		t.Error(gotErr)
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

	gotSessions, gotErr := sessionsRepository.Find(map[string]interface{}{
		"Username": testSession1.Username,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	if len(gotSessions) != 2 {
		t.Error("Expected len(sessions) to be 2, instead got", len(gotSessions))
	}

	if gotSessions[0].Username != gotSessions[1].Username {
		t.Error("Expected sessionname to equal to each other, instead got",
			gotSessions[0].Username,
			gotSessions[1].Username)
	}
}

func TestFindOne(t *testing.T) {
	db, sessionsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sessions", testSession1, testSession2)

	gotSession, gotErr := sessionsRepository.FindOne(map[string]interface{}{
		"Username": testSession1.Username,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotSession.Username != testSession1.Username {
		t.Error("Expected sessionname to equal '" + testSession1.Username + "', instead got " + gotSession.Username)
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

	gotInsertedID, gotErr := sessionsRepository.InsertOne(testSession1)

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotInsertedID != testSession1.ID.Hex() {
		t.Error("Expected inserted id to be ", testSession1.ID.Hex(), "but instead got", gotInsertedID)
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

	gotInsertedIDs, gotErr := sessionsRepository.InsertMany(sessions)

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotInsertedIDs[0] != sessions[0].ID.Hex() ||
		gotInsertedIDs[1] != sessions[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", sessions[0].ID.Hex(), sessions[1].ID.Hex(), "but instead got", gotInsertedIDs)
	}
}
