package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/shota-aa/grpc-pr/domain"
	"github.com/shota-aa/grpc-pr/interfaces/repository/model"
	"github.com/shota-aa/grpc-pr/usecase/repository"
)

type UserRepository struct {
	client *firestore.Client
}

func NewUserRepository(client *firestore.Client) repository.UserRepository {
	return &UserRepository{client: client}
}

func (repo *UserRepository) GetUser(ctx context.Context, userId uuid.UUID) (*domain.User, error) {
	data, err := repo.client.Collection("users").
		Doc(fmt.Sprint(userId)).
		Get(ctx)
	if err != nil {
		return nil, err
	}
	var user model.User
	if err = data.DataTo(&user); err != nil {
		return nil, err
	}
	return &domain.User{
		Id:        userId,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, arg *repository.CreateUserArg) (*domain.User, error) {
	ID := uuid.New()
	_, err := repo.client.Collection("users").
		Doc(ID.String()).
		Set(ctx, map[string]interface{}{
			"id":         ID.String(),
			"name":       arg.Name,
			"email":      arg.Email,
			"created_at": firestore.ServerTimestamp,
			"updated_at": firestore.ServerTimestamp,
		})
	if err != nil {
		return nil, err
	}
	data, err := repo.client.Collection("users").
		Doc(ID.String()).
		Get(ctx)
	if err != nil {
		return nil, err
	}
	var user model.User
	if err = data.DataTo(&user); err != nil {
		return nil, err
	}
	return &domain.User{
		Id:        ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// func mapToUser(userMap map[string]interface{}, val interface{}) error {
// 	tmp, err := json.Marshal(userMap)
// 	if err != nil {
// 		return err
// 	}
// 	err = json.Unmarshal(tmp, val)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
