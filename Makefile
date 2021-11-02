.PHONY: precommit test build

precommit:
	golangci-lint run ./...
	go mod tidy
	go mod verify

test:
	go test -race -cover ./...

build:
	go build -o cron