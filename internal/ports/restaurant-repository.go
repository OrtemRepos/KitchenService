package ports

import "github.com/OrtemRepos/KitchenService/internal/domain"

type RestaurantRepository interface {
	Save(domain.Restaurant) error
	Load(id uint) (*domain.Restaurant, error)
}
