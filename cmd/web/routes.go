package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sseudes108/go-course/pkg/config"
	"github.com/sseudes108/go-course/pkg/handlers"
)

// routes manages the templates routes
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
