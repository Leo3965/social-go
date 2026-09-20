run: build
	.\bin\api.exe

run-local:
	go run ./cmd/api

test:
	go test ./...

race:
	go test ./... -race

build:
	go build -o bin/api.exe ./cmd/api

