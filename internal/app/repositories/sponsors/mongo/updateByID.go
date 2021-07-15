package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) UpdateByID(sponsorID string, update interface{}) error {
	if sponsorID == "" {
		return errors.New("SponsorID is empty string")
	}

	sponsorObjectID, err := primitive.ObjectIDFromHex(sponsorID)

	if err != nil {
		return errors.New("SponsorID is invalid ObjectID")
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": sponsorObjectID}, &bson.M{"$set": update})

	return err
}
