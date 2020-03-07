package mongo

import (
	"context"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
)

func (r *Repository) FindOne(filter interface{}) (*models.Sponsor, error) {
	result := r.collection.FindOne(context.Background(), filter)

	sponsor := &models.Sponsor{}

	err := result.Decode(sponsor)

	return sponsor, err
}

func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}
