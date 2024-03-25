package main

import (
	"fmt"

	"github.com/skeletonkey/coldStorage/app/httpServer"
	"github.com/skeletonkey/lib-core-go/logger"
)

func main() {
	fmt.Println("Starting up!")

	log := logger.Get()
	log.Info().Msg("Starting app")

	httpServer.RunServer()

	fmt.Println("Shutting down!")
}
