package main

import (
	"chatapp/internal/config"
	"chatapp/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Get("/", handlers.Repo.Home)
	return router
}
