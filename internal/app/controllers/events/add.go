package events

import (
	"errors"
	"strconv"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
)

func (c *Controller) Add(name string, description string, tags []string, eType string) (string, error) {
	if name == "" {
		return "", errors.New("Name is empty string")
	}

	if description == "" {
		return "", errors.New("Description is empty string")
	}

	for index, tag := range tags {
		if tag == "" {
			return "", errors.New("tags[" + strconv.Itoa(index) + "] is empty string")
		}
	}

	if eType == "" {
		return "", errors.New("Type is empty string")
	}

	isDuplicate := c.repositories.EventsRepository.IsDuplicate(name)

	if isDuplicate {
		return "", errors.New("An event already exists with the same name")
	}

	newEvent := models.Event{
		Name:        name,
		Description: description,
		Tags:        tags,
		Type:        eType,
	}

	eventHexID, err := c.repositories.EventsRepository.InsertOne(newEvent)

	return eventHexID, err
}
