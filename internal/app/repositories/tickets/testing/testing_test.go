package testing_test

import (
	"reflect"
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	tickets "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/tickets/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() (*testingDatabase.DatabaseSession, *tickets.Repository) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	ticketsRepository := tickets.MakeRepository(database)

	return database, ticketsRepository
}

var testTicket1 = models.Ticket{
	ID:       primitive.NewObjectID(),
	Email:    "joe@mail.com",
	Fullname: "joe jordinson",
	Message:  "Hello stuff. What's up? This is BJ talking to u.",
	State:    "open",
}

var testTicket2 = models.Ticket{
	ID:       primitive.NewObjectID(),
	Email:    "bill@mail.com",
	Fullname: "Bill Smith",
	Message:  "Is this reality? Or just a fantasy?",
	State:    "open",
}

var testTicket3 = models.Ticket{
	ID:       primitive.NewObjectID(),
	Email:    "john@mail.com",
	Fullname: "John Doe",
	Message:  "Easy come, easy go.",
	State:    "closed",
}

func TestFindByID(t *testing.T) {
	db, ticketsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "tickets", testTicket1)

	gotTicket, gotErr := ticketsRepository.FindByID(testTicket1.ID.Hex())

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotTicket == nil {
		t.Error("Expected result to be an ticket object, got nil instead")
	}

	if gotTicket != nil && gotTicket.ID != testTicket1.ID {
		t.Error("Expected ticket's id to be", testTicket1.ID.Hex(), "but is", gotTicket.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	db, ticketsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "tickets", testTicket1)

	newFullname := "New Fullname"
	gotErr := ticketsRepository.UpdateByID(testTicket1.ID.Hex(), map[string]interface{}{
		"Fullname": newFullname,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	fullnameStored := ticketsRepository.Collection.Columns["Fullname"][0]
	fullnameChanged := fullnameStored != newFullname
	if fullnameChanged {
		t.Error("Expected ticketname to be '"+newFullname+"', but instead got", fullnameStored)
	}
}

func TestDeleteByID(t *testing.T) {
	db, ticketsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "tickets", testTicket1)

	gotErr := ticketsRepository.DeleteByID(testTicket1.ID.Hex())

	if gotErr != nil {
		t.Error(gotErr)
	}

	for key, column := range ticketsRepository.Collection.Columns {
		if len(column) > 0 {
			t.Error("Expected column", key, "to have length of 0, but instead got", len(column))
		}
	}
}

func TestFind(t *testing.T) {
	db, ticketsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "tickets", testTicket1, testTicket1)

	gotTickets, gotErr := ticketsRepository.Find(map[string]interface{}{
		"Email": testTicket1.Email,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	if len(gotTickets) != 2 {
		t.Error("Expected len(tickets) to be 2, instead got", len(gotTickets))
	}

	if gotTickets[0].Email != gotTickets[1].Email {
		t.Error("Expected Email to equal to each other, instead got",
			gotTickets[0].Email,
			gotTickets[1].Email)
	}
}

func TestFindOne(t *testing.T) {
	db, ticketsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "tickets", testTicket1, testTicket2)

	gotTicket, gotErr := ticketsRepository.FindOne(map[string]interface{}{
		"Fullname": testTicket1.Fullname,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotTicket.Fullname != testTicket1.Fullname {
		t.Error("Expected ticketname to equal 'ticketname2', instead got", gotTicket.Fullname)
	}
}

func TestUpdateMany(t *testing.T) {
	// TODO: Not implemented
}

func TestDeleteMany(t *testing.T) {
	// TODO: Not implemented
}

func TestInsertOne(t *testing.T) {

	db, ticketsRepository := makeRepository()

	// Regular example
	testUtils.ResetCollection(db, "tickets")

	gotInsertedID, gotErr := ticketsRepository.InsertOne(testTicket1)

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotInsertedID != testTicket1.ID.Hex() {
		t.Error("Expected inserted id to be ", testTicket1.ID.Hex(), "but instead got", gotInsertedID)
	}

	storedTicket := testUtils.GetInterfaceAtCollectionRow(
		db,
		"tickets",
		reflect.TypeOf(models.Ticket{}),
		0,
	).(models.Ticket)

	if storedTicket.ID.Hex() != gotInsertedID {
		t.Error("Expected stored user's ID to equal gotInsertedID")
	}
}

func TestInsertMany(t *testing.T) {
	db, ticketsRepository := makeRepository()

	// Regular example
	testUtils.ResetCollection(db, "tickets")

	tickets := []models.Ticket{
		testTicket1,
		testTicket2,
		testTicket3,
	}

	gotInsertedIDs, gotErr := ticketsRepository.InsertMany(tickets)

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotInsertedIDs[0] != tickets[0].ID.Hex() ||
		gotInsertedIDs[1] != tickets[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", tickets[0].ID.Hex(), tickets[1].ID.Hex(), "but instead got", gotInsertedIDs)
	}
}
