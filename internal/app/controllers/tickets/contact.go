package tickets

import (
	"errors"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
)

// Contact creates a new ticket
func (c *Controller) Contact(email string, fullname string, message string) (string, error) {
	if email == "" {
		return "", errors.New("Email is empty string")
	}

	if fullname == "" {
		return "", errors.New("Fullname is empty string")
	}

	if message == "" {
		return "", errors.New("Message is empty string")
	}

	ticket := models.Ticket{
		Email:    email,
		Fullname: fullname,
		Message:  message,
		State:    "open",
	}

	ticketID, err := c.repositories.TicketsRepository.InsertOne(ticket)

	return ticketID, err
}
