package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os/exec"
	"time"

	"go.uber.org/zap"

	"github.com/hibiken/asynq"

	"github.com/variety-jones/cfstress/pkg/models"
	"github.com/variety-jones/cfstress/pkg/store/mongodb"
)

func handler(ctx context.Context, t *asynq.Task) error {
	ticket := new(models.Ticket)
	if err := json.Unmarshal(t.Payload(), ticket); err != nil {
		// If unmarshalling fails, there's no point in retrying the task.
		// Hence, we log the error but don't return it.
		zap.S().Errorf("Could not unmarshal ticket: %s with error: %v",
			string(t.Payload()), err)
		return nil
	}
	zap.S().Infof("Tikect: %d Simulating stress test...", ticket.TicketID)

	// TODO: Write the code for invoking the bash script that creates the
	// required files in a certain directory (input/expectedOutput/yourOutput).
	// Then, read these files and update the MongoDB databse.
	// This handler does not have access to the Mongo conenction, but you can
	// use dependency injection (by wrapping it as a struct's method) to get
	// the connection.
	ctx, cancel := context.WithTimeout(
		context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
	}

	return nil
}

func main() {
	var workers int
	var redisAddr, mongoAddr string
	flag.IntVar(&workers, "workers", 10,
		"Number of workers")
	flag.StringVar(&redisAddr, "redis-addr", "localhost:6379",
		"redis address")
	flag.StringVar(&mongoAddr, "mongo-addr", "mongodb://localhost:27017",
		"mongoDB address")

	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	ticketStore, err := mongodb.NewMongoStore(mongoAddr, "cfstress")
	defer func() {
		if err := ticketStore.Close(); err != nil {
			zap.S().Fatal(err)
		}
	}()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: redisAddr,
		},
		asynq.Config{
			Concurrency: workers,
		})

	mux := asynq.NewServeMux()
	mux.HandleFunc("stress:test", handler)
	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}
