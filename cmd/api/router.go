package main

import (
	"api-management/internal/modules/api_key"
	"api-management/internal/modules/user"
	"api-management/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(dbPool *pgxpool.Pool, jwt *utils.JWTManager) chi.Router {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// API routes
	r.Route("/api", func(api chi.Router) {
		// Public routes (no auth required)
		user.Register(api, dbPool, jwt)

		// Protected routes
		api.Group(func(protected chi.Router) {
			protected.Use(jwt.Middleware)
			api_key.Register(protected, dbPool)
		})
	})

	return r
}
