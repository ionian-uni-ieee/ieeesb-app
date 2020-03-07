package mongo

import (
	"context"
	"errors"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) FindByID(userID string) (*models.User, error) {
	if userID == "" {
		return nil, errors.New("UserID is empty string")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, errors.New("UserID is invalid ObjectID")
	}

	result := r.collection.FindOne(context.Background(), &bson.M{"_id": userObjectID})

	user := &models.User{}

	err = result.Decode(user)

	return user, err
}
