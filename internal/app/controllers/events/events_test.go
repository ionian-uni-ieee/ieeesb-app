package events_test

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/events"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockEvent = models.Event{
	ID:          primitive.NewObjectID(),
	Name:        "arduino",
	Description: "Learn how to handle Arduino microcontrollers",
	Tags:        []string{"arduino", "hardware"},
	Type:        "seminar",
	Sponsors:    []models.Sponsor{},
	Logo:        models.MediaMeta{},
	Media:       []models.MediaMeta{},
}

var mockEvent2 = models.Event{
	ID:          primitive.NewObjectID(),
	Name:        "C++ from zero to hero",
	Description: "Programming C++ for beginners that want to become heroes",
	Tags:        []string{"language", "programming"},
	Type:        "workshop",
	Sponsors:    []models.Sponsor{},
	Logo:        models.MediaMeta{},
	Media:       []models.MediaMeta{},
}

var mockEvent3 = models.Event{
	ID:          primitive.NewObjectID(),
	Name:        "What is Machine Learning",
	Description: "Learn what is machine learning and how to train your own models like a pro",
	Tags:        []string{"statistics", "machine learning", "programming", "python"},
	Type:        "workshop",
	Sponsors:    []models.Sponsor{},
	Logo:        models.MediaMeta{},
	Media:       []models.MediaMeta{},
}

func makeController() (*testingDatabase.DatabaseSession, *events.Controller) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	reps := repositories.MakeRepositories(database)
	controller := events.MakeController(reps)

	return database, controller
}
