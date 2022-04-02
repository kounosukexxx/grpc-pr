package service

import (
	"context"

	"github.com/shota-aa/grpc-pr/domain"
	"github.com/shota-aa/grpc-pr/usecase/repository"
)

type RoomService interface {
	CreateRoom(ctx context.Context, arg *repository.CreateRoomArg) (*domain.Room, error)
	CreateComment(ctx context.Context, arg *repository.CreateCommentArg) (*domain.Comment, error)
}

type roomService struct {
	repo repository.RoomRepository
}

func NewRoomService(room repository.RoomRepository) RoomService {
	return &roomService{room}
}

func (s *roomService) CreateRoom(ctx context.Context, arg *repository.CreateRoomArg) (*domain.Room, error) {
	return nil, nil
}

func (s *roomService) CreateComment(ctx context.Context, arg *repository.CreateCommentArg) (*domain.Comment, error) {
	return nil, nil
}