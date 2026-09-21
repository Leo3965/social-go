package postgres

import (
	"context"
	"database/sql"
	"log"
)

type PGUsersRepository struct {
	db *sql.DB
}

func (u *PGUsersRepository) Create(ctx context.Context) error {
	// INSERT INTO users ...
	log.Println("creating users")
	return nil
}
