package store

import (
	"context"
)

// Storage Any type that implements Storage must have two methods: PostsRepository() and UsersRepository().
type Storage interface {
	Posts() PostsRepository
	Users() UsersRepository
}

type PostsRepository interface {
	Create(ctx context.Context) error
}

type UsersRepository interface {
	Create(ctx context.Context) error
}
