package mongo

import (
	"context"
	"errors"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) FindByID(sponsorID string) (*models.Sponsor, error) {
	if sponsorID == "" {
		return nil, errors.New("SponsorID is empty string")
	}

	sponsorObjectID, err := primitive.ObjectIDFromHex(sponsorID)

	if err != nil {
		return nil, errors.New("SponsorID is invalid ObjectID")
	}

	result := r.collection.FindOne(context.Background(), &bson.M{"_id": sponsorObjectID})

	sponsor := &models.Sponsor{}

	err = result.Decode(sponsor)

	return sponsor, err
}
