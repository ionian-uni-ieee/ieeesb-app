package mongo

import (
	"context"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) Find(filter interface{}, skip int64, limit int64) ([]models.Sponsor, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.Sponsor{}),
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

	sponsors := []models.Sponsor{}

	for result.Next(context.Background()) {
		sponsor := models.Sponsor{}

		result.Decode(&sponsor)

		if result.Err() != nil {
			return nil, err
		}

		sponsors = append(sponsors, sponsor)
	}

	return sponsors, nil
}
