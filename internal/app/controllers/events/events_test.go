package events_test

import (
	"reflect"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/events"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
)

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
func setupData(db *testingDatabase.DatabaseSession, collectionName string, data ...models.Ticket) {
	resetCollection(db, collectionName)

	ticketFieldNames, err := reflections.GetFieldNames(&models.Ticket{})
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
