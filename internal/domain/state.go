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
	String() string
}

type defaultState struct {
	aggregateTicket *Ticket
}

func newDefaultState(aggregate *Ticket) *defaultState {
	s := defaultState{aggregate}
	return &s
}

func (s defaultState) confirmCreate() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) cancelCreate() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) accept(readyBy, acceptTime time.Time) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) preparing(preparingTime time.Time) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) readyForPickUp(readyForPickUp time.Time) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) pickedUp(pickedUp time.Time) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) changeLineItem(lineItems []TicketLineItem) error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) cancel() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) confirmCancel() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) undoCancel() error {
	return NewErrUnsupportedStateTransition(s.aggregateTicket.State)
}

func (s defaultState) String() string {
	return ""
}

type createPendingState struct {
	*defaultState
}

func (s createPendingState) confirmCreate() error {
	s.aggregateTicket.previouState = s.aggregateTicket.State
	s.aggregateTicket.State = s.aggregateTicket.awaitingAcceptanceState
	return nil
}

func (s createPendingState) cancelCreate() error {
	s.aggregateTicket.previouState = s.aggregateTicket.State
	s.aggregateTicket.State = s.aggregateTicket.cancelPendingState
	return nil
}

func (s createPendingState) String() string {
	return "CREATE_PENDING"
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
	s.aggregateTicket.State = s.aggregateTicket.acceptedState
	return nil
}

func (s awaitingAcceptanceState) changeLineItem(lineItems []TicketLineItem) error {
	s.aggregateTicket.LineItems = lineItems
	return nil
}

func (s awaitingAcceptanceState) cancel() error {
	s.aggregateTicket.State = s.aggregateTicket.canceledState
	return nil
}

func (s awaitingAcceptanceState) String() string {
	return "AWAITING_ACCEPTANCE"
}

type acceptedState struct {
	*defaultState
}

func (s acceptedState) preparing(preparingTime time.Time) error {
	s.aggregateTicket.State = s.aggregateTicket.preparingState
	s.aggregateTicket.PreparingTime = preparingTime
	return nil
}

func (s acceptedState) cancel() error {
	s.aggregateTicket.previouState = s.aggregateTicket.State
	s.aggregateTicket.State = s.aggregateTicket.cancelPendingState
	return nil
}

func (s acceptedState) String() string {
	return "ACCEPTED"
}

type preparingState struct {
	*defaultState
}

func (s preparingState) readyForPickUp(readyForPickup time.Time) error {
	s.aggregateTicket.State = s.aggregateTicket.readyForPickUpState
	s.aggregateTicket.ReadyForPickupTime = readyForPickup
	return nil
}

func (s preparingState) changeLineItem(lineItems []TicketLineItem) error {
	_ = lineItems
	return ErrTooLate
}

func (s preparingState) String() string {
	return "PREPARING"
}

type readyForPickUpState struct {
	*defaultState
}

func (s readyForPickUpState) pickedUp(pickedUp time.Time) error {
	s.aggregateTicket.State = s.aggregateTicket.pickedUpState
	s.aggregateTicket.PickedUpTime = pickedUp
	return nil
}

func (s readyForPickUpState) String() string {
	return "READY_FOR_PICKUP"
}

type pickedUpState struct {
	*defaultState
}

func (s pickedUpState) String() string {
	return "PICKED_UP"
}

type cancelPendingState struct {
	*defaultState
}

func (s cancelPendingState) confirmCancel() error {
	s.aggregateTicket.State = s.aggregateTicket.canceledState
	return nil
}

func (s cancelPendingState) undoCancel() error {
	s.aggregateTicket.State = s.aggregateTicket.previouState
	return nil
}

func (s cancelPendingState) String() string {
	return "CANCEL_PENDING"
}

type canceledState struct {
	*defaultState
}

func (s canceledState) String() string {
	return "CANCELED"
}
