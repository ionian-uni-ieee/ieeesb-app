package mongo

import (
	"context"
	"errors"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) FindByID(eventID string) (*models.Event, error) {
	if eventID == "" {
		return nil, errors.New("EventID is empty string")
	}

	eventObjectID, err := primitive.ObjectIDFromHex(eventID)

	if err != nil {
		return nil, errors.New("EventID is invalid ObjectID")
	}

	result := r.collection.FindOne(context.Background(), &bson.M{"_id": eventObjectID})

	event := &models.Event{}

	err = result.Decode(event)

	return event, err
}
