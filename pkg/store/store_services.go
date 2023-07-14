package store

import "github.com/variety-jones/cfstress/pkg/models"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . TicketStore

type TicketStore interface {
	Add(ticket *models.Ticket) (int, error)
	Query(id int) (*models.Ticket, error)
	Update(id int, updatedTicket *models.Ticket) error
	Close() error
}
