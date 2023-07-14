package main

// hello

import (
	"flag"
	"log"
	"sync"

	"go.uber.org/zap"

	"github.com/variety-jones/cfstress/pkg/executioner"
	"github.com/variety-jones/cfstress/pkg/store/mongodb"
	"github.com/variety-jones/cfstress/pkg/web"
)

func main() {
	var serverAddr, redisAddr, mongoAddr string
	flag.StringVar(&redisAddr, "redis-addr", "localhost:6379",
		"redis address")
	flag.StringVar(&mongoAddr, "mongo-addr", "mongodb://localhost:27017",
		"mongoDB address")
	flag.StringVar(&serverAddr, "server-addr", ":4000",
		"Server address")

	flag.Parse()

	ticketStore, err := mongodb.NewMongoStore(mongoAddr, "cfstress")
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := ticketStore.Close(); err != nil {
			zap.S().Fatal(err)
		}
	}()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	zap.S().Infof("redis-addr: %s", redisAddr)
	judge := executioner.NewJudge(redisAddr)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		srv := web.CreateWebServer(ticketStore, judge)
		if err := srv.ListenAndServe(serverAddr); err != nil {
			zap.S().Fatal(err)
		}
	}()
	wg.Wait()
}
