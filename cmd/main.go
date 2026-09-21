package main

import (
	"log"

	"github.com/Leo3965/social/cmd/api"
	"github.com/Leo3965/social/internal/env"
	"github.com/Leo3965/social/internal/store/postgres"
)

func main() {
	cfg := api.Config{
		Addr: env.GetString("ADDR", ":8080"),
	}

	store := postgres.NewStorage(nil)

	app := &api.Application{
		Config: cfg,
		Store:  store,
	}

	httpHandler := app.Mount()

	log.Fatal(app.Run(httpHandler))
}
