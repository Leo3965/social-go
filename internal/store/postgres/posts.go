package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Leo3965/social/internal/data/application"
	"github.com/Leo3965/social/internal/data/model"
	"github.com/Leo3965/social/internal/store"
	"github.com/lib/pq"
)

type PostRepository struct {
	db *sql.DB
}

func (pr *PostRepository) Find(ctx context.Context, id int64) (*model.Post, error) {
	query := `SELECT id, content, title, user_id, created_at, updated_at, tags, version 
			  FROM posts WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, store.QueryTimeoutDuration)
	defer cancel()

	var post model.Post

	err := pr.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Content,
		&post.Title,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
		&post.Version,
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

	ctx, cancel := context.WithTimeout(ctx, store.QueryTimeoutDuration)
	defer cancel()

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

func (pr *PostRepository) Update(ctx context.Context, post *model.Post) error {
	query := `UPDATE posts
       		  SET title = $1,
            	content = $2,
            	version = version + 1,
            	updated_at = NOW()
        	  WHERE id = $3 AND version = $4
        	  RETURNING version`

	if err := pr.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(&post.Version); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return store.ErrConcurrentUpdate
		default:
			return err
		}
	}

	return nil
}

func (pr *PostRepository) GetUserFeed(ctx context.Context, userID int64, pagination application.PaginatedFeedQuery) ([]model.FeedPost, error) {
	query := `
		SELECT
			p.id,
			p.user_id,
			p.title,
			p.content,
			p.created_at,
			p.updated_at,
			p.version,
			p.tags,
			u.username,
			COUNT(c.id) AS comments_count
		FROM posts p
		LEFT JOIN comments c ON c.post_id = p.id
		LEFT JOIN users u ON p.user_id = u.id
		JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1
		WHERE
			f.user_id = $1 OR p.user_id = $1
		GROUP BY p.id, u.id, u.username
		ORDER BY p.created_at ` + pagination.Sort + `
		LIMIT $2 OFFSET $3
	`

	ctx, cancel := context.WithTimeout(ctx, store.QueryTimeoutDuration)
	defer cancel()

	rows, err := pr.db.QueryContext(
		ctx,
		query,
		userID,
		pagination.Limit,
		pagination.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feed := make([]model.FeedPost, 0)

	for rows.Next() {
		var p model.FeedPost

		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Version,
			pq.Array(&p.Tags),
			&p.User.Username,
			&p.CommentsCount,
		)
		if err != nil {
			return nil, err
		}

		feed = append(feed, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return feed, nil
}
