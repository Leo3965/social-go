package postgres

import (
	"context"
	"database/sql"

	"github.com/Leo3965/social/internal/data/model"
)

type PGUsersRepository struct {
	db *sql.DB
}

func (p *PGUsersRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO post (username, password, email)
			  VALUES (1$, 2$, 3$)
			  RETURNING id, created_at, updated_at`

	err := p.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
