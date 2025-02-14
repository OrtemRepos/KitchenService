package service

import (
	"github.com/OrtemRepos/KitchenService/internal/domain"
	"github.com/OrtemRepos/KitchenService/internal/ports"
)

type KitchenService struct {
	ticketRepository           ports.TicketRepository
	ticketDomainEventPublisher ports.DomainEventPublisher
	restaurantRepository       ports.RestaurantRepository
}

func (ks *KitchenService) CreateMenu(id uint, menu []domain.MenuItem) error {
	restaurant := domain.Restaurant{ID: id, Menu: menu}
	return ks.restaurantRepository.Save(restaurant)
}
