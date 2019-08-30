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
		ID:       primitive.NewObjectID(),
		Email:    "email",
		Fullname: "fullname",
		Message:  "message",
		State:    "active",
	}
	setupData(ticketsRepository, ticket)

	ticketFound, err := ticketsRepository.FindByID(ticket.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if ticketFound == nil {
		t.Error("Expected result to be an ticket object, got nil instead")
	}

	if ticketFound != nil && ticketFound.ID != ticket.ID {
		t.Error("Expected ticket's id to be", ticket.ID.Hex(), "but is", ticketFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	ticket := models.Ticket{
		ID:       primitive.NewObjectID(),
		Email:    "email",
		Fullname: "fullname",
		Message:  "message",
		State:    "active",
	}
	setupData(ticketsRepository, ticket)

	err := ticketsRepository.UpdateByID(ticket.ID.Hex(), map[string]interface{}{
		"Fullname": "new name",
	})

	if err != nil {
		t.Error(err)
	}

	if name := ticketsRepository.Collection.Columns["Fullname"][0]; name != "new name" {
		t.Error("Expected fullname to be 'new name', but instead got", name)
	}
}

func TestDeleteByID(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	ticket := models.Ticket{
		ID:       primitive.NewObjectID(),
		Email:    "email",
		Fullname: "fullname",
		Message:  "message",
		State:    "active",
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
			ID:       primitive.NewObjectID(),
			Email:    "email",
			Fullname: "fullname",
			Message:  "message",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email2",
			Fullname: "fullname2",
			Message:  "message2",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email2",
			Fullname: "fullname3",
			Message:  "message3",
			State:    "active",
		},
	}
	setupData(ticketsRepository, tickets...)

	ticketsFound, err := ticketsRepository.Find(map[string]interface{}{
		"Email": "email2",
	})

	if err != nil {
		t.Error(err)
	}

	if len(ticketsFound) != 2 {
		t.Error("Expected len(tickets) to be 2, instead got", len(ticketsFound))
	}

	if ticketsFound[0].Email != ticketsFound[1].Email {
		t.Error("Expected email to equal to each other, instead got",
			ticketsFound[0].Email,
			ticketsFound[1].Email)
	}
}

func TestFindOne(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	tickets := []models.Ticket{
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email",
			Fullname: "fullname",
			Message:  "message",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email2",
			Fullname: "fullname2",
			Message:  "message2",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email3",
			Fullname: "fullname3",
			Message:  "message3",
			State:    "active",
		},
	}
	setupData(ticketsRepository, tickets...)

	ticketFound, err := ticketsRepository.FindOne(map[string]interface{}{
		"Email": "email2",
	})

	if err != nil {
		t.Error(err)
	}

	if ticketFound.Email != "email2" {
		t.Error("Expected email to equal 'email2', instead got", ticketFound.Email)
	}
}

func TestUpdateMany(t *testing.T) {
	ticketsRepository := makeRepository()

	// Regular example
	tickets := []models.Ticket{
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email",
			Fullname: "fullname",
			Message:  "message",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email2",
			Fullname: "fullname2",
			Message:  "message2",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email3",
			Fullname: "fullname3",
			Message:  "message3",
			State:    "active",
		},
	}
	setupData(ticketsRepository, tickets...)
}

func TestDeleteMany(t *testing.T) {

	ticketsRepository := makeRepository()

	// Regular example
	tickets := []models.Ticket{
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email",
			Fullname: "fullname",
			Message:  "message",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email2",
			Fullname: "fullname2",
			Message:  "message2",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email3",
			Fullname: "fullname3",
			Message:  "message3",
			State:    "active",
		},
	}
	setupData(ticketsRepository, tickets...)
}

func TestInsertOne(t *testing.T) {

	ticketsRepository := makeRepository()

	// Regular example
	resetCollection(ticketsRepository)

	newTicket := models.Ticket{
		ID:       primitive.NewObjectID(),
		Email:    "email3",
		Fullname: "fullname3",
		Message:  "message3",
		State:    "active",
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

	newTickets := []models.Ticket{
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email",
			Fullname: "fullname",
			Message:  "message",
			State:    "active",
		},
		models.Ticket{
			ID:       primitive.NewObjectID(),
			Email:    "email2",
			Fullname: "fullname2",
			Message:  "message2",
			State:    "active",
		},
	}

	insertedIDs, err := ticketsRepository.InsertMany(newTickets)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != newTickets[0].ID.Hex() ||
		insertedIDs[1] != newTickets[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", newTickets[0].ID.Hex(), newTickets[1].ID.Hex(), "but instead got", insertedIDs)
	}
}
