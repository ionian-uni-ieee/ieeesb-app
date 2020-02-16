package users

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"

// Repository interface contains all the necessary
// functions for a user's repository
type Repository interface {
	FindByID(userID string) (*models.User, error)
	UpdateByID(userID string, update interface{}) error
	DeleteByID(userID string) error
	Find(filter interface{}, skip int64, limit int64) ([]models.User, error)
	FindOne(filter interface{}) (*models.User, error)
	UpdateMany(filter interface{}, update interface{}) ([]string, error)
	DeleteMany(filter interface{}) (int64, error)
	InsertOne(document models.User) (string, error)
	InsertMany(documents []models.User) ([]string, error)
	Exists(email string, username string, fullname string) bool
}
