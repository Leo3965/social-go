package postgres

import (
	"database/sql"

	"github.com/Leo3965/social/internal/store"
)

type PGStorage struct {
	posts store.Posts
	users store.Users
}

func (pg *PGStorage) Posts() store.Posts {
	return pg.posts
}

func (pg *PGStorage) Users() store.Users {
	return pg.users
}

func NewPostgresStorage(db *sql.DB) store.Storage {
	return &PGStorage{
		posts: &PostsStore{db},
		users: &UsersStore{db},
	}
}
