package routes

import (
	"api-management/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, h *handlers.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUserHandler)
		r.Get("/", h.GetUserByEmailHandler)
	})

}

func RegisterAPIKeysRoutes(r chi.Router, h *handlers.APIKeyHandler) {
	r.Route("/api-keys", func(r chi.Router) {
		r.Post("/", h.CreateAPIKeyHandler)
	})
}
