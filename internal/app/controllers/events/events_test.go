package events_test

import (
	"reflect"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/events"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockEvent = models.Event{
	ID:          primitive.NewObjectID(),
	Name:        "arduino",
	Description: "Learn how to handle Arduino microcontrollers",
	Tags:        []string{"arduino", "hardware"},
	Type:        "seminar",
	Sponsors:    []models.Sponsor{},
	Logo:        models.MediaMeta{},
	Media:       []models.MediaMeta{},
}

var mockEvents = []models.Event{
	mockEvent,
	models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "C++ from zero to hero",
		Description: "Programming C++ for beginners that want to become heroes",
		Tags:        []string{"language", "programming"},
		Type:        "workshop",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	},
	models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "What is Machine Learning",
		Description: "Learn what is machine learning and how to train your own models like a pro",
		Tags:        []string{"statistics", "machine learning", "programming", "python"},
		Type:        "workshop",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	},
}

func makeController() (*testingDatabase.DatabaseSession, *events.Controller) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	reps := repositories.MakeRepositories(database)
	controller := events.MakeController(reps)

	return database, controller
}

func isEventEqualToDbRow(db *testingDatabase.DatabaseSession, event models.Event, rowIndex int) bool {

	fieldNames, err := reflections.GetFieldNames(&event)

	if err != nil {
		panic(err)
	}

	events := db.GetCollection("events").(*testingDatabase.Collection)

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

// Clears the collection's data
func resetCollection(db *testingDatabase.DatabaseSession, collectionName string) {
	collection := db.GetCollection(collectionName).(*testingDatabase.Collection)
	for key, _ := range collection.Columns {
		collection.Columns[key] = []interface{}{}
	}
}

// setupData resets the collection and inserts an array of data in it
func setupData(db *testingDatabase.DatabaseSession, collectionName string, data ...models.Event) {
	resetCollection(db, collectionName)

	ticketFieldNames, err := reflections.GetFieldNames(&models.Event{})
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
