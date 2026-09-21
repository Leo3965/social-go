package postgres

import (
	"context"
	"database/sql"
	"log"
)

type PGPostsRepository struct {
	db *sql.DB
}

func (p *PGPostsRepository) Create(ctx context.Context) error {
	// INSERT INTO posts ...
	log.Println("creating posts")
	return nil
}
