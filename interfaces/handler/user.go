package handler

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/shota-aa/grpc-pr/pb/proto"
	
)

type UserHandler struct{}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received: %v", req.Id)
	now, _ := ptypes.TimestampProto(time.Now())
	return &pb.GetUserResponse{Id: req.Id, Name: "John Smith", Email: "johnsmith@example.com", UpdatedAt: now}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, nil
}
