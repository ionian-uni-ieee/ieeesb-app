package mongo

import (
	"context"
	"errors"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) UpdateByID(userID string, update interface{}) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return errors.New("UserID is invalid ObjectID")
	}

	user := r.collection.FindOne(context.Background(), &bson.M{"_id": userObjectID})

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("No user with this ID exists")
	}

	updateBSON, err := reflections.ConvertFieldNamesToTagNames(
		update.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": userObjectID}, &bson.M{"$set": updateBSON})

	return err
}
