package sponsors

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) Delete(sponsorID string) error {
	if sponsorID == "" {
		return errors.New("SponsorID is empty string")
	}

	if _, err := primitive.ObjectIDFromHex(sponsorID); err != nil {
		return errors.New("SponsorID is invalid ObjectID")
	}

	err := c.repositories.SponsorsRepository.DeleteByID(sponsorID)

	return err
}
