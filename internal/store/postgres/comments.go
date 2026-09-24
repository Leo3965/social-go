package postgres

import (
	"context"
	"database/sql"

	"github.com/Leo3965/social/internal/data/model"
)

type CommentRepository struct {
	db *sql.DB
}

func (cr *CommentRepository) Create(ctx context.Context, comment *model.Comment) error {
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
