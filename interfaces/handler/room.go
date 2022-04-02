package handler

import (
	"context"

	pb "github.com/shota-aa/grpc-pr/pb/proto"
	"github.com/shota-aa/grpc-pr/usecase/service"
)

type RoomHandler struct {
	repo service.RoomService
}

func NewRoomHandler(service service.RoomService) *RoomHandler {
	return &RoomHandler{service}
}

func (h *RoomHandler) GetRoom(ctx context.Context, req *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
	return nil, nil
}

func (h *RoomHandler) GetRooms(ctx context.Context, req *pb.GetRoomsRequest) (*pb.GetRoomsResponse, error) {
	return nil, nil
}

func (h *RoomHandler) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	return nil, nil
}

func (g *RoomHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	return nil, nil
}