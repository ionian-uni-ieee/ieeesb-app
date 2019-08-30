package repositories

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/events"
	eventsMongo "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/events/mongo"
	eventsTesting "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/events/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sessions"
	sessionsMongo "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sessions/mongo"
	sessionsTesting "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sessions/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sponsors"
	sponsorsMongo "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sponsors/mongo"
	sponsorsTesting "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sponsors/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/tickets"
	ticketsMongo "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/tickets/mongo"
	ticketsTesting "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/tickets/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/users"
	usersMongo "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/users/mongo"
	usersTesting "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/users/testing"
)

// Repositories stores all the model repositories
// with which you can request data exchanges or modifications
// for a specific model
type Repositories struct {
	database           database.Driver
	UsersRepository    users.Repository
	TicketsRepository  tickets.Repository
	EventsRepository   events.Repository
	SponsorsRepository sponsors.Repository
	SessionsRepository sessions.Repository
}

// MakeNewRepositories is a factory of Repositories
// which creates a set of repositories
func MakeNewRepositories(database database.Driver) (repositories *Repositories) {
	repositories = &Repositories{}
	repositories.database = database
	if database.GetDatabaseType() == "mongo" {
		repositories.UsersRepository = usersMongo.MakeRepository(database)
		repositories.EventsRepository = eventsMongo.MakeRepository(database)
		repositories.SponsorsRepository = sponsorsMongo.MakeRepository(database)
		repositories.TicketsRepository = ticketsMongo.MakeRepository(database)
		repositories.SessionsRepository = sessionsMongo.MakeRepository(database)
	} else if database.GetDatabaseType() == "testing" {
		repositories.UsersRepository = usersTesting.MakeRepository(database)
		repositories.EventsRepository = eventsTesting.MakeRepository(database)
		repositories.SponsorsRepository = sponsorsTesting.MakeRepository(database)
		repositories.TicketsRepository = ticketsTesting.MakeRepository(database)
		repositories.SessionsRepository = sessionsTesting.MakeRepository(database)
	} else {
		panic("The database type '" + database.GetDatabaseType() + "' is not supported")
	}

	return repositories
}
