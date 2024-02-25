package main

import (
	"fmt"
	"time"

	"github.com/skeletonkey/coldStorage/app/dataStore"
	"github.com/skeletonkey/coldStorage/app/queue"
	"github.com/skeletonkey/lib-core-go/logger"
)

func main() {
	fmt.Println("Starting up!")

	log := logger.Get()
	log.Info().Msg("Starting app")

	// Check on the DB
	if _, err := dataStore.GetDB(); err != nil {
		log.Error().Err(err).Msg("Got error when trying to get DB")
	}

	queue.GetChannel()
	defer func() {
		queue.CloseChannel()
		queue.CloseConnection()
	}()
	go func() {
		for {
		queue.PublishMsg("thaw", "Testing at: "+time.Now().String())
		time.Sleep(10 * time.Second)
		}
	}()

	// httpServer.RunServer()

	time.Sleep(1 * time.Minute)
	fmt.Println("Bye Bye!")
}
