package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/qlfzn/viewlater/internal/store"
	"go.uber.org/zap"
)

// encapsulates video notes data
type Notes struct {
	Description string
	Type        string
}

type NewVideoPayload struct {
	Url   string `json:"url"`
	Notes string `json:"notes"`
}

type Handler struct {
	Logger *zap.SugaredLogger
}

func (h *Handler) SaveVideoHandler(w http.ResponseWriter, r *http.Request) {
	var payload NewVideoPayload

	Validate := validator.New()

	if err := readJSON(w, r, &payload); err != nil {
		// error handling
		h.Logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := Validate.Struct(payload); err != nil {
		h.Logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
		writeJSON(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// TODO: call store service
	video := &store.NewVideo{
		Url: payload.Url,
	}

	// success response
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "video received successfully",
		"url":     payload.Url,
		"notes":   payload.Notes,
	})
}
