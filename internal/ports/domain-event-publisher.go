package ports

import "github.com/OrtemRepos/KitchenService/internal/domain"

type DomainEventPublisher interface {
	Publish(domain.EventDomainTicket) error
}