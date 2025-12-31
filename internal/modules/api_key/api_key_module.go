package api_key

import (
	"api-management/internal/db/repositories"
	"api-management/internal/handlers"
	"api-management/internal/routes"
	"api-management/internal/services"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(r chi.Router, pool *pgxpool.Pool) {
	apiSecret := os.Getenv("API_KEY_SECRET")

	if apiSecret == "" {
		log.Fatal("NO api secret key")
	}
	repo := repositories.NewAPIKeyRepository(pool)
	service := services.NewAPIKeyService(repo, apiSecret)
	handler := handlers.NewAPIKeyHandler(service)

	routes.RegisterAPIKeysRoutes(r, handler)

}
