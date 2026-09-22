local-setup:
	go install github.com/air-verse/air@latest

up:
	docker-compose up -d

down:
	docker-compose down -v

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

