package ports

import (
	"github.com/OrtemRepos/KitchenService/internal/domain"
	"github.com/jmoiron/sqlx"
)

type TicketRepository interface {
	Save(tx *sqlx.Tx, ticket domain.Ticket) error
	Load(tx *sqlx.Tx, ticketID uint) (*domain.Ticket, error)
}
