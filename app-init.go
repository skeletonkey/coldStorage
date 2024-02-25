package main

//go:generate go run app-init.go

import instanceGen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	goVersion := "1.22"
	// primary coordinator
	app := instanceGen.NewApp("cold-storage", "app")
	app.
		WithPackages("httpServer", "dataStore", "queue").
		WithDependencies(
			"github.com/labstack/echo/v4",
			"github.com/mattn/go-sqlite3",
			"github.com/rabbitmq/amqp091-go",
		).
		WithGithubWorkflows("linter", "test").
		WithGoVersion(goVersion).
		WithCGOEnabled().
		WithMakefile()

	remoteApp := instanceGen.NewApp("cold-storage-remote", "appRemote")
	remoteApp.WithPackages("queue")
}
