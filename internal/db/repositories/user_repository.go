package repositories

import (
	"context"
	"errors"

	sqlc "api-management/internal/db/sqlc"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUserNotFound = errors.New("user not found")
var ErrUserAlreadyExists = errors.New("user already exists")

const pgUniqueViolation = "23505"

type UserRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		q: sqlc.New(pool),
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, password, email string) (sqlc.User, error) {
	user, err := u.q.CreateUser(ctx, sqlc.CreateUserParams{
		PasswordHash: password,
		Email:        email,
	})
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == pgUniqueViolation {
			return sqlc.User{}, ErrUserAlreadyExists
		}
		return sqlc.User{}, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (sqlc.User, error) {
	user, err := u.q.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sqlc.User{}, ErrUserNotFound
		}
		return sqlc.User{}, err
	}
	return user, nil
}
