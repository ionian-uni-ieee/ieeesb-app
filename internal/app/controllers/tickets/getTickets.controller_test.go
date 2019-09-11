package tickets_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
)

func TestGetTickets(t *testing.T) {
	// Setup
	db, ticketsController := makeController()

	t.Run("Should return an array of 2 tickets", func(t *testing.T) {
		testUtils.SetupData(db, "tickets", mockActiveTicket, mockClosedTicket)

		gotTickets, err := ticketsController.GetTickets(0, 2)

		if err != nil {
			t.Error(err)
			t.SkipNow()
		}

		if len(gotTickets) != 2 {
			t.Error("Expected tickets array with length of 2, but got", len(gotTickets))
			t.SkipNow()
		}

		gotValidData := gotTickets[0].ID == mockActiveTicket.ID ||
			gotTickets[1].ID == mockClosedTicket.ID
		if !gotValidData {
			t.Error("Expected returned data to have the same IDs with stored data")
		}
	})
}
