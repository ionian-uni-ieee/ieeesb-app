package events_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
)

func TestGetEvents(t *testing.T) {
	// Setup
	db, eventsController := makeController()

	t.Run("Should return an array of 2 events", func(t *testing.T) {
		testUtils.SetupData(db, "events", mockEvent, mockEvent2)

		gotEvents, err := eventsController.GetEvents(0, 2)

		if err != nil {
			t.Error(err)
			t.SkipNow()
		}

		if len(gotEvents) != 2 {
			t.Error("Expected events array with length of 2, but got", len(gotEvents))
			t.SkipNow()
		}

		gotValidData := gotEvents[0].ID == mockEvent.ID ||
			gotEvents[1].ID == mockEvent2.ID
		if !gotValidData {
			t.Error("Expected returned data to have the same IDs with stored data")
		}
	})
}
