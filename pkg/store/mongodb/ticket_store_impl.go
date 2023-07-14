package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/variety-jones/cfstress/pkg/models"
)

type counter struct {
	Seq int `json:"seq"`
}

func (store *mongoStore) Add(ticket *models.Ticket) (int, error) {
	filter := bson.D{{"_id", "ticket_counter"}}
	update := bson.D{{"$inc", bson.D{{"seq", 1}}}}

	createDocumentIfNotPresent := true
	previousRecord := store.countersCollection.
		FindOneAndUpdate(context.TODO(), filter, update,
			&options.FindOneAndUpdateOptions{
				Upsert: &createDocumentIfNotPresent,
			})

	// If the collection does not exist, previousRecord.Err() would be non-nil.
	// This indicates that this is our first time using the store.
	// We initialize the counter from 1. If it was indeed an error,
	// we will catch it via the UNIQUE_INDEX constraint on ticket_id in the
	// tickets collection.
	lastCounter := new(counter)
	if previousRecord != nil && previousRecord.Err() == nil {
		if err := previousRecord.Decode(lastCounter); err != nil {
			return -1, fmt.
				Errorf("could not decode last counter with error %w", err)
		}
	}

	ticket.TicketID = lastCounter.Seq + 1
	_, err := store.ticketsCollection.InsertOne(context.TODO(), ticket)
	if err != nil {
		return -1, err
	}

	return ticket.TicketID, nil
}

func (store *mongoStore) Query(id int) (*models.Ticket, error) {
	ticket := new(models.Ticket)
	return ticket, nil
}

func (store *mongoStore) Update(id int, updatedTicket *models.Ticket) error {
	return nil
}

func (store *mongoStore) Close() error {
	return store.mongoClient.Disconnect(context.TODO())
}
