package handler

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	pb "github.com/shota-aa/grpc-pr/pb/proto"
	"github.com/shota-aa/grpc-pr/usecase/repository"
	"github.com/shota-aa/grpc-pr/usecase/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	var userIds []*uuid.UUID
	for _, user := range req.Users {
		userId, err := uuid.Parse(user)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		userIds = append(userIds, &userId)
	}
	room, err := h.repo.CreateRoom(ctx, &repository.CreateRoomArg{
		Name:    req.Name,
		UserIds: userIds,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	createdAt, _ := ptypes.TimestampProto(room.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(room.UpdatedAt)
	var resUsers []*pb.User
	for _, u := range room.Users {
		createdAt, _ := ptypes.TimestampProto(u.CreatedAt)
		updatedAt, _ := ptypes.TimestampProto(u.UpdatedAt)
		resUsers = append(resUsers, &pb.User{
			Id:        u.Id.String(),
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return &pb.CreateRoomResponse{
		Id:        room.Id.String(),
		Name:      room.Name,
		Users:     resUsers,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (h *RoomHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	roomId, err := uuid.Parse(req.RoomId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	comment, err := h.repo.CreateComment(ctx, &repository.CreateCommentArg{
		RoomId:  roomId,
		Comment: req.Comment,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	createdAt, _ := ptypes.TimestampProto(comment.User.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(comment.User.UpdatedAt)
	return &pb.CreateCommentResponse{
		Id:      comment.Id.String(),
		RoomId:  comment.RoomId.String(),
		Comment: comment.Comment,
		CreatedBy: &pb.User{
			Id:        comment.User.Id.String(),
			Name:      comment.User.Name,
			Email:     comment.User.Email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}, nil
}
