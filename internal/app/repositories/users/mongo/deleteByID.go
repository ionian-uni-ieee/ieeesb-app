package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) DeleteByID(userID string) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return errors.New("UserID is invalid ObjectID")
	}

	user := r.collection.FindOne(context.Background(), &bson.M{"_id": userObjectID})

	if user == nil {
		return errors.New("No user with this ID exists")
	}

	_, err = r.collection.DeleteOne(context.Background(), &bson.M{"_id": userObjectID})

	return err
}
