package events_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/events"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeController() (*testingDatabase.DatabaseSession, *events.Controller) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	reps := repositories.MakeRepositories(database)
	controller := events.MakeController(reps)

	return database, controller
}

func isEventEqualToDbRow(db *testingDatabase.DatabaseSession, event models.Event, rowIndex int) bool {

	events := db.GetCollection("events").(*testingDatabase.Collection)
	fieldNames, err := reflections.GetFieldNames(&event)

	if err != nil {
		panic(err)
	}

	for _, fieldName := range fieldNames {
		field, err := reflections.GetField(&event, fieldName)

		if err != nil {
			panic(err)
		}

		if !reflect.DeepEqual(events.Columns[fieldName][rowIndex], field) {
			return false
		}
	}

	return true
}

func TestAdd(t *testing.T) {
	// Setup
	db, eventsController := makeController()

	// Normal
	event := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "Arduino",
		Description: "desc",
		Tags:        []string{"arduino", "hardware"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	eventsController.Add(event)

	if !isEventEqualToDbRow(db, event, 0) {
		t.Error("Expected event to have been added in database")
	}
}
