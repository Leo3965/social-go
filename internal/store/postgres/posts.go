package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Leo3965/social/internal/data/model"
	"github.com/Leo3965/social/internal/store"
	"github.com/lib/pq"
)

type PostRepository struct {
	db *sql.DB
}

func (pr *PostRepository) FindById(ctx context.Context, id int64) (*model.Post, error) {
	query := `SELECT id, content, title, user_id, created_at, updated_at, tags 
			  FROM posts WHERE id = $1`

	var post model.Post

	err := pr.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Content,
		&post.Title,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, store.ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (pr *PostRepository) Create(ctx context.Context, post *model.Post) error {
	query := `INSERT INTO posts (content, title, user_id, tags)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id, created_at, updated_at`

	err := pr.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
