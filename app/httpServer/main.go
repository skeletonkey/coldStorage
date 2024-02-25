package httpServer

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/skeletonkey/lib-core-go/logger"
)

func RunServer() {
	config := getConfig()
	log := logger.Get()
	log.Info().Msg("Starting Echo server")

	e := echo.New()

	// routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	err := e.Start(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Error().Err(err).Int("Port", config.Port).Msg("error while attempting to start echo server")
	}
	log.Trace().Msg("Echo server should be running")
}
