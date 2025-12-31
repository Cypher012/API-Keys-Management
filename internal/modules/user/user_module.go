package user

import (
	"api-management/internal/db/repositories"
	"api-management/internal/handlers"
	"api-management/internal/routes"
	"api-management/internal/services"
	"api-management/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(r chi.Router, pool *pgxpool.Pool, jwt *utils.JWTManager) {
	repo := repositories.NewUserRepository(pool)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service, jwt)

	routes.RegisterUserRoutes(r, handler)
}
