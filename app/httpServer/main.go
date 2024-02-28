package httpServer

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"

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

var mediaFiles library = library{
	entities:     make([]entity, 0),
	updateNeeded: true,
}

func app(c echo.Context) error {
	cfg := getConfig()

	if mediaFiles.updateNeeded {
		for _, subDir := range cfg.MediaTypes {
			err := filepath.WalkDir(cfg.StorageDir+"/"+subDir+"/", mediaFiles.visit)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			mediaFiles.updateNeeded = false
		}
	}

	var b strings.Builder
	b.WriteString("<HTML><BODY><table><tr><td>Type</td><td>Title</td></tr>")
	for _, media := range mediaFiles.entities {
		b.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", media.media, media.title))
	}
	b.WriteString("</table></BODY></HTML>")

	return c.HTML(http.StatusOK, b.String())
}

type library struct {
	entities     []entity
	updateNeeded bool
}
type entity struct {
	media string
	title string
}

func (v *library) visit(basePath string, entry fs.DirEntry, err error) error {
	if entry.Name() == "Movies" || entry.Name() == "TV Shows" {
		return nil
	}
	if err != nil {
		return err
	}

	parts := strings.Split(basePath, string(os.PathSeparator))

	temp := entity{
		media: parts[len(parts)-2],
		title: entry.Name(),
	}

	v.entities = append(v.entities, temp)

	return nil
}
