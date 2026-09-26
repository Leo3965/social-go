package main

import (
	"log"
	"time"

	"github.com/Leo3965/social/internal/env"
	"github.com/Leo3965/social/internal/store/postgres"
)

func main() {

	arr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	db, err := postgres.New(arr, 3, 3, time.Minute*15)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	store := postgres.NewStorage(db)

	postgres.Seed(store)
}
