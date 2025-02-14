package domain

import (
	"fmt"
	"time"
)

type state interface {
	confirmCreate() error
	cancelCreate() error
	accept(readyBy, acceptTime time.Time) error
	preparing(preparingTime time.Time) error
	readyForPickUp(readyForPickUp time.Time) error
	pickedUp(pickedUp time.Time) error
	changeLineItem(lineItems []TicketLineItem) error
	cancel() error
	confirmCancel() error
	undoCancel() error
}

type defaultState struct {
	aggregateTicket *Ticket
}

func newDefaultState(aggregate *Ticket) *defaultState {
	s := defaultState{aggregate}
	return &s
}

func (s defaultState) confirmCreate() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) cancelCreate() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) accept(redyBy, acceptTime time.Time) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) preparing(preparingTime time.Time) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) readyForPickUp(readyForPickUp time.Time) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) pickedUp(pickedUp time.Time) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) changeLineItem(lineItems []TicketLineItem) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) cancel() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) confirmCancel() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

func (s defaultState) undoCancel() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.state)
}

type createPendingState struct {
	*defaultState
}

func (s createPendingState) confirmCreate() error {
	s.aggregateTicket.previouState = s.aggregateTicket.state
	s.aggregateTicket.state = s.aggregateTicket.awaitingAcceptanceState
	return nil
}

func (s createPendingState) cancelCreate() error {
	s.aggregateTicket.previouState = s.aggregateTicket.state
	s.aggregateTicket.state = s.aggregateTicket.cancelPendingState
	return nil
}

type awaitingAcceptanceState struct {
	*defaultState
}

func (s awaitingAcceptanceState) accept(readyBy, acceptTime time.Time) error {
	if !acceptTime.Before(readyBy) {
		return NewErrIllegalArgument(fmt.Sprintf("readyBy %v is not after now %v", readyBy, acceptTime))
	}
	s.aggregateTicket.ReadyBy = readyBy
	s.aggregateTicket.AcceptTime = acceptTime
	s.aggregateTicket.state = s.aggregateTicket.acceptedState
	return nil
}

func (s awaitingAcceptanceState) changeLineItem(lineItems []TicketLineItem) error {
	s.aggregateTicket.LineItems = lineItems
	return nil
}

func (s awaitingAcceptanceState) cancel() error {
	s.aggregateTicket.state = s.aggregateTicket.canceledState
	return nil
}

type acceptedState struct {
	*defaultState
}

func (s acceptedState) preparing(preparingTime time.Time) error {
	s.aggregateTicket.state = s.aggregateTicket.preparingState
	s.aggregateTicket.PreparingTime = preparingTime
	return nil
}

func (s acceptedState) cancel() error {
	s.aggregateTicket.previouState = s.aggregateTicket.state
	s.aggregateTicket.state = s.aggregateTicket.cancelPendingState
	return nil
}

type preparingState struct {
	*defaultState
}

func (s preparingState) readyForPickUp(readyForPickup time.Time) error {
	s.aggregateTicket.state = s.aggregateTicket.readyForPickUpState
	s.aggregateTicket.ReadyForPickupTime = readyForPickup
	return nil
}

func (s preparingState) changeLineItem(lineItems []TicketLineItem) error {
	_ = lineItems
	return ErrTooLate
}

type readyForPickUpState struct {
	*defaultState
}

func (s readyForPickUpState) pickedUp(pickedUp time.Time) error {
	s.aggregateTicket.state = s.aggregateTicket.pickedUpState
	s.aggregateTicket.PickedUpTime = pickedUp
	return nil
}

type pickedUpState struct {
	*defaultState
}

type cancelPendingState struct {
	*defaultState
}

func (s cancelPendingState) confirmCancel() error {
	s.aggregateTicket.state = s.aggregateTicket.canceledState
	return nil
}

func (s cancelPendingState) undoCancel() error {
	s.aggregateTicket.state = s.aggregateTicket.previouState
	return nil
}

type canceledState struct {
	*defaultState
}
