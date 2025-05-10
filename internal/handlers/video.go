package handlers

import (
	"net/http"
	"time"
)

// struct for notes
type Notes struct {
	Description string
	Type        string
}

// struct for video params
type SaveVideoParams struct {
	Url       string
	Notes     Notes
	UpdatedAt time.Time
}

// TODO: middleware to validate input URL
func SaveVideoHandler(w http.ResponseWriter, r *http.Request) {
	err := app.VideoService.Save()
}
