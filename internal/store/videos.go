package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Video struct {
	ID          int64     `json:"id"`
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	UserID      int64     `json:"user_id"`
	Tags        []string  `json:"tags"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type VideoStore struct {
	DB *pgxpool.Pool
}

func (v *VideoStore) SaveVideo(ctx context.Context, video *Video) error {
	// insert into videos table
	query := `
		INSERT INTO videos (url)
		VALUES ($1) RETURNING id, created_at
	`
	var createdAt time.Time

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := v.DB.QueryRow(ctx, query, video.Url).Scan(&video.ID, &createdAt)
	if err != nil {
		return err
	}

	// insert into user_videos
	queryUserVideos := `
		INSERT INTO user_videos (video_id, title, description, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	err = v.DB.QueryRow(
		ctx,
		queryUserVideos,
		video.ID,
		video.Title,
		video.Description,
		video.Tags,
	).Scan(&video.ID, &video.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (v *VideoStore) GetVideoById(ctx context.Context, videoID int64) (*Video, error) {
	query := `
		SELECT * FROM videos (id)
		LEFT JOIN user_videos
		ON videos.id = user_videos.video_id
		VALUES ($1)
	`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var video Video

	err := v.DB.QueryRow(ctx, query, videoID).Scan(
		&video.ID,
		&video.Url,
		&video.Title,
		&video.Description,
		&video.Tags,
		&video.CreatedAt,
		&video.UserID,
	)
	if err != nil {
		return nil, err
	}

	return &video, nil
}
