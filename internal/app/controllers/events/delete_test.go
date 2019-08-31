package events_test

import (
	"testing"
)

func TestDelete(t *testing.T) {
	// Setup
	db, eventsController := makeController()

	// Normal
	setupData(db, "events", mockEvent)

	err := eventsController.Delete(mockEvent.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if isEventEqualToDbRow(db, mockEvent, 0) {
		t.Error("Expected event to have been deleted from the database")
	}
}
