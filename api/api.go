package api

// Save video params
type SaveVideoParams struct {
	Url string
}

// Save video response
type SaveVideoResponse struct {
	Code int
}

// Error struct
type Error struct {
	Code    int
	Message string
}
