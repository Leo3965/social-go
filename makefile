include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: local-setup
local-setup:
	go install github.com/air-verse/air@latest

# make migrate-create name=create_users
.PHONY: migrate-create
migrate-create:
	migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY: migrate-up
migrate-up:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATION_ADDR) up

.PHONY: migrate-down
migrate-down:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DB_MIGRATION_ADDR) down


.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down -v

.PHONY: run
run:
	air

.PHONY: run-local
run-local:
	go run ./cmd

.PHONY: test
test:
	go test ./...

.PHONY: race
race:
	go test ./... -race

.PHONY: build
build:
	go build -o bin/api ./cmd

