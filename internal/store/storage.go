package store

import (
	"context"

	"github.com/Leo3965/social/internal/data/model"
)

// Storage Any type that implements Storage must have two methods: PostsRepository() and UsersRepository().
type Storage interface {
	Posts() PostsRepository
	Users() UsersRepository
}

type PostsRepository interface {
	Create(ctx context.Context, post *model.Post) error
}

type UsersRepository interface {
	Create(ctx context.Context, user *model.User) error
}
