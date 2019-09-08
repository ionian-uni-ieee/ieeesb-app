package media_test

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/media"
	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories"
)

func makeController() (*testingDatabase.DatabaseSession, *media.Controller) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	reps := repositories.MakeRepositories(database)
	controller := media.MakeController(reps)

	return database, controller
}
