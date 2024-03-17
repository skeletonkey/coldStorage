package httpServer

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"

	"github.com/skeletonkey/coldStorage/pkg/library"
	"github.com/skeletonkey/lib-core-go/logger"
)

func RunServer() {
	config := getConfig()
	log := logger.Get()
	log.Info().Msg("Starting Echo server")

	e := echo.New()

	// routes
	e.GET("/", app)

	err := e.Start(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Error().Err(err).Int("Port", config.Port).Msg("error while attempting to start echo server")
	}
	log.Trace().Msg("Echo server should be running")
}

func app(c echo.Context) error {
	config := getConfig()
	library.Initialize(context.Background(), config.StorageDir, config.MediaTypes[0], config.MediaTypes[1], time.Duration(config.RefreshInterval)*time.Second)
	lib, err := library.Get()
	returnCode := http.StatusOK

	var b strings.Builder
	if err != nil {
		returnCode = http.StatusInternalServerError
		b.WriteString("<HTML><BODY>")
		b.WriteString(err.Error())
		b.WriteString("</BODY></HTML>")

	} else {
		b.WriteString("<HTML><BODY><table><tr><td>Type</td><td>Title</td></tr>")
		for _, movie := range lib.Movies {
			b.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", library.MoviesTopicTitle, movie.Name))
		}
		b.WriteString("</table><table><tr><td>Type</td><td>Show</td><td>Episode</td></tr>")
		for seriesKey, series := range lib.TVShows {
			fmt.Printf("Key: %s -> %+v\n", seriesKey, series)
			for _, episode := range series.Episodes {
				fmt.Printf("Ep title: %s\n", episode.Title)
				b.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>", library.TVShowsTopicTitle, series.Title, episode.Title))
			}
		}
		b.WriteString("</table></BODY></HTML>")
	}

	return c.HTML(returnCode, b.String())
}
