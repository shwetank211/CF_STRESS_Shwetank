package executioner

import (
	"github.com/hibiken/asynq"

	"github.com/variety-jones/cfstress/pkg/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . Judge

// Judge is responsible for processing all the tickets.
type Judge interface {
	ProcessTicket(ticket *models.Ticket) (*asynq.TaskInfo, error)
}
