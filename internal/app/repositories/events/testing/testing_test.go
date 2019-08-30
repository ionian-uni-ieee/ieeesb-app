package events_test

import (
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	events "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/events/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() *events.Repository {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	eventsRepository := events.MakeRepository(database)

	return eventsRepository
}

// Clears the collection's data
func resetCollection(repository *events.Repository) {
	for key, _ := range repository.Collection.Columns {
		repository.Collection.Columns[key] = []interface{}{}
	}
}

// setupData resets the collection and inserts an array of data in it
func setupData(repository *events.Repository, data ...models.Event) {
	resetCollection(repository)

	eventFieldNames, err := reflections.GetFieldNames(&models.Event{})
	if err != nil {
		panic(err)
	}

	for _, item := range data {
		for _, fieldName := range eventFieldNames {
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
	eventsRepository := makeRepository()

	// Regular example
	event := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(eventsRepository, event)

	eventFound, err := eventsRepository.FindByID(event.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if eventFound == nil {
		t.Error("Expected result to be an event object, got nil instead")
	}

	if eventFound != nil && eventFound.ID != event.ID {
		t.Error("Expected event's id to be", event.ID.Hex(), "but is", eventFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	eventsRepository := makeRepository()

	// Regular example
	event := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(eventsRepository, event)

	err := eventsRepository.UpdateByID(event.ID.Hex(), map[string]interface{}{
		"Name": "new name",
	})

	if err != nil {
		t.Error(err)
	}

	if name := eventsRepository.Collection.Columns["Name"][0]; name != "new name" {
		t.Error("Expected event name to be 'new name', but instead got", name)
	}
}

func TestDeleteByID(t *testing.T) {
	eventsRepository := makeRepository()

	// Regular example
	event := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(eventsRepository, event)

	err := eventsRepository.DeleteByID(event.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if ids := eventsRepository.Collection.Columns["ID"]; len(ids) == 0 {
		t.Error("Expected event id to have length of 0, but instead got", len(ids))
	}
}

func TestFind(t *testing.T) {
	eventsRepository := makeRepository()

	// Regular example
	events := []models.Event{
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
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
	setupData(eventsRepository, events...)

	eventsFound, err := eventsRepository.Find(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if len(eventsFound) != 2 {
		t.Error("Expected len(events) to be 2, instead got", len(eventsFound))
	}

	if eventsFound[0].Description != eventsFound[1].Description {
		t.Error("Expected events' description to equal to each other, instead got",
			eventsFound[0].Description,
			eventsFound[1].Description)
	}
}

func TestFindOne(t *testing.T) {
	eventsRepository := makeRepository()

	// Regular example
	events := []models.Event{
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
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
	setupData(eventsRepository, events...)

	eventFound, err := eventsRepository.FindOne(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if eventFound.Description != "desc3" {
		t.Error("Expected event description to equal 'desc3', instead got", eventFound.Description)
	}
}

func TestUpdateMany(t *testing.T) {
	eventsRepository := makeRepository()

	// Regular example
	events := []models.Event{
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
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
	setupData(eventsRepository, events...)
}

func TestDeleteMany(t *testing.T) {

	eventsRepository := makeRepository()

	// Regular example
	events := []models.Event{
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
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
	setupData(eventsRepository, events...)
}

func TestInsertOne(t *testing.T) {

	eventsRepository := makeRepository()

	// Regular example
	resetCollection(eventsRepository)

	newEvent := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "name3",
		Description: "desc3",
		Tags:        []string{"tag3"},
		Type:        "workshop",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	insertedID, err := eventsRepository.InsertOne(newEvent)

	if err != nil {
		t.Error(err)
	}

	if insertedID != newEvent.ID.Hex() {
		t.Error("Expected inserted id to be ", newEvent.ID.Hex(), "but instead got", insertedID)
	}
}

func TestInsertMany(t *testing.T) {
	eventsRepository := makeRepository()

	// Regular example
	resetCollection(eventsRepository)

	newEvents := []models.Event{
		models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Event{
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

	insertedIDs, err := eventsRepository.InsertMany(newEvents)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != newEvents[0].ID.Hex() ||
		insertedIDs[1] != newEvents[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", newEvents[0].ID.Hex(), newEvents[1].ID.Hex(), "but instead got", insertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	eventsRepository := makeRepository()

	// Name is duplicate
	events := []models.Event{
		models.Event{
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
	setupData(eventsRepository, events...)

	isDuplicate := eventsRepository.IsDuplicate("name2")

	if !isDuplicate {
		t.Error("Expected name to be duplicate")
	}
}
