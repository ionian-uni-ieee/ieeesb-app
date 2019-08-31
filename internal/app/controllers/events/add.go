package events

import (
	"errors"
	"strconv"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
)

// Add adds an event to the repository
func (c *Controller) Add(newEvent models.Event) (string, error) {
	if newEvent.Name == "" {
		return "", errors.New("Name is empty string")
	}

	if newEvent.Description == "" {
		return "", errors.New("Description is empty string")
	}

	for index, tag := range newEvent.Tags {
		if tag == "" {
			return "", errors.New("tags[" + strconv.Itoa(index) + "] is empty string")
		}
	}

	if newEvent.Type == "" {
		return "", errors.New("Type is empty string")
	}

	isDuplicate := c.repositories.EventsRepository.IsDuplicate(newEvent.Name)

	if isDuplicate {
		return "", errors.New("An event already exists with the same name")
	}

	eventHexID, err := c.repositories.EventsRepository.InsertOne(newEvent)

	return eventHexID, err
}
