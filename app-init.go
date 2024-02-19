package main

//go:generate go run app-init.go

import instanceGen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	app := instanceGen.NewApp("cold-storage", "app")
	app.
		WithPackages("archimedes", "woody", "httpServer", "dataStore").
		WithDependencies(
			"github.com/mattn/go-sqlite3",
			"github.com/labstack/echo/v4",
		).
		WithGithubWorkflows("linter", "test").
		WithGoVersion("1.22").
		WithCGOEnabled().
		WithMakefile()
}
