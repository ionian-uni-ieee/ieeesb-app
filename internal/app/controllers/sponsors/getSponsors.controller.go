package sponsors

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"

func (c *Controller) GetSponsors(skip int64, limit int64) ([]models.Sponsor, error) {
	sponsors, err := c.repositories.SponsorsRepository.Find(nil, skip, limit)

	return sponsors, err
}
