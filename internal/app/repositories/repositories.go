package repositories

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/events"
	eventsMongo "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/events/mongo"
	eventsTesting "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/events/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/sessions"
	sessionsMongo "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/sessions/mongo"
	sessionsTesting "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/sessions/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/sponsors"
	sponsorsMongo "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/sponsors/mongo"
	sponsorsTesting "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/sponsors/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/tickets"
	ticketsMongo "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/tickets/mongo"
	ticketsTesting "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/tickets/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/users"
	usersMongo "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/users/mongo"
	usersTesting "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/users/testing"
)

// Repositories stores all the model repositories
// with which you can request data exchanges or modifications
// for a specific model
type Repositories struct {
	database database.Driver
	Users    users.Repository
	Tickets  tickets.Repository
	Events   events.Repository
	Sponsors sponsors.Repository
	Sessions sessions.Repository
}

// MakeRepositories is a factory of Repositories
// which creates a set of repositories
func MakeRepositories(database database.Driver) (repositories *Repositories) {
	repositories = &Repositories{}
	repositories.database = database
	if database.GetDatabaseType() == "mongo" {
		repositories.Users = usersMongo.MakeRepository(database)
		repositories.Events = eventsMongo.MakeRepository(database)
		repositories.Sponsors = sponsorsMongo.MakeRepository(database)
		repositories.Tickets = ticketsMongo.MakeRepository(database)
		repositories.Sessions = sessionsMongo.MakeRepository(database)
	} else if database.GetDatabaseType() == "testing" {
		repositories.Users = usersTesting.MakeRepository(database)
		repositories.Events = eventsTesting.MakeRepository(database)
		repositories.Sponsors = sponsorsTesting.MakeRepository(database)
		repositories.Tickets = ticketsTesting.MakeRepository(database)
		repositories.Sessions = sessionsTesting.MakeRepository(database)
	} else {
		panic("The database type '" + database.GetDatabaseType() + "' is not supported")
	}

	return repositories
}
