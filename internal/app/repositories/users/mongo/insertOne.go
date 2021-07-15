package mongo

import (
	"context"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) InsertOne(document models.User) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}
