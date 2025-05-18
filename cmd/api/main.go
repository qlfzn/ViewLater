package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/qlfzn/viewlater/internal/handlers"
)

func main() {
	cfg := config{
		":8080",
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// initialise handler
	handler := &handlers.Handler{
		Logger: logger,
	}

	app := application{
		cfg,
		logger,
		handler,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
