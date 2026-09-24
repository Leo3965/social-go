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

func (pr *PostRepository) Find(ctx context.Context, id int64) (*model.Post, error) {
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

func (pr *PostRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts
			  WHERE id = $1`

	result, err := pr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return store.ErrNotFound
	}

	return nil
}

func (pr PostRepository) Update(ctx context.Context, post *model.Post) error {
	query := `UPDATE posts
       		  SET title = $1,
            	content = $2,
            	updated_at = NOW()
        	  WHERE id = $3`

	result, err := pr.db.ExecContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return store.ErrNotFound
	}

	return nil
}
