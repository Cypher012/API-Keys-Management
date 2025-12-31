package services

import (
	"api-management/internal/db/repositories"
	"api-management/internal/utils"
	"context"
)

type UserService struct {
	repo *repositories.UserRepository
}

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, password, email string) (*User, error) {
	hashed_password, err := utils.GenerateHashFromPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.CreateUser(ctx, hashed_password, email)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:    user.ID.String(),
		Email: user.Email,
	}, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    user.ID.String(),
		Email: user.Email,
	}, nil
}
