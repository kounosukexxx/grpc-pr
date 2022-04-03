package repository

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/shota-aa/grpc-pr/domain"
	"github.com/shota-aa/grpc-pr/interfaces/repository/model"
	"github.com/shota-aa/grpc-pr/usecase/repository"
)

type RoomRepository struct {
	client *firestore.Client
}

func NewRoomRepository(client *firestore.Client) repository.RoomRepository {
	return &RoomRepository{client: client}
}

func (repo *RoomRepository) CreateRoom(ctx context.Context, arg *repository.CreateRoomArg) (*domain.Room, error) {
	ID := uuid.New()
	err := repo.client.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		roomRef := repo.client.Collection("rooms").Doc(ID.String())
		err := t.Set(roomRef, map[string]interface{}{
			"id":         ID.String(),
			"name":       arg.Name,
			"created_at": firestore.ServerTimestamp,
			"updated_at": firestore.ServerTimestamp,
		})
		if err != nil {
			return err
		}
		for _, userId := range arg.UserIds {
			userRef := repo.client.Collection("rooms").
				Doc(ID.String()).
				Collection("users").
				Doc(userId.String())
			err = t.Set(userRef, map[string]interface{}{
				"id": userId,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	// 取らなくてもいけるが
	doc, err := repo.client.Collection("rooms").
		Doc(ID.String()).
		Get(ctx)
	if err != nil {
		return nil, err
	}
	var room model.Room
	if err = doc.DataTo(&room); err != nil {
		return nil, err
	}
	return &domain.Room{
		Id:        ID,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}, nil
}

func (repo *RoomRepository) CreateComment(ctx context.Context, arg *repository.CreateCommentArg) (*domain.Comment, error) {
	return nil, nil
}
