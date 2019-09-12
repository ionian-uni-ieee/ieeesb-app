package tickets

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"

type Repository interface {
	FindByID(ticketID string) (*models.Ticket, error)
	UpdateByID(ticketID string, update interface{}) error
	DeleteByID(ticketID string) error
	Find(filter interface{}, skip int64, limit int64) ([]models.Ticket, error)
	FindOne(filter interface{}) (*models.Ticket, error)
	UpdateMany(filter interface{}, update interface{}) ([]string, error)
	DeleteMany(filter interface{}) (int64, error)
	InsertOne(document models.Ticket) (string, error)
	InsertMany(documents []models.Ticket) ([]string, error)
}
