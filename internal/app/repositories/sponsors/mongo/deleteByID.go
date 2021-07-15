package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) DeleteByID(sponsorID string) error {
	if sponsorID == "" {
		return errors.New("SponsorID is empty string")
	}

	sponsorObjectID, err := primitive.ObjectIDFromHex(sponsorID)

	if err != nil {
		return errors.New("SponsorID is invalid ObjectID")
	}

	_, err = r.collection.DeleteOne(context.Background(), &bson.M{"_id": sponsorObjectID})

	return err
}
