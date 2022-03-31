package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shota-aa/grpc-pr/domain"
	"github.com/shota-aa/grpc-pr/usecase/repository"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]*domain.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, arg *repository.CreateUserArg) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
} 

func (s *userService) GetUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUser(ctx context.Context, userId uuid.UUID) (*domain.User, error) {
	return nil, nil
}

func (s *userService) CreateUser(ctx context.Context, arg *repository.CreateUserArg) (*domain.User, error) {
	return nil, nil
}