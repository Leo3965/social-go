package postgres

import (
	"context"
	"database/sql"
	"log"
)

type UsersStore struct {
	db *sql.DB
}

func (u *UsersStore) Create(ctx context.Context) error {
	// INSERT INTO users ...
	log.Println("creating users")
	return nil
}
