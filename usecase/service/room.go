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
	room repository.RoomRepository
	user repository.UserRepository
}

func NewRoomService(room repository.RoomRepository, user repository.UserRepository) RoomService {
	return &roomService{room, user}
}

func (s *roomService) CreateRoom(ctx context.Context, arg *repository.CreateRoomArg) (*domain.Room, error) {
	users, err := s.user.GetUsersByIDs(ctx, arg.UserIds)
	if err != nil {
		return nil, err
	}
	room, err := s.room.CreateRoom(ctx, arg)
	if err != nil {
		return nil, err
	}
	room.Users = users
	return room, nil
}

func (s *roomService) CreateComment(ctx context.Context, arg *repository.CreateCommentArg) (*domain.Comment, error) {
	return nil, nil
}