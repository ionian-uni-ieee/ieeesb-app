package main

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/mongo"
)

func main() {
	app := &app.Application{
		DatabasePort: "27017",
		DatabaseHost: "mongodb://localhost",
		DatabaseName: "test",
	}

	database := mongo.MakeDatabaseDriver()

	app.Initialize(database)
}
