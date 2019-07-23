package sessions

import "github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"

type Repository interface {
	FindByID(sessionID string) (*models.Session, error)
	UpdateByID(sessionID string, update interface{}) error
	DeleteByID(sessionID string) error
	Find(filter interface{}) ([]models.Session, error)
	FindOne(filter interface{}) (*models.Session, error)
	UpdateMany(filter interface{}, update interface{}) ([]string, error)
	DeleteMany(filter interface{}) (int64, error)
	InsertOne(document models.Session) (string, error)
	InsertMany(documents []models.Session) ([]string, error)
}
