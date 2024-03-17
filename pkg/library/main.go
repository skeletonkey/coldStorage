package library

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

var libraryInstance Library

const (
	moviesKey         = "movies"
	MoviesTopicTitle  = "Movie"
	tvShowsKey        = "tvShows"
	TVShowsTopicTitle = "TV Episode"
)

type Library struct {
	dirs            map[string]string
	Movies          []*Movie
	TVShows         map[string]*Series
	refreshRequired bool
}

type Series struct {
	Title    string
	Episodes []*Show
}

type Show struct {
	Season  int
	Episode int
	Title   string
}

type Movie struct {
	Name string
}

func (v *Library) addMovie(title string) error {
	fmt.Printf("addMovie: %s\n", title)
	v.Movies = append(v.Movies, &Movie{
		Name: title,
	})

	return nil
}

func (v *Library) addEpisode(title string) error {
	fmt.Printf("addEpisode: %s\n", title)
	parts := strings.Split(title, " - ")
	seriesName := parts[0]
	episodeName := parts[0]
	if len(parts) == 3 {
		episodeName = parts[2]
	}

	// series -> parts[0]
	var found bool
	var tvSeries *Series
	if tvSeries, found = v.TVShows[seriesName]; !found {
		fmt.Printf("creating series: %s\n", seriesName)
		tvSeries = &Series{
			Title:    seriesName,
			Episodes: make([]*Show, 0),
		}
		v.TVShows[seriesName] = tvSeries
	} else {
		fmt.Printf("adding episode: %s\n", episodeName)
		tvSeries.Episodes = append(tvSeries.Episodes, &Show{
			Season:  0,
			Episode: 0,
			Title:   episodeName,
		})
	}

	return nil
}

func Initialize(ctx context.Context, baseDir string, moviesDir string, tvShowsDir string, refreshInterval time.Duration) error {
	libraryInstance = Library{
		dirs: map[string]string{
			moviesKey:  fmt.Sprintf("%s/%s", baseDir, moviesDir),
			tvShowsKey: fmt.Sprintf("%s/%s", baseDir, tvShowsDir),
		},
		Movies:          make([]*Movie, 0),
		TVShows:         make(map[string]*Series),
		refreshRequired: true,
	}

	if refreshInterval > 0 {
		ticker := time.NewTicker(refreshInterval)

		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					libraryInstance.refresh()
				}
			}
		}()
	}

	return nil
}

func (v *Library) refresh() error {
	fmt.Println("refresh")
	// movies
	dir, ok := v.dirs[moviesKey]
	if !ok {
		return fmt.Errorf("no %s directory found", moviesKey)
	}
	if err := v.processMovies(dir); err != nil {
		return err
	}

	// tv shows
	dir, ok = v.dirs[tvShowsKey]
	if !ok {
		return fmt.Errorf("no %s directory found", tvShowsKey)
	}
	if err := v.processTVShows(dir); err != nil {
		return err
	}

	return nil
}

// func (v library) processNode(dir string, process func(string) error) error {
// 	files, err := os.ReadDir(dir)
// 	if err != nil {
// 		return fmt.Errorf("unable to read dir (%s): %s", dir, err)
// 	}
// 	for _, file := range files {
// 		if file.IsDir() {
// 			if info, err := file.Info(); err != nil {
// 				return err
// 			} else {
// 				return v.processNode(info.Name(), process)
// 			}
// 		} else {
// 			return process(file.Name())
// 		}
// 	}

//		return nil
//	}
func (v *Library) processMovies(dir string) error {
	fmt.Printf("movie: %s\n", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("unable to read dir (%s): %s", dir, err)
	}
	for _, file := range files {
		if file.IsDir() {
			if err = v.processMovies(fmt.Sprintf("%s/%s", dir, file.Name())); err != nil {
				return err
			}
		} else {
			if err = v.addMovie(file.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (v *Library) processTVShows(dir string) error {
	fmt.Printf("tv: %s\n", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("unable to read dir (%s): %s", dir, err)
	}
	for _, file := range files {
		if file.IsDir() {
			if err = v.processTVShows(fmt.Sprintf("%s/%s", dir, file.Name())); err != nil {
				return err
			}
		} else {
			if err = v.addEpisode(file.Name()); err != nil {
				return err
			}
		}
	}

	return nil
}

func Get() (Library, error) {
	var err error
	if libraryInstance.refreshRequired {
		err = libraryInstance.refresh()
	}
	return libraryInstance, err
}
