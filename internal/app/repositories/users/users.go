package users

import "github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"

type Repository interface {
	FindByID(userID string) (*models.User, error)
	UpdateByID(userID string, update interface{}) error
	DeleteByID(userID string) error
	Find(filter interface{}) ([]models.User, error)
	FindOne(filter interface{}) (*models.User, error)
	UpdateMany(filter interface{}, update interface{}) ([]string, error)
	DeleteMany(filter interface{}) (int64, error)
	InsertOne(document models.User) (string, error)
	InsertMany(documents []models.User) ([]string, error)
}
