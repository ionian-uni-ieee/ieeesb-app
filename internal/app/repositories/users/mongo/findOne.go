package mongo

import (
	"context"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
)

func (r *Repository) FindOne(filter interface{}) (*models.User, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	result := r.collection.FindOne(context.Background(), filterBSON)

	user := &models.User{}

	err = result.Decode(user)

	return user, err
}
