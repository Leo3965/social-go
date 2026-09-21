package postgres

import (
	"database/sql"

	"github.com/Leo3965/social/internal/store"
)

type PGStorage struct {
	posts store.PostsRepository
	users store.UsersRepository
}

func (pg *PGStorage) Posts() store.PostsRepository {
	return pg.posts
}

func (pg *PGStorage) Users() store.UsersRepository {
	return pg.users
}

func NewPostgresStorage(db *sql.DB) store.Storage {
	return &PGStorage{
		posts: &PGPostsRepository{db},
		users: &PGUsersRepository{db},
	}
}
