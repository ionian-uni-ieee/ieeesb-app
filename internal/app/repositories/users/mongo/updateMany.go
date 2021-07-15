package mongo

import (
	"context"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
)

func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	updateBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	result, err := r.collection.UpdateMany(context.Background(), filterBSON, updateBSON)

	return result.UpsertedID.([]string), err
}
