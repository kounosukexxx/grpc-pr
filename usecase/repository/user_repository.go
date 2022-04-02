package repository

import (
	"context"

	"github.com/shota-aa/grpc-pr/domain"
)

type CreateUserArg struct {
	Name  string
	Email string
}

type UserRepository interface {
	// GetUsers(ctx context.Context) ([]*domain.User, error)
	GetUser(ctx context.Context, userId int) (*domain.User, error)
	CreateUser(ctx context.Context, arg *CreateUserArg) (*domain.User, error)
}
