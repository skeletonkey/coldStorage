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

var mediaFiles library

func RunServer() {
	config := getConfig()
	log := logger.Get()
	log.Info().Msg("Starting Echo server")

	mediaFiles = newLibrary(config)

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
	for mediaTitle, mediaNode := range mediaFiles.entities {
		if mediaNode.parent == "" {
			// top of the pyramid is a category
			continue
		}
		b.WriteString(fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", mediaNode.media, mediaTitle))
	}
	b.WriteString("</table></BODY></HTML>")

	return c.HTML(http.StatusOK, b.String())
}

type library struct {
	entities     map[string]entity
	updateNeeded bool
}
type entity struct {
	media string
	title string
	parent string
}

func newLibrary(cfg *httpServer) library {
	l := library{
		entities: make(map[string]entity, 2),
		updateNeeded: true,
	}

	for _, mediaType := range cfg.MediaTypes {
		l.entities[mediaType] = entity{
			media: mediaType,
			title: mediaType,
			parent: "",
		}
	}

	return l
}

func (v *library) visit(basePath string, entry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if _, ok := v.entities[entry.Name()]; ok {
		// do not reprocess existing entries
		return nil
	}

	parts := strings.Split(basePath, string(os.PathSeparator))
	parent, ok := v.entities[parts[len(parts)-2]]
	if !ok {
		return fmt.Errorf("could not find '%s' in entities", parts[len(parts)-2])
	}

	v.entities[entry.Name()] = entity{
		media: parent.media,
		title: entry.Name(),
		parent: parent.title,
	}

	return nil
}
