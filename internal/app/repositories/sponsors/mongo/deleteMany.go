package mongo

import (
	"context"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) DeleteMany(filter interface{}) (int64, error) {
	result, err := r.collection.DeleteMany(context.Background(), filter)

	return result.DeletedCount, err
}

func (r *Repository) InsertOne(document models.Sponsor) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}
