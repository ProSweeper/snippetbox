package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	address := flag.String("address", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("Starting Server", slog.String("address", *address))

	err := http.ListenAndServe(*address, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}
