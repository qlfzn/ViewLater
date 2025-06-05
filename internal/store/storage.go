package store

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotFound  = errors.New("resource not found")
	ErrConflict  = errors.New("resource already exists")
	QueryTimeout = time.Second * 5
)

type VideoRepository interface {
	SaveVideo(context.Context, *NewVideo) error
	GetVideoById(context.Context, int64) error
	Update(context.Context, *NewVideo) error
}

type Storage struct {
	Videos VideoRepository
}
