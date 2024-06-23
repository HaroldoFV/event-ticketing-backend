package domain

import (
	"errors"
	"github.com/google/uuid"
)

type TicketKind string

var (
	ErrTicketPriceZero   = errors.New("ticket price must be greater than zero")
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
)

const (
	TicketKindHalf TicketKind = "half" // Half-price ticket
	TicketKindFull TicketKind = "full" // Full-price ticket
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

func NewTicket(event *Event, spot *Spot, ticketKind TicketKind) (*Ticket, error) {
	if !IsValidTicketType(ticketKind) {
		return nil, ErrInvalidTicketKind
	}
	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketKind: ticketKind,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

func IsValidTicketType(ticketKind TicketKind) bool {
	//switch ticketKind {
	//case TicketKindHalf, TicketKindFull:
	//	return true
	//default:
	//	return false
	//}
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}
	return nil
}
