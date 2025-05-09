package api

// struct for saving user's additonal metadata
type Notes struct {
	Description string
	Tags        string
}

// Save video params
type SaveVideoParams struct {
	Url   string
	Notes Notes
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
