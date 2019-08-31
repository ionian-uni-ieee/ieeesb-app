package events_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAdd(t *testing.T) {
	// Setup
	db, eventsController := makeController()

	// Normal
	event := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "Arduino",
		Description: "desc",
		Tags:        []string{"arduino", "hardware"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	eventsController.Add(event)

	if !isEventEqualToDbRow(db, event, 0) {
		t.Error("Expected event to have been added in database")
	}
}
