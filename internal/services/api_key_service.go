package services

import (
	"api-management/internal/db/repositories"
	"api-management/internal/utils"
	"context"
	"errors"
)

type APIKeyService struct {
	repo      *repositories.APIKeyRepository
	apiSecret string
}

type CreatedAPIKey struct {
	ApiKey repositories.APIKey
	RawKey string
}

func NewAPIKeyService(repo *repositories.APIKeyRepository, apiSecret string) *APIKeyService {
	return &APIKeyService{
		repo:      repo,
		apiSecret: apiSecret,
	}
}

func (s *APIKeyService) CreateApiKey(ctx context.Context, user_id, name string) (*CreatedAPIKey, error) {
	if s.apiSecret == "" {
		return nil, errors.New("API_KEY_SECRET Missing")
	}

	rawKey := utils.GenerateAPIKey()
	key_hash := utils.HashAPIKey(rawKey, s.apiSecret)

	apiKey, err := s.repo.CreateApiKey(ctx, user_id, name, key_hash)
	if err != nil {
		return nil, err
	}

	return &CreatedAPIKey{
		ApiKey: apiKey,
		RawKey: rawKey,
	}, nil
}
