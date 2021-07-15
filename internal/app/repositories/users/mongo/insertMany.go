package mongo

import (
	"context"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) InsertMany(documents []models.User) ([]string, error) {
	var i []interface{}

	for _, user := range documents {
		i = append(i, user)
	}

	result, err := r.collection.InsertMany(context.Background(), i)

	insertedObjectIDs := result.InsertedIDs
	insertedIDs := []string{}

	for _, insertedID := range insertedObjectIDs {
		insertedIDs = append(insertedIDs, insertedID.(primitive.ObjectID).Hex())
	}

	return insertedIDs, err
}
