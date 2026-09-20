setup:
	go install github.com/air-verse/air@latest

run:
	air

run-local:
	go run ./cmd/api

test:
	go test ./...

race:
	go test ./... -race

build:
	go build -o bin/api.exe ./cmd/api

