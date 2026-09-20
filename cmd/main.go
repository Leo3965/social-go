package main

import (
	"log"

	"github.com/Leo3965/social/cmd/api"
	"github.com/Leo3965/social/internal/env"
)

func main() {
	cfg := api.Config{
		Addr: env.GetString("ADDR", ":8080"),
	}

	app := &api.Application{
		Config: cfg,
	}

	httpHandler := app.Mount()

	log.Fatal(app.Run(httpHandler))
}
