package users

import "gitlab.com/gphub/app/internal/app/models"

type Repository interface {
	FindByID(userID string) (*models.User, error)
	FindUserCollectionsByUserID(userID string) ([]models.Collection, error)
	FindUserCollectionByIDs(userID string, collectionID string) (*models.Collection, error)
	UpdateByID(userID string, update interface{}) error
	DeleteByID(userID string) error
	Find(filter interface{}) ([]models.User, error)
	FindOne(filter interface{}) (*models.User, error)
	UpdateMany(filter interface{}, update interface{}) ([]string, error)
	DeleteMany(filter interface{}) (int64, error)
	InsertOne(document models.User) (string, error)
	InsertMany(documents []models.User) ([]string, error)
}
