package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/variety-jones/cfstress/pkg/store"
)

type mongoStore struct {
	mongoClient        *mongo.Client
	countersCollection *mongo.Collection
	ticketsCollection  *mongo.Collection
}

func NewMongoStore(mongoURI, databaseName string) (store.TicketStore, error) {
	// Create a new client and connect to the server
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		return nil, err
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected and pinged.")
	mStore := new(mongoStore)
	mStore.mongoClient = client
	mStore.countersCollection = client.Database(databaseName).
		Collection("counters")
	mStore.ticketsCollection = client.Database(databaseName).
		Collection("tickets")

	uniqueIndex := true
	indexName, err := mStore.ticketsCollection.Indexes().
		CreateOne(context.TODO(), mongo.IndexModel{
			Keys: bson.D{{"ticket_id", 1}},
			Options: &options.IndexOptions{
				Unique: &uniqueIndex,
			},
		})

	fmt.Println("Created index: ", indexName)
	if err != nil {
		return nil, err
	}
	return mStore, nil
}
