package executioner

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/variety-jones/cfstress/pkg/models"
)

// Jury implements the Judge interface.
type Jury struct {
	asynqClient *asynq.Client
}

// ProcessTicket enqueues a ticket to the redis queue.
func (j *Jury) ProcessTicket(ticket *models.Ticket) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(ticket)
	if err != nil {
		return nil, fmt.Errorf("could not marshal ticket with error %w",
			err)
	}

	task := asynq.NewTask(ticket.Type, payload)
	taskInfo, err := j.asynqClient.Enqueue(task)
	if err != nil {
		return nil, fmt.Errorf("could not enqueue task with error %w",
			err)
	}

	return taskInfo, nil
}

// NewJudge returns a concrete implementation of the Judge interface.
func NewJudge(redisAddr string) Judge {
	return &Jury{
		asynqClient: asynq.NewClient(
			asynq.RedisClientOpt{
				Addr: redisAddr,
			}),
	}
}
