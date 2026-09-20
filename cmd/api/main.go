package main

import (
	"log"

	"github.com/Leo3965/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	httpHandler := app.mount()

	log.Fatal(app.run(httpHandler))
}
