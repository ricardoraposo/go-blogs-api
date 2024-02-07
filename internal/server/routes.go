package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ricardoraposo/blogs-api-go/internal/database"
)

func NewRouter() *chi.Mux {
	db := database.NewDatabase()
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", db.HealthCheck)

	return r
}
