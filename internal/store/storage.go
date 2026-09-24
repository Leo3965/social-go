package store

import (
	"context"
	"errors"

	"github.com/Leo3965/social/internal/data/model"
)

var (
	ErrNotFound = errors.New("record not found")
)

// Storage Any type that implements Storage must have two methods: PostRepository() and UserRepository().
type Storage interface {
	Posts() PostRepository
	Users() UserRepository
	Comments() CommentRepository
}

type PostRepository interface {
	Create(context.Context, *model.Post) error
	Find(context.Context, int64) (*model.Post, error)
	Delete(context.Context, int64) error
	Update(context.Context, *model.Post) error
}

type UserRepository interface {
	Create(context.Context, *model.User) error
}

type CommentRepository interface {
	Create(context.Context, *model.Comment) error
	FindByPostId(context.Context, int64) ([]model.Comment, error)
}
