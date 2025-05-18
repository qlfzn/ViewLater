package store

import (
	"context"
	"database/sql"
	"time"
)

type NewVideo struct {
	ID        int64    `json:"id"`
	Url       string   `json:"url"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"create_at"`
}

type VideoStore struct {
	db *sql.DB
}

func (v *VideoStore) SaveVideo(ctx context.Context, video *NewVideo) error {
	query := `
								INSERT INTO videos (url)
								VALUES $1 RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := v.db.QueryRowContext(
		ctx,
		query,
		video.Url,
	).Scan(
		&video.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
