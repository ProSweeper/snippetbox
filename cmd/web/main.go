package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	address := flag.String("address", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("Starting Server", slog.String("address", *address))
	str := "sting"

	err := http.ListenAndServe(*address, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
