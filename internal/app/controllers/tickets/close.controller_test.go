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

	t.Run("Should change ticket's state to closed", func(t *testing.T) {
		testUtils.SetupData(db, "tickets", mockActiveTicket)

		gotErr := ticketController.Close(mockActiveTicket.ID.Hex())

		if gotErr != nil {
			t.Error(gotErr)
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

	t.Run("Should return ticket is already closed error", func(t *testing.T) {
		testUtils.SetupData(db, "tickets", mockClosedTicket)

		gotErr := ticketController.Close(mockClosedTicket.ID.Hex())

		expectedError := "Ticket is already closed"
		if gotErr.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' gotError but got '" + gotErr.Error() + "'")
		}
	})

	t.Run("Should return ticket was not found error", func(t *testing.T) {
		testUtils.ResetCollection(db, "tickets")

		gotErr := ticketController.Close(mockClosedTicket.ID.Hex())

		expectedError := "No ticketID was found with this ObjectID"
		if gotErr.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' gotError but got '" + gotErr.Error() + "'")
		}
	})
}
