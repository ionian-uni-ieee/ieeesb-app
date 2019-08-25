package events

import "github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"

type Repository interface {
	FindByID(eventID string) (*models.Event, error)
	UpdateByID(eventID string, update interface{}) error
	DeleteByID(eventID string) error
	Find(filter interface{}) ([]models.Event, error)
	FindOne(filter interface{}) (*models.Event, error)
	UpdateMany(filter interface{}, update interface{}) ([]string, error)
	DeleteMany(filter interface{}) (int64, error)
	InsertOne(document models.Event) (string, error)
	InsertMany(documents []models.Event) ([]string, error)
	IsDuplicate(name string) bool
}
