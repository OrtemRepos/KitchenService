package domain

import "github.com/google/uuid"

type TicketLineItem struct {
	Quantity   int
	MenuItemID uuid.UUID
	Name       string
}

func NewTicketItem(quantity int, menuItemID uuid.UUID, name string) *TicketLineItem {
	return &TicketLineItem{quantity, menuItemID, name}
}
