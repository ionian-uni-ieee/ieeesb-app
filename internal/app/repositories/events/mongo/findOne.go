package mongo

import (
	"context"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
)

func (r *Repository) FindOne(filter interface{}) (*models.Event, error) {
	result := r.collection.FindOne(context.Background(), filter)

	event := &models.Event{}

	err := result.Decode(event)

	return event, err
}
