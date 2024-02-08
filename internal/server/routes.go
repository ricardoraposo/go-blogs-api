package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ricardoraposo/blogs-api-go/internal/database"
	"github.com/ricardoraposo/blogs-api-go/internal/handlers"
)

func NewRouter() *chi.Mux {
	db := database.NewDatabase()
	r := chi.NewRouter()

	userDB := database.NewUserDB(db.DB)
    userHandler := handlers.NewUserHandler(userDB)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", db.HealthCheck)

	r.Post("/users", userHandler.CreateUser)
    r.Get("/users", userHandler.GetUsers)
    r.Get("/users/search", userHandler.GetUserByEmail)
    r.Get("/users/{id}", userHandler.GetUserByID)

	return r
}
