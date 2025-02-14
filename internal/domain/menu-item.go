package domain

import "github.com/google/uuid"

type Money float64

type MenuItem struct {
	ID    uuid.UUID
	Name  string
	Price Money
}

func NewMenuItem(id uuid.UUID, name string, price Money) *MenuItem {
	return &MenuItem{id, name, price}
}
