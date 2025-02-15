package domain

import "time"

type EventDomainTicket interface {
	GetID() uint
	GetTicketID() uint
	Type() string
}

type EventTicket struct {
	id       uint
	ticketID uint
}

func (e EventTicket) GetID() uint {
	return e.id
}

func (e EventTicket) GetTicketID() uint {
	return e.ticketID
}

func NewEventTicket(ticketID uint) EventTicket {
	return EventTicket{ticketID: ticketID}
}

type TicketAcceptedEvent struct {
	EventTicket
	ReadyBy    time.Time
	AcceptTime time.Time
}

func NewTicketAcceptedEvent(ticketID uint, readyBy, acceptBy time.Time) *TicketAcceptedEvent {
	return &TicketAcceptedEvent{NewEventTicket(ticketID), readyBy, acceptBy}
}

func (t TicketAcceptedEvent) Type() string {
	return "TicketAcceptedEvent"
}

type TicketCancelCreateEvent struct {
	EventTicket
}

func (t TicketCancelCreateEvent) Type() string {
	return "TicketCancelCreateEvent"
}

func NewTicketCancelCreateEvent(ticketID uint) *TicketCancelCreateEvent {
	return &TicketCancelCreateEvent{NewEventTicket(ticketID)}
}

type TicketCanceledEvent struct {
	EventTicket
	Force bool
}

func (t TicketCanceledEvent) Type() string {
	return "TicketCanceledEvent"
}

func NewTicketCanceledEvent(ticketID uint, force bool) *TicketCanceledEvent {
	return &TicketCanceledEvent{NewEventTicket(ticketID), force}
}

type TicketUndoCanceledEvent struct {
	EventTicket
}

func (t TicketUndoCanceledEvent) Type() string {
	return "TicketUndoCanceledEvent"
}

func NewTicketUndoCanceledEvent(ticketID uint) *TicketUndoCanceledEvent {
	return &TicketUndoCanceledEvent{NewEventTicket(ticketID)}
}

type TicketChangeLineItemEvent struct {
	EventTicket
	LineItems []TicketLineItem
}

func (t TicketChangeLineItemEvent) Type() string {
	return "TicketChangeLineItemEvent"
}

func NewTicketChangeLineItemEvent(ticketID uint, lineItems []TicketLineItem) *TicketChangeLineItemEvent {
	return &TicketChangeLineItemEvent{NewEventTicket(ticketID), lineItems}
}

type TicketCreateEvent struct {
	EventTicket
	RestaurantID  uint
	TicketDetails []TicketLineItem
}

func (t TicketCreateEvent) Type() string {
	return "TicketCreateEvent"
}

func NewTicketCreateEvent(restaurantID, ticketID uint, ticketDetails []TicketLineItem) *TicketCreateEvent {
	return &TicketCreateEvent{NewEventTicket(ticketID), restaurantID, ticketDetails}
}

type TicketPickedUpEvent struct {
	EventTicket
	PickedUpTime time.Time
}

func (t TicketPickedUpEvent) Type() string {
	return "TicketPickedUpEvent"
}

func NewTicketPickedUpEvent(ticketID uint, pickedUpTime time.Time) *TicketPickedUpEvent {
	return &TicketPickedUpEvent{NewEventTicket(ticketID), pickedUpTime}
}

type TicketPreparationCompletedEvent struct {
	EventTicket
	ReadyForPickupTime time.Time
}

func (t TicketPreparationCompletedEvent) Type() string {
	return "TicketPreparationCompletedEvent"
}

func NewTicketPreparationCompletedEvent(ticketID uint, redyForPickupTime time.Time) *TicketPreparationCompletedEvent {
	return &TicketPreparationCompletedEvent{NewEventTicket(ticketID), redyForPickupTime}
}

type TicketPreparingStartedEvent struct {
	EventTicket
	PreparingTime time.Time
}

func (t TicketPreparingStartedEvent) Type() string {
	return "TicketPreparingStartedEvent"
}

func NewTicketPreparingStartedEvent(ticketID uint, preparingTime time.Time) *TicketPreparingStartedEvent {
	return &TicketPreparingStartedEvent{NewEventTicket(ticketID), preparingTime}
}
