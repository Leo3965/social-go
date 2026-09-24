package postgres

import (
	"database/sql"

	"github.com/Leo3965/social/internal/store"
)

type Storage struct {
	posts    store.PostRepository
	users    store.UserRepository
	comments store.CommentRepository
}

func (pg *Storage) Posts() store.PostRepository {
	return pg.posts
}

func (pg *Storage) Users() store.UserRepository {
	return pg.users
}

func (pg *Storage) Comments() store.CommentRepository {
	return pg.comments
}

func NewStorage(db *sql.DB) store.Storage {
	return &Storage{
		posts:    &PostRepository{db},
		users:    &UserRepository{db},
		comments: &CommentRepository{db},
	}
}
