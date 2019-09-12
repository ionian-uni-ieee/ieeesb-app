package tickets

import "github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"

func (c *Controller) GetTickets(skip int64, limit int64) ([]models.Ticket, error) {
	tickets, err := c.repositories.TicketsRepository.Find(nil, skip, limit)

	return tickets, err
}
