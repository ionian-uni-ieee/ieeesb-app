package main

import (
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/mongo"
)

func main() {
	app := &app.Application{}

	database := mongo.MakeDatabaseDriver()

	app.Initialize(database)
}
