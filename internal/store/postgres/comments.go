package postgres

import (
	"context"
	"database/sql"

	"github.com/Leo3965/social/internal/data/model"
	"github.com/Leo3965/social/internal/store"
)

type CommentRepository struct {
	db *sql.DB
}

func (cr *CommentRepository) Create(ctx context.Context, comment *model.Comment) error {
	query := `
        INSERT INTO comments (user_id, post_id, content)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `

	ctx, cancel := context.WithTimeout(ctx, store.QueryTimeoutDuration)
	defer cancel()

	err := cr.db.QueryRowContext(
		ctx,
		query,
		comment.UserID,
		comment.PostID,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepository) FindByPostId(ctx context.Context, postID int64) ([]model.Comment, error) {
	query := `SELECT c.id, c.user_id, c.post_id, c.content, c.created_at, c.updated_at,
			  u.username, u.id
 			  FROM comments c
   			  JOIN users u ON u.id = c.user_id
			  WHERE c.post_id = $1
			  ORDER BY  c.created_at DESC;`

	rows, err := cr.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var c model.Comment
		c.User = model.User{}
		err := rows.Scan(&c.ID, &c.UserID, &c.PostID, &c.Content, &c.CreatedAt, &c.UpdatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}
