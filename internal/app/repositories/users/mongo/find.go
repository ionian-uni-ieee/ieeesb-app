package mongo

import (
	"context"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) Find(filter interface{}, skip int64, limit int64) ([]models.User, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	if skip < 0 {
		skip = 0
	}
	if limit <= 0 {
		skip = 12
	}

	result, err := r.collection.Find(
		context.Background(),
		filterBSON,
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		})
	defer result.Close(context.Background())

	users := []models.User{}

	for result.Next(context.Background()) {
		user := models.User{}

		result.Decode(&user)

		if result.Err() != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
