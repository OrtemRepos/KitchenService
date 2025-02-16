package domain

import (
	"fmt"
)

type ErrUnsupportedStateTransition struct {
	state state
}

func (e ErrUnsupportedStateTransition) Error() string {
	return fmt.Sprintf("Unsupported transition for state %v", e.state)
}

func NewErrUnsupportedStateTransition(state state) *ErrUnsupportedStateTransition {
	return &ErrUnsupportedStateTransition{state}
}

type ErrIllegalArgument struct {
	msg string
}

func NewErrIllegalArgument(msg string) *ErrIllegalArgument {
	return &ErrIllegalArgument{msg}
}

func (e ErrIllegalArgument) Error() string {
	return e.msg
}

var ErrMismatchedID = fmt.Errorf("Aggregate and entity ids do not match")

var ErrUnsupportedEvent = fmt.Errorf("Unsupported event type")

var ErrTooLate = fmt.Errorf("Too late to change the order")

var ErrNotFound = fmt.Errorf("Not found")