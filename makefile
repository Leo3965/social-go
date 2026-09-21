setup:
	go install github.com/air-verse/air@latest

run:
	air

run-local:
	go run ./cmd

test:
	go test ./...

race:
	go test ./... -race

build:
	go build -o bin/api ./cmd

