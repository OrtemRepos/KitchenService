package domain

import "time"

type EventDomainTicket interface {
	GetID() uint
	GetTicketID() uint
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

type TicketCancelCreateEvent struct {
	EventTicket
}

func NewTicketCancelCreateEvent(ticketID uint) *TicketCancelCreateEvent {
	return &TicketCancelCreateEvent{NewEventTicket(ticketID)}
}

type TicketCanceledEvent struct {
	EventTicket
	Force bool
}

func NewTicketCanceledEvent(ticketID uint, force bool) *TicketCanceledEvent {
	return &TicketCanceledEvent{NewEventTicket(ticketID), force}
}

type TicketUndoCanceledEvent EventTicket

func NewTicketUndoCanceledEvent(ticketID uint) *TicketUndoCanceledEvent {
	e := TicketUndoCanceledEvent(NewEventTicket(ticketID))
	return &e
}

type TicketChangeLineItemEvent struct {
	EventTicket
	LineItems []TicketLineItem
}

func NewTicketChangeLineItemEvent(ticketID uint, lineItems []TicketLineItem) *TicketChangeLineItemEvent {
	return &TicketChangeLineItemEvent{NewEventTicket(ticketID), lineItems}
}

type TicketCreateEvent struct {
	EventTicket
	TicketDetails []TicketLineItem
}

func NewTicketCreateEvent(ticketID uint, ticketDetails []TicketLineItem) *TicketCreateEvent {
	return &TicketCreateEvent{NewEventTicket(ticketID), ticketDetails}
}

type TicketPickedUpEvent struct {
	EventTicket
	PickedUpTime time.Time
}

func NewTicketPickedUpEvent(ticketID uint, pickedUpTime time.Time) *TicketPickedUpEvent {
	return &TicketPickedUpEvent{NewEventTicket(ticketID), pickedUpTime}
}

type TicketPreparationCompletedEvent struct {
	EventTicket
	ReadyForPickupTime time.Time
}

func NewTicketPreparationCompletedEvent(ticketID uint, redyForPickupTime time.Time) *TicketPreparationCompletedEvent {
	return &TicketPreparationCompletedEvent{NewEventTicket(ticketID), redyForPickupTime}
}

type TicketPreparingStartedEvent struct {
	EventTicket
	PreparingTime time.Time
}

func NewTicketPreparingStartedEvent(ticketID uint, preparingTime time.Time) *TicketPreparingStartedEvent {
	return &TicketPreparingStartedEvent{NewEventTicket(ticketID), preparingTime}
}
