package model

import "time"

type Follower struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	FollowerID int64     `json:"follower_id"`
	CreatedAt  time.Time `json:"created_at"`
}
