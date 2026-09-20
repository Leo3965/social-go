package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	httpHandler := app.mount()

	log.Fatal(app.run(httpHandler))
}
