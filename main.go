package main

import (
	"gitlab.com/ionio-uni-ieee/ieee-webapp/internal/app"
)

func main() {
	app := 	&app.Application{
		DatabasePort: "27017",
		DatabaseHost: "mongodb://localhost",
		DatabaseName: "test",
	}

	database := mongo.MakeDatabaseDriver()

	app.Initialize(database)
}