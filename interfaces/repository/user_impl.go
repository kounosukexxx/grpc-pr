package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/shota-aa/grpc-pr/domain"
	"github.com/shota-aa/grpc-pr/usecase/repository"
)

type UserRepository struct {
	client *firestore.Client
}

func NewUserRepository(client *firestore.Client) repository.UserRepository {
	return &UserRepository{client: client}
}

func (repo *UserRepository) GetUser(ctx context.Context, userId int) (*domain.User, error) {
	data, err := repo.client.Collection("users").Doc(fmt.Sprint(userId)).Get(ctx)
	if err != nil {
		return nil, err
	}
	user := new(domain.User)
	mapToUser(data.Data(), &user)
	return user, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, arg *repository.CreateUserArg) (*domain.User, error) {
	return nil, nil
}

func mapToUser(userMap map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(userMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
