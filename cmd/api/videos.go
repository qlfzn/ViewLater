package main

import (
	"encoding/json"
	"net/http"
)

// encapsulates video notes data
type Notes struct {
	Description string
	Type string
}

type CreateVideoPayload struct {
	Url string `json:"url"`
	Notes string
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // max 1mb allowed 
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func (app *application) saveVideoHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateVideoPayload
	if err := readJSON(w, r , &payload); if err != nil {
		// error handling 
	}


	
}
