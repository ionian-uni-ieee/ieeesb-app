package sponsors

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"

type Repository interface {
	FindByID(sponsorID string) (*models.Sponsor, error)
	UpdateByID(sponsorID string, update interface{}) error
	DeleteByID(sponsorID string) error
	Find(filter interface{}) ([]models.Sponsor, error)
	FindOne(filter interface{}) (*models.Sponsor, error)
	UpdateMany(filter interface{}, update interface{}) ([]string, error)
	DeleteMany(filter interface{}) (int64, error)
	InsertOne(document models.Sponsor) (string, error)
	InsertMany(documents []models.Sponsor) ([]string, error)
	IsDuplicate(name string) bool
}
