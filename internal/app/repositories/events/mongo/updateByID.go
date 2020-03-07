package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) UpdateByID(eventID string, update interface{}) error {
	if eventID == "" {
		return errors.New("EventID is empty string")
	}

	eventObjectID, err := primitive.ObjectIDFromHex(eventID)

	if err != nil {
		return errors.New("EventID is invalid ObjectID")
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": eventObjectID}, &bson.M{"$set": update})

	return err
}
