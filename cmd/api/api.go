package main

import (
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/qlfzn/viewlater/config"
	"github.com/qlfzn/viewlater/internal/handlers"
)

type application struct {
	config  config.Config
	logger  *zap.SugaredLogger
	handler *handlers.Handler
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Add prefix in related routes
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Post("/videos", app.handler.SaveVideoHandler)
		r.Get("/videos?id", app.handler.GetVideoHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:    app.config.PORT,
		Handler: mux,
	}

	log.Printf("server has started at http://localhost:8080")

	return srv.ListenAndServe()
}
