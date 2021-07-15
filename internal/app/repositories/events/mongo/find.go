package mongo

import (
	"context"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) Find(filter interface{}, skip int64, limit int64) ([]models.Event, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.Event{}),
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

	events := []models.Event{}

	for result.Next(context.Background()) {
		event := models.Event{}

		result.Decode(&event)

		if result.Err() != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
