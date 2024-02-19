package main

import (
	"fmt"
	"time"

	"github.com/skeletonkey/coldStorage/app/dataStore"
	"github.com/skeletonkey/coldStorage/app/httpServer"
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

	httpServer.RunServer()

	time.Sleep(1 * time.Minute)
	fmt.Println("Bye Bye!")
}
