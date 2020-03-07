package mongo

import (
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database"
	mongod "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	database   database.Driver
	collection *mongod.Collection
}

func MakeRepository(database database.Driver) *Repository {
	collection := database.GetCollection("events").(*mongod.Collection)
	return &Repository{database, collection}
}
