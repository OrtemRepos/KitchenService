package ports

import (
	"github.com/OrtemRepos/KitchenService/internal/domain"
	"github.com/google/uuid"
)

type EventStore interface {
	Save(aggregateID uuid.UUID, events []domain.EventDomainTicket, expectedVersion int) error
	Load(aggregateID uuid.UUID) domain.EventDomainTicket
}