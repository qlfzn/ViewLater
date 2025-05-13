package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// encapsulates video notes data
type Notes struct {
	Description string
	Type        string
}

type CreateVideoPayload struct {
	Url   string `json:"url"`
	Notes string `json:"notes"`
}

func (n Notes) New(s string) string {
	return "This is notes: " + s
}

func (c CreateVideoPayload) New(p string) string {
	return "This is payload: " + p
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // max 1mb allowed
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func (app *application) saveVideoHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateVideoPayload

	Validate := validator.New()

	if err := readJSON(w, r, &payload); err != nil {
		// error handling
		app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
		writeJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// success response
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "video received successfully",
		"url":     payload.Url,
		"notes":   payload.Notes,
	})
}
