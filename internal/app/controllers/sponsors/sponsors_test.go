package sponsors_test

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/sponsors"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockSponsor = models.Sponsor{
	ID:     primitive.NewObjectID(),
	Name:   "Microsoft",
	Emails: []string{"info@microsoft.com"},
	Phones: []string{"+1 234 343 324"},
	Logo:   models.MediaMeta{},
}

var mockSponsor2 = models.Sponsor{
	ID:     primitive.NewObjectID(),
	Name:   "Netflix",
	Emails: []string{"info@netflix.com"},
	Phones: []string{"+1 562 623 484"},
	Logo:   models.MediaMeta{},
}

func makeController() (*testingDatabase.DatabaseSession, *sponsors.Controller) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	reps := repositories.MakeRepositories(database)
	controller := sponsors.MakeController(reps)

	return database, controller
}
