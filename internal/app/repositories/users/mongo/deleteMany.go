package mongo

import (
	"context"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
)

func (r *Repository) DeleteMany(filter interface{}) (int64, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return -1, err
	}

	result, err := r.collection.DeleteMany(context.Background(), filterBSON)

	return result.DeletedCount, err
}
