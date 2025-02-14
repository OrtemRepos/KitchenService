package domain

import (
	"time"
)

type Ticket struct {
	ID           uint
	state        state
	previouState state
	RestaurantID uint
	LineItems    []TicketLineItem

	ReadyBy            time.Time
	AcceptTime         time.Time
	PreparingTime      time.Time
	PickedUpTime       time.Time
	ReadyForPickupTime time.Time

	createPendingState      createPendingState
	awaitingAcceptanceState awaitingAcceptanceState
	acceptedState           acceptedState
	preparingState          preparingState
	readyForPickUpState     readyForPickUpState
	pickedUpState           pickedUpState
	cancelPendingState      cancelPendingState
	canceledState           canceledState
}

func NewTicket(restaurantID, id uint, lineItems []TicketLineItem) *Ticket {
	t := &Ticket{
		ID:           id,
		RestaurantID: restaurantID,
		LineItems:    lineItems,
	}
	defaultState := newDefaultState(t)
	t.createPendingState = createPendingState{defaultState}
	t.awaitingAcceptanceState = awaitingAcceptanceState{defaultState}
	t.acceptedState = acceptedState{defaultState}
	t.preparingState = preparingState{defaultState}
	t.readyForPickUpState = readyForPickUpState{defaultState}
	t.pickedUpState = pickedUpState{defaultState}
	t.cancelPendingState = cancelPendingState{defaultState}
	t.canceledState = canceledState{defaultState}
	
	t.state = t.createPendingState
	return t
}

func (t *Ticket) ApplyEvent(e EventDomainTicket) error {
	if e.GetTicketID() != t.ID {
		return ErrMismatchedID
	}
	switch v := e.(type) {
	case *TicketCreateEvent:
		if err := t.state.confirmCreate(); err != nil {
			return err
		}
	case *TicketAcceptedEvent:
		if err := t.state.accept(v.ReadyBy, v.AcceptTime); err != nil {
			return err
		}
	case *TicketPreparingStartedEvent:
		if err := t.state.preparing(v.PreparingTime); err != nil {
			return err
		}
	case *TicketPreparationCompletedEvent:
		if err := t.state.readyForPickUp(v.ReadyForPickupTime); err != nil {
			return err
		}
	case *TicketPickedUpEvent:
		if err := t.state.pickedUp(v.PickedUpTime); err != nil {
			return err
		}
	default:
		return ErrUnsupportedEvent
	}
	return nil
}

func (t *Ticket) ConfirmCreate() (*TicketCreateEvent, error) {
	if err := t.state.confirmCreate(); err != nil {
		return nil, err
	}
	return NewTicketCreateEvent(t.ID, t.LineItems), nil
}

func (t *Ticket) CancelCreate() (*TicketCancelCreateEvent, error) {
	if err := t.state.cancelCreate(); err != nil {
		return nil, err
	}
	return NewTicketCancelCreateEvent(t.ID), nil
}

func (t *Ticket) Accept(redyBy time.Time) (*TicketAcceptedEvent, error) {
	acceptTime := time.Now()
	if err := t.state.accept(redyBy, acceptTime); err != nil {
		return nil, err
	}
	return NewTicketAcceptedEvent(t.ID, redyBy, acceptTime), nil
}

func (t *Ticket) Preparing() (*TicketPreparingStartedEvent, error) {
	preparingTime := time.Now()
	if err := t.state.preparing(preparingTime); err != nil {
		return nil, err
	}
	return NewTicketPreparingStartedEvent(t.ID, preparingTime), nil
}

func (t *Ticket) ReadyForPickup() (*TicketPreparationCompletedEvent, error) {
	readyForPickUp := time.Now()
	if err := t.state.readyForPickUp(readyForPickUp); err != nil {
		return nil, err
	}
	return NewTicketPreparationCompletedEvent(t.ID, readyForPickUp), nil
}

func (t *Ticket) PickedUp() (*TicketPickedUpEvent, error) {
	pickedUpTime := time.Now()
	if err := t.state.pickedUp(pickedUpTime); err != nil {
		return nil, err
	}
	return NewTicketPickedUpEvent(t.ID, pickedUpTime), nil
}

func (t *Ticket) ChangeLineItemQuantity(lineItems []TicketLineItem) (*TicketChangeLineItemEvent, error) {
	if err := t.state.changeLineItem(lineItems); err != nil {
		return nil, err
	}
	return NewTicketChangeLineItemEvent(t.ID, lineItems), nil
}

func (t *Ticket) Cancel() (*TicketCanceledEvent, error) {
	if err := t.state.cancel(); err != nil {
		return nil, err
	}
	return NewTicketCanceledEvent(t.ID, t.state == t.canceledState), nil
}

func (t *Ticket) ConfirmCancel() (*TicketCanceledEvent, error) {
	if err := t.state.confirmCancel(); err != nil {
		return nil, err
	}
	return NewTicketCanceledEvent(t.ID, true), nil
}

func (t *Ticket) UndoCancel() (*TicketUndoCanceledEvent, error) {
	if err := t.state.undoCancel(); err != nil {
		return nil, err
	}
	return NewTicketUndoCanceledEvent(t.ID), nil
}
