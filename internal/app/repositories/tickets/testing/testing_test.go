package tickets_test

import (
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	tickets "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/tickets/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() *tickets.Repository {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	ticketsRepository := tickets.MakeRepository(database)

	return ticketsRepository
}

// Clears the collection's data
func resetCollection(repository *tickets.Repository) {
	for key, _ := range repository.Collection.Columns {
		repository.Collection.Columns[key] = []interface{}{}
	}
}

// setupData resets the collection and inserts an array of data in it
func setupData(repository *tickets.Repository, data ...models.Ticket) {
	resetCollection(repository)

	ticketFieldNames, err := reflections.GetFieldNames(&models.Ticket{})
	if err != nil {
		panic(err)
	}

	for _, item := range data {
		for _, fieldName := range ticketFieldNames {
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
	ticketsRepository := makeRepository()

	// Regular example
	ticket := models.Ticket{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(ticketsRepository, ticket)

	ticketFound, err := ticketsRepository.FindByID(event.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if ticketFound == nil {
		t.Error("Expected result to be an ticket object, got nil instead")
	}

	if ticketFound != nil && ticketFound.ID != event.ID {
		t.Error("Expected ticket's id to be", ticket.ID.Hex(), "but is", eventFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	ticket := models.Ticket{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(ticketsRepository, ticket)

	err := ticketsRepository.UpdateByID(ticket.ID.Hex(), map[string]interface{}{
		"Name": "new name",
	})

	if err != nil {
		t.Error(err)
	}

	if name := ticketsRepository.Collection.Columns["Name"][0]; name != "new name" {
		t.Error("Expected ticket name to be 'new name', but instead got", name)
	}
}

func TestDeleteByID(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	ticket := models.Ticket{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(ticketsRepository, ticket)

	err := ticketsRepository.DeleteByID(ticket.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if ids := ticketsRepository.Collection.Columns["ID"]; len(ids) == 0 {
		t.Error("Expected ticket id to have length of 0, but instead got", len(ids))
	}
}

func TestFind(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	tickets := []models.Ticket{
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
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
	setupData(ticketsRepository, tickets...)

	ticketsFound, err := ticketsRepository.Find(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if len(ticketsFound) != 2 {
		t.Error("Expected len(tickets) to be 2, instead got", len(ticketsFound))
	}

	if ticketsFound[0].Description != ticketsFound[1].Description {
		t.Error("Expected tickets' description to equal to each other, instead got",
			ticketsFound[0].Description,
			ticketsFound[1].Description)
	}
}

func TestFindOne(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	tickets := []models.Ticket{
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
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
	setupData(ticketsRepository, tickets...)

	ticketFound, err := ticketsRepository.FindOne(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if ticketFound.Description != "desc3" {
		t.Error("Expected ticket description to equal 'desc3', instead got", ticketFound.Description)
	}
}

func TestUpdateMany(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	tickets := []models.Ticket{
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
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
	setupData(ticketsRepository, tickets...)
}

func TestDeleteMany(t *testing.T) {

	ticketsRepository := makeRepository()

	// Regular example
	tickets := []models.Ticket{
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
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
	setupData(ticketsRepository, tickets...)
}

func TestInsertOne(t *testing.T) {

	ticketsRepository := makeRepository()

	// Regular example
	resetCollection(ticketsRepository)

	newTicket := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "name3",
		Description: "desc3",
		Tags:        []string{"tag3"},
		Type:        "workshop",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	insertedID, err := ticketsRepository.InsertOne(newTicket)

	if err != nil {
		t.Error(err)
	}

	if insertedID != newTicket.ID.Hex() {
		t.Error("Expected inserted id to be ", newTicket.ID.Hex(), "but instead got", insertedID)
	}
}

func TestInsertMany(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	resetCollection(ticketsRepository)

	newTickets := []models.Event{
		models.Ticket{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Ticket{
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

	insertedIDs, err := ticketsRepository.InsertMany(newTickets)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != newTickets[0].ID.Hex() ||
		insertedIDs[1] != newTickets[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", newTickets[0].ID.Hex(), newEvents[1].ID.Hex(), "but instead got", insertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	ticketsRepository := makeRepository()

	// Name is duplicate
	tickets := []models.Ticket{
		models.Ticket{
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
	setupData(ticketsRepository, tickets...)

	isDuplicate := ticketsRepository.IsDuplicate("name2")

	if !isDuplicate {
		t.Error("Expected name to be duplicate")
	}
}
