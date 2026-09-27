package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Leo3965/social/internal/store"
	"github.com/lib/pq"
)

type FollowerRepository struct {
	db *sql.DB
}

func (fr *FollowerRepository) Follow(ctx context.Context, userID int64, followerID int64) error {
	query := `INSERT INTO followers (user_id, follower_id, created_at)
   			  VALUES ($1, $2, NOW())`

	_, err := fr.db.ExecContext(
		ctx,
		query,
		userID,
		followerID,
	)
	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return store.ErrConflict
		}

		return err
	}

	return nil
}

func (fr *FollowerRepository) Unfollow(ctx context.Context, userID int64, followerID int64) error {
	query := `DELETE FROM followers
    		  WHERE followers.user_id = $1
			  AND followers.follower_id = $2`

	_, err := fr.db.ExecContext(
		ctx,
		query,
		userID,
		followerID,
	)
	if err != nil {
		return err
	}

	return nil
}
