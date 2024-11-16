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
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	app := &application{
		logger: logger,
	}
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	logger.Info("Starting Server", slog.String("address", *address))

	err := http.ListenAndServe(*address, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
