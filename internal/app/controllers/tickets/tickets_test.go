package tickets_test

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/tickets"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockActiveTicket = models.Ticket{
	ID:       primitive.NewObjectID(),
	Email:    "joe@mail.com",
	Fullname: "Joe Jordinson",
	Message:  "Later bitches",
	State:    "open",
}

var mockClosedTicket = models.Ticket{
	ID:       primitive.NewObjectID(),
	Email:    "bill@mail.com",
	Fullname: "Bill Smith",
	Message:  "Later bitches",
	State:    "closed",
}

func makeController() (*testingDatabase.DatabaseSession, *tickets.Controller) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	reps := repositories.MakeRepositories(database)
	controller := tickets.MakeController(reps)

	return database, controller
}
