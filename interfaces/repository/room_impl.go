package repository

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	"github.com/shota-aa/grpc-pr/domain"
	"github.com/shota-aa/grpc-pr/usecase/repository"
)

type RoomRepository struct {
	client *firestore.Client
}

func NewRoomRepository(client *firestore.Client) repository.RoomRepository {
	return &RoomRepository{client: client}
}

func (repo *RoomRepository) CreateRoom(ctx context.Context, arg *repository.CreateRoomArg) (*domain.Room, error) {
	return nil, nil
}

func (repo *RoomRepository) CreateComment(ctx context.Context, arg *repository.CreateCommentArg) (*domain.Comment, error) {
	return nil, nil
}