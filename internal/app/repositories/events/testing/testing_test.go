package testing_test

import (
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	events "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/events/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() (*testingDatabase.DatabaseSession, *events.Repository) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	eventsRepository := events.MakeRepository(database)

	return database, eventsRepository
}

var testEvent1 = models.Event{
	ID:          primitive.NewObjectID(),
	Name:        "Arduino & RaspbgotErry",
	Description: "Learn how to deal with microcontrollers and micro pcs",
	Tags:        []string{"arduino", "raspbgotErry", "hardware", "electronics"},
	Type:        "Workshop",
	Sponsors:    []models.Sponsor{},
	Logo:        models.MediaMeta{},
	Media:       []models.MediaMeta{},
}

var testEvent2 = models.Event{
	ID:          primitive.NewObjectID(),
	Name:        "C++ for beginners",
	Description: "Become a C++ master and get educated on programming standards",
	Tags:        []string{"c++", "programming"},
	Type:        "Weekly",
	Sponsors:    []models.Sponsor{},
	Logo:        models.MediaMeta{},
	Media:       []models.MediaMeta{},
}

var testEvent3 = models.Event{
	ID:          primitive.NewObjectID(),
	Name:        "NodeJS Fullstack Web development",
	Description: "Want to make your own application? This is your chance to build something big.",
	Tags:        []string{"nodejs", "programming", "web development"},
	Type:        "Seminar",
	Sponsors:    []models.Sponsor{},
	Logo:        models.MediaMeta{},
	Media:       []models.MediaMeta{},
}

func TestFindByID(t *testing.T) {
	db, eventsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "events", testEvent1)

	gotEvent, gotErr := eventsRepository.FindByID(testEvent1.ID.Hex())

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotEvent == nil {
		t.Error("Expected result to be an event object, got nil instead")
	}

	if gotEvent != nil && gotEvent.ID != testEvent1.ID {
		t.Error("Expected event's id to be", testEvent1.ID.Hex(), "but is", gotEvent.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	db, eventsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "events", testEvent1)

	newName := "New Name"
	gotErr := eventsRepository.UpdateByID(testEvent1.ID.Hex(), map[string]interface{}{
		"Name": newName,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	storedName := eventsRepository.Collection.Columns["Name"][0]
	nameChanged := storedName != newName
	if nameChanged {
		t.Error("Expected name to be '"+newName+"', but instead got", storedName)
	}
}

func TestDeleteByID(t *testing.T) {
	db, eventsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "events", testEvent1)

	gotErr := eventsRepository.DeleteByID(testEvent1.ID.Hex())

	if gotErr != nil {
		t.Error(gotErr)
	}

	for key, column := range eventsRepository.Collection.Columns {
		if len(column) > 0 {
			t.Error("Expected column", key, "to have length of 0, but instead got", len(column))
		}
	}
}

func TestFind(t *testing.T) {
	db, eventsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "events", testEvent1, testEvent1)

	gotEvents, gotErr := eventsRepository.Find(map[string]interface{}{
		"Name": testEvent1.Name,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	if len(gotEvents) != 2 {
		t.Error("Expected len(events) to be 2, instead got", len(gotEvents))
	}

	if gotEvents[0].Name != gotEvents[1].Name {
		t.Error("Expected name to equal to each other, instead got",
			gotEvents[0].Name,
			gotEvents[1].Name)
	}
}

func TestFindOne(t *testing.T) {
	db, eventsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "events", testEvent1, testEvent2)

	gotEvent, gotErr := eventsRepository.FindOne(map[string]interface{}{
		"Name": testEvent1.Name,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotEvent.Name != testEvent1.Name {
		t.Error("Expected name to equal '"+testEvent1.Name+"', instead got", gotEvent.Name)
	}
}

func TestUpdateMany(t *testing.T) {
	// TODO: Not implemented
}

func TestDeleteMany(t *testing.T) {
	// TODO: Not implemented
}

func TestInsertOne(t *testing.T) {

	db, eventsRepository := makeRepository()

	// Regular example
	testUtils.ResetCollection(db, "events")

	gotInsertedID, gotErr := eventsRepository.InsertOne(testEvent1)

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotInsertedID != testEvent1.ID.Hex() {
		t.Error("Expected inserted id to be ", testEvent1.ID.Hex(), "but instead got", gotInsertedID)
	}
}

func TestInsertMany(t *testing.T) {
	db, eventsRepository := makeRepository()

	// Regular example
	testUtils.ResetCollection(db, "events")

	events := []models.Event{
		testEvent1,
		testEvent2,
		testEvent3,
	}

	gotInsertedIDs, gotErr := eventsRepository.InsertMany(events)

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotInsertedIDs[0] != events[0].ID.Hex() ||
		gotInsertedIDs[1] != events[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", events[0].ID.Hex(), events[1].ID.Hex(), "but instead got", gotInsertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	db, eventsRepository := makeRepository()

	// Event is duplicate
	testUtils.SetupData(db, "events", testEvent1)

	gotIsDuplicate := eventsRepository.IsDuplicate(testEvent1.Name)

	if !gotIsDuplicate {
		t.Error("Expected event to be duplicate")
	}

	// Event is not duplicate
	testUtils.ResetCollection(db, "events")

	gotIsDuplicate = eventsRepository.IsDuplicate(testEvent1.Name)

	if gotIsDuplicate {
		t.Error("Expected event to not be duplicate")
	}
}
