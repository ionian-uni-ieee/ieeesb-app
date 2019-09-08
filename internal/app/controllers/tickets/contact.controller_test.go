package tickets_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
)

func TestContact(t *testing.T) {
	// Setup
	db, ticketController := makeController()

	t.Run("Creates a ticket", func(t *testing.T) {
		testUtils.ResetCollection(db, "tickets")
		gotTicketID, gotErr := ticketController.Contact(
			mockActiveTicket.Email,
			mockActiveTicket.Fullname,
			mockActiveTicket.Message,
		)

		if gotErr != nil {
			t.Error(gotErr)
		}

		firstTicket := testUtils.GetInterfaceAtCollectionRow(
			db,
			"tickets",
			reflect.TypeOf(models.Ticket{}),
			0,
		).(models.Ticket)

		if firstTicket.ID.Hex() != gotTicketID {
			t.Error("Expected first ticket's ID to equal the controller's returned ID, instead it's " + firstTicket.ID.Hex())
		}

		if firstTicket.State != "open" {
			t.Error("Expected ticket state to be 'open', but it's '" + firstTicket.State + "'")
		}
	})
}
