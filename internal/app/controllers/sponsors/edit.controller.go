package sponsors

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Edit updates the stored sponsor with the new values on the specific fields that are given in the map
func (c *Controller) Edit(sponsorID string, editMap map[string]interface{}) error {
	if sponsorID == "" {
		return errors.New("SponsorID is empty string")
	}

	if _, err := primitive.ObjectIDFromHex(sponsorID); err != nil {
		return errors.New("SponsorID is not valid ObjectID")
	}

	err := c.repositories.SponsorsRepository.UpdateByID(sponsorID, editMap)

	return err
}
