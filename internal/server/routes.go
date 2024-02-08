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

	r.Route("/users", UserRoutes(db.DB))

	return r
}
func UserRoutes(db *sql.DB) func(r chi.Router) {
	return func(r chi.Router) {
		userDB := database.NewUserDB(db)
		userHandler := handlers.NewUserHandler(userDB)
		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetUsers)
		r.Get("/search", userHandler.GetUserByEmail)
		r.Get("/{id}", userHandler.GetUserByID)
	}
}
