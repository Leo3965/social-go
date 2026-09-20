package postgres

import (
	"context"
	"database/sql"
	"log"
)

type PostsStore struct {
	db *sql.DB
}

func (p *PostsStore) Create(ctx context.Context) error {
	// INSERT INTO posts ...
	log.Println("creating posts")
	return nil
}
