package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	cfg := config{
		":8080",
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	app := application{
		cfg,
		logger,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
