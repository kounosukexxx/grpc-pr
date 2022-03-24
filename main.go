package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"
	user "github.com/shota-aa/grpc-pr/pb/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("Listening on %v", ":8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {
	log.Printf("Received: %v", req.Id)
	now, _ := ptypes.TimestampProto(time.Now())
	return &user.User{Id: req.Id, Name: "John Smith", Email: "johnsmith@example.com", UpdatedAt: now}, nil
}
