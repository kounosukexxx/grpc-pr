package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shota-aa/grpc-pr/domain"
	"github.com/shota-aa/grpc-pr/usecase/repository"
)

type UserService interface {
	// GetUsers(ctx context.Context) ([]*domain.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, arg *repository.CreateUserArg) (*domain.User, error)
}

type userService struct {
	user repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{user}
}
// func (s *userService) GetUsers(ctx context.Context) ([]*domain.User, error) {
// 	users, err := s.repo.GetUsers(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

func (s *userService) GetUser(ctx context.Context, userId uuid.UUID) (*domain.User, error) {
	user, err := s.user.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, arg *repository.CreateUserArg) (*domain.User, error) {
	user, err := s.user.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return user, nil
}
