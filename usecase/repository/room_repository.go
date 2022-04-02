package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/shota-aa/grpc-pr/domain"
)

type CreateRoomArg struct {
	Name    string
	UserIds []*uuid.UUID
}

type CreateCommentArg struct {
	RoomId  uuid.UUID
	Comment string
}

type RoomRepository interface {
	CreateRoom(ctx context.Context, arg *CreateRoomArg) (*domain.Room, error)
	CreateComment(ctx context.Context, arg*CreateCommentArg) (*domain.Comment, error)
}
