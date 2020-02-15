package main

import (
	"github.com/ionian-uni-ieee/ieeesb-app/cmd/app"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database/mongo"
)

func main() {
	database := mongo.MakeDatabaseDriver()
	app.Initialize(database)
}
