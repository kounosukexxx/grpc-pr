package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/shota-aa/grpc-pr/domain"
)

type CreateUserArg struct {
	Name  string
	Email string
}

type UserRepository interface {
	// GetUsers(ctx context.Context) ([]*domain.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (*domain.User, error)
	CreateUser(ctx context.Context, arg *CreateUserArg) (*domain.User, error)
	GetUsersByIDs(ctx context.Context, userIds []*uuid.UUID) ([]*domain.User, error)
}
