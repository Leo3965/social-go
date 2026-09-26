package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Leo3965/social/internal/data/model"
	"github.com/Leo3965/social/internal/store"
)

type UserRepository struct {
	db *sql.DB
}

func (ur *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (username, password, email)
			  VALUES ($1, $2, $3)
			  RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, store.QueryTimeoutDuration)
	defer cancel()

	err := ur.db.QueryRowContext(
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

func (ur *UserRepository) Find(ctx context.Context, userID int64) (*model.User, error) {
	query := `SELECT id, username, email, created_at, updated_at 
			  FROM users
			  WHERE users.id = $1`

	var user = model.User{}
	if err := ur.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, store.ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
