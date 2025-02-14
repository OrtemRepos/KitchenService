package domain

type TicketState int

const (
	CREATE_PENDING = iota
	AWAITING_ACCEPTANCE
	ACCEPTED
	PREPARING
	READY_FOR_PICKUP
	PICKED_UP
	CANCEL_PENDING
	CANCELED
	REVISION_PENDING
)

func (s TicketState) String() string {
	switch s {
	case CREATE_PENDING:
		return "CREATE_PENDING"
	case AWAITING_ACCEPTANCE:
		return "AWAITING_ACCEPTANCE"
	case ACCEPTED:
		return "ACCEPTED"
	case PREPARING:
		return "PREPARING"
	case READY_FOR_PICKUP:
		return "READY_FOR_PICKUP"
	case PICKED_UP:
		return "PICKED_UP"
	case CANCEL_PENDING:
		return "CANCEL_PENDING"
	case CANCELED:
		return "CANCELLED"
	case REVISION_PENDING:
		return "REVISION_PENDING"
	default:
		return "unidentified state"
	}
}
