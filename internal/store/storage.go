package store

import (
	"context"
	"errors"
	"time"

	"github.com/Leo3965/social/internal/data/model"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("record already exists")
	ErrConcurrentUpdate  = errors.New("the record was modified by another request")
	QueryTimeoutDuration = time.Second * 5
)

// Storage Any type that implements Storage must have two methods: PostRepository() and UserRepository().
type Storage interface {
	Posts() PostRepository
	Users() UserRepository
	Comments() CommentRepository
	Followers() FollowerRepository
}

type PostRepository interface {
	Create(context.Context, *model.Post) error
	Find(context.Context, int64) (*model.Post, error)
	Delete(context.Context, int64) error
	Update(context.Context, *model.Post) error
	GetUserFeed(context.Context, int64) ([]model.FeedPost, error)
}

type UserRepository interface {
	Create(context.Context, *model.User) error
	Find(context.Context, int64) (*model.User, error)
}

type CommentRepository interface {
	Create(context.Context, *model.Comment) error
	FindByPostId(context.Context, int64) ([]model.Comment, error)
}

type FollowerRepository interface {
	Follow(ctx context.Context, userID int64, followerID int64) error
	Unfollow(ctx context.Context, userID int64, followerID int64) error
}
