package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/qlfzn/viewlater/config"
	"github.com/qlfzn/viewlater/internal/handlers"
	"github.com/qlfzn/viewlater/internal/repository"
	"github.com/qlfzn/viewlater/internal/store"
)

func main() {
	// Load config
	if err := config.Init(); err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	cfg := config.AppConf

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// initialise db
	db := repository.New(cfg.DSN())

	// intialise store service
	srv := store.VideoStore{
		DB: db,
	}

	// initialise handler
	handler := &handlers.Handler{
		Logger: logger,
		Store:  srv,
	}

	app := application{
		cfg,
		logger,
		handler,
	}

	// start app
	mux := app.mount()
	log.Fatal(app.run(mux))
}
