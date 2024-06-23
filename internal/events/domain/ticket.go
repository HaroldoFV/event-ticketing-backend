package domain

import (
	"errors"
	"github.com/google/uuid"
)

type TicketKind string

var ErrTicketPriceZero = errors.New("ticket price must be greater than zero")

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

func NewTicket(event *Event, spot *Spot, ticketType TicketKind) (*Ticket, error) {
	if !IsValidTicketType(ticketType) {
		return nil, errors.New("invalid ticket type")
	}
	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketKind: ticketType,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

func IsValidTicketType(ticketType TicketKind) bool {
	switch ticketType {
	case TicketKindHalf, TicketKindFull:
		return true
	default:
		return false
	}
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
