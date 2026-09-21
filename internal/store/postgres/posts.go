package postgres

import (
	"context"
	"database/sql"

	"github.com/Leo3965/social/internal/data/model"
	"github.com/lib/pq"
)

type PGPostsRepository struct {
	db *sql.DB
}

func (p *PGPostsRepository) Create(ctx context.Context, post *model.Post) error {
	query := `INSERT INTO post (content, title, user_id, tags)
			  VALUES (1$, 2$, 3$, 4$)
			  RETURNING id, created_at, updated_at`

	err := p.db.QueryRowContext(
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
