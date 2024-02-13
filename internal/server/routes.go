package server

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ricardoraposo/blogs-api-go/internal/database"
	"github.com/ricardoraposo/blogs-api-go/internal/handlers"
	"github.com/ricardoraposo/blogs-api-go/internal/middlewares"
)

func NewRouter() *chi.Mux {
	db := database.NewDatabase()
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", db.HealthCheck)

	r.Route("/users", UserRoutes(db.DB))
	r.Route("/categories", CategoryRoutes(db.DB))
	r.Route("/posts", BlogPostRoutes(db.DB))

	return r
}

func UserRoutes(db *sql.DB) func(r chi.Router) {
	return func(r chi.Router) {
		userDB := database.NewUserDB(db)
		userHandler := handlers.NewUserHandler(userDB)

		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetUsers)
		r.Use(middlewares.VerifyJWT)
		r.Get("/search", userHandler.GetUserByEmail)
		r.Get("/{id}", userHandler.GetUserByID)
	}
}

func CategoryRoutes(db *sql.DB) func(r chi.Router) {
	return func(r chi.Router) {
		categoryDB := database.NewCategoryDB(db)
		categoryHandler := handlers.NewCategoryHandler(categoryDB)

		r.Use(middlewares.VerifyJWT)
		r.Post("/", categoryHandler.CreateCategory)
		r.Get("/", categoryHandler.GetCategories)
	}
}

func BlogPostRoutes(db *sql.DB) func(r chi.Router) {
	return func(r chi.Router) {
		blogPostDB := database.NewBlogPostDB(db)
		postCategoryDB := database.NewPostCategoryDB(db)
		blogPostHandler := handlers.NewBlogPostHandler(blogPostDB, postCategoryDB)

		r.Use(middlewares.VerifyJWT)
		r.Get("/", blogPostHandler.GetBlogPosts)
		r.Post("/", blogPostHandler.CreateBlogPost)
	}
}
