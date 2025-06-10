package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/qlfzn/viewlater/internal/middleware"
	"github.com/qlfzn/viewlater/internal/store"
	"go.uber.org/zap"
)

type NewVideoPayload struct {
	ID          int64    `json:"id"`
	Url         string   `json:"url"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// consists of all dependencies related
type Handler struct {
	Logger *zap.SugaredLogger
	Store  store.VideoStore
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

	video := &store.Video{
		Url:         payload.Url,
		Title:       payload.Title,
		Description: payload.Description,
		Tags:        payload.Tags,
	}

	ctx := r.Context()

	source := &middleware.Source{}
	if err := source.ParseAndValidateUrl(video.Url); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid origin",
		})
		return
	}

	if err := h.Store.SaveVideo(ctx, video); err != nil {
		h.Logger.Errorf("error saving video: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to save video",
		})
		return
	}

	// success response
	writeJSON(w, http.StatusOK, map[string]string{
		"message":     "video received successfully",
		"url":         payload.Url,
		"description": payload.Description,
	})
}

func (h *Handler) GetVideoHandler(w http.ResponseWriter, r *http.Request) *store.Video {
	var payload NewVideoPayload

	Validate := validator.New()

	err := readJSON(w, r, payload)
	if err != nil {
		h.Logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
		writeJSON(w, http.StatusBadRequest, err.Error())
		return nil
	}

	if err := Validate.Struct(payload); err != nil {
		h.Logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
		writeJSON(w, http.StatusUnprocessableEntity, err.Error())
		return nil
	}

	video := &store.Video{
		ID:          payload.ID,
		Url:         payload.Url,
		Title:       payload.Title,
		Description: payload.Description,
		Tags:        payload.Tags,
	}

	ctx := r.Context()

	videoData, err := h.Store.GetVideoById(ctx, video.ID)
	if err != nil {
		h.Logger.Errorf("error fetching video of id-%i: %v", video.ID, err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch video",
		})
		return nil
	}

	return videoData

}
