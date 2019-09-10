package tickets

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) Close(ticketID string) error {
	if ticketID == "" {
		return errors.New("TicketID is empty string")
	}

	if _, err := primitive.ObjectIDFromHex(ticketID); err != nil {
		return errors.New("TicketID is not valid ObjectID")
	}

	ticket, err := c.repositories.TicketsRepository.FindByID(ticketID)

	if err != nil {
		return err
	}

	if ticket == nil {
		return errors.New("No ticketID was found with this ObjectID")
	}

	if ticket.State == "closed" {
		return errors.New("Ticket is already closed")
	}

	ticket.State = "closed"

	err = c.repositories.TicketsRepository.UpdateByID(ticketID, *ticket)

	return err
}
