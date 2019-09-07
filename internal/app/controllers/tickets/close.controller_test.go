package tickets_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
)

func TestClose(t *testing.T) {
	// Setup
	db, ticketController := makeController()

	t.Run("Closes ticket", func(t *testing.T) {
		testUtils.SetupData(db, "tickets", mockActiveTicket)

		err := ticketController.Close(mockActiveTicket.ID.Hex())

		if err != nil {
			t.Error(err)
		}

		firstRowTicket := testUtils.GetInterfaceAtCollectionRow(
			db,
			"tickets",
			reflect.TypeOf(models.Ticket{}),
			0,
		).(models.Ticket)

		if firstRowTicket.State != "closed" {
			t.Error("Expected ticket state to be 'closed', but it's '" + firstRowTicket.State + "'")
		}
	})

	t.Run("Ticket already closed", func(t *testing.T) {
		testUtils.SetupData(db, "tickets", mockClosedTicket)

		err := ticketController.Close(mockClosedTicket.ID.Hex())

		expectedError := "Ticket is already closed"
		if err.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' error but got '" + err.Error() + "'")
		}
	})

	t.Run("No such ticket", func(t *testing.T) {
		testUtils.ResetCollection(db, "tickets")

		err := ticketController.Close(mockClosedTicket.ID.Hex())

		expectedError := "No ticketID was found with this ObjectID"
		if err.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' error but got '" + err.Error() + "'")
		}
	})
}
