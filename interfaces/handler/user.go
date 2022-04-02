package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/shota-aa/grpc-pr/pb/proto"
	"github.com/shota-aa/grpc-pr/usecase/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	repo service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "received 0 or nothing")
	}
	user, err := h.repo.GetUser(ctx, int(req.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}
	createdAt, _ := ptypes.TimestampProto(user.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(user.UpdatedAt)
	return &pb.GetUserResponse{Id: int32(user.Id), Name: user.Name, Email: user.Email, CreatedAt: createdAt, UpdatedAt: updatedAt}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, nil
}
