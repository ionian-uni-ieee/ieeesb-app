package sponsors

import (
	"errors"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
)

// Add inserts a sponsor in the repository
func (c *Controller) Add(sponsor models.Sponsor) (string, error) {
	if sponsor.Name == "" {
		return "", errors.New("Name is empty string")
	}

	insertedID, err := c.repositories.SponsorsRepository.InsertOne(sponsor)

	return insertedID, err
}
