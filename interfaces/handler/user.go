package handler

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	pb "github.com/shota-aa/grpc-pr/pb/proto"
	"github.com/shota-aa/grpc-pr/usecase/repository"
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
	ID, err := uuid.Parse(req.Id) 
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "received 0 or nothing")
	}
	user, err := h.repo.GetUser(ctx, ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	createdAt, _ := ptypes.TimestampProto(user.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(user.UpdatedAt)
	return &pb.GetUserResponse{
		Id: ID.String(),
		Name: user.Name,
		Email: user.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req == nil || len(req.Name) == 0 || len(req.Email) == 0 {
		return nil, status.Error(codes.InvalidArgument, "bad request")
	}
	user, err := h.repo.CreateUser(ctx, &repository.CreateUserArg{
		Name: req.Name, Email: req.Email,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	createdAt, _ := ptypes.TimestampProto(user.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(user.UpdatedAt)
	return &pb.CreateUserResponse{
		Id: user.Id.String(),
		Name: user.Name,
		Email: user.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
