package ports

import (
	"github.com/OrtemRepos/KitchenService/internal/domain"
	"github.com/google/uuid"
)

type TicketRepository interface {
	Save(domain.Ticket) error
	Load(ticketID uuid.UUID) (*domain.Ticket, error)	
}