package repositories

import (
	sqlc "api-management/internal/db/sqlc"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type APIKeyRepository struct {
	q *sqlc.Queries
}

type APIKey struct {
	ID         string
	UserID     string
	Name       string
	CreatedAt  time.Time
	RevokedAt  *time.Time
	LastUsedAt *time.Time
}

func pgTimeToPtr(t pgtype.Timestamp) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

func NewAPIKeyRepository(pool *pgxpool.Pool) *APIKeyRepository {
	return &APIKeyRepository{
		q: sqlc.New(pool),
	}
}

func (a *APIKeyRepository) CreateApiKey(ctx context.Context, user_id, name, key_hash string) (APIKey, error) {
	var userUUId pgtype.UUID
	err := userUUId.Scan(user_id)
	if err != nil {
		return APIKey{}, err
	}

	dbKey, err := a.q.CreateApiKey(ctx, sqlc.CreateApiKeyParams{
		UserID:  userUUId,
		Name:    name,
		KeyHash: key_hash,
	})

	if err != nil {
		return APIKey{}, err
	}

	return APIKey{
		ID:         dbKey.ID.String(),
		UserID:     dbKey.UserID.String(),
		Name:       dbKey.Name,
		CreatedAt:  dbKey.CreatedAt.Time,
		RevokedAt:  pgTimeToPtr(dbKey.RevokedAt),
		LastUsedAt: pgTimeToPtr(dbKey.LastUsedAt),
	}, nil
}
