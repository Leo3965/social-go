run:
	go run ./cmd/api

test:
	go test ./...

race:
	go test ./... -race

build:
	go build -o bin/app .

