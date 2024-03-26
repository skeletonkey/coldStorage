// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
.DEFAULT_GOAL=build

build:
	go fmt ./...
	go vet ./...
	CGO_ENABLED=1 go build -o bin/cold-storage app/*.go

install:
	cp bin/cold-storage /usr/local/sbin/cold-storage

golib-latest:
	go get -u github.com/labstack/echo/v4@latest
	go get -u github.com/mattn/go-sqlite3@latest
	go get -u github.com/rabbitmq/amqp091-go@latest
	go get -u github.com/skeletonkey/lib-core-go@latest
	go get -u github.com/skeletonkey/lib-instance-gen-go@latest

	go mod tidy

app-init:
	go generate
