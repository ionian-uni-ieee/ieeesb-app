package repositories

import (
	"gitlab.com/gphub/app/internal/app/drivers/database"
	"gitlab.com/gphub/app/internal/app/repositories/sessions"
	"gitlab.com/gphub/app/internal/app/repositories/users"
)

type Repositories struct {
	database           database.Driver
	UsersRepository    users.Repository
	TicketsRepository  tickets.Repository
	EventsRepository   events.Repository
	SponsorsRepository sponsors.Repository
	SessionsRepository sessions.Repository
}

func MakeNewRepositories(database database.Driver) (repositories *Repositories) {
	repositories = &Repositories{}
	repositories.database = database
	if database.GetDatabaseType() == "mongo" {
		repositories.UsersRepository = users.MakeMongoRepository(database)
		repositories.EventsRepository = events.MakeMongoRepository(database)
		repositories.SponsorsRepository = sponsors.MakeMongoRepository(database)
		repositories.TicketsRepository = tickets.MakeMongoRepository(database)
		repositories.SessionsRepository = sessions.MakeMongoRepository(database)
	} else {
		panic("The database type '" + database.GetDatabaseType() + "' is not supported")
	}

	return repositories
}
