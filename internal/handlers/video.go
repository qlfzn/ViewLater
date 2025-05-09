package handlers

import (
	"encoding/json"
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
	Notes     string
	UpdatedAt time.Time
}

// TODO: middleware to validate input URL
func saveVideos(w http.ResponseWriter, r *http.Request) {
	var video SaveVideoParams
	_ = json.NewDecoder(r.Body).Decode(&video)

	if video.Url == "" {
		http.Error(w, "Please provide a valid URL", http.StatusBadRequest)
	}
}
