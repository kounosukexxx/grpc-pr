package infrastructure

import (
	"github.com/shota-aa/grpc-pr/interfaces/handler"
	pb "github.com/shota-aa/grpc-pr/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Setup(s *grpc.Server, grpc handler.GRPC) {
	pb.RegisterUserServiceServer(s, grpc.User)
	reflection.Register(s)
}