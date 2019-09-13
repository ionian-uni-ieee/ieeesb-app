package events

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"

func (c *Controller) GetEvents(skip int64, limit int64) ([]models.Event, error) {
	events, err := c.repositories.EventsRepository.Find(nil, skip, limit)

	return events, err
}
