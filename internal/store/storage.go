package store

import (
	"context"
)

// Storage Any type that implements Storage must have two methods: Posts() and Users().
type Storage interface {
	Posts() Posts
	Users() Users
}

type Posts interface {
	Create(ctx context.Context) error
}

type Users interface {
	Create(ctx context.Context) error
}
