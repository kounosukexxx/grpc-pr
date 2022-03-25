package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	user "github.com/shota-aa/grpc-pr/pb/rest"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// logging
	zap, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to set: %v", err)
	}
	zap_opt := grpc_zap.WithLevels( // --- ②
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel

			case codes.Internal:
				l = zapcore.ErrorLevel

			default:
				l = zapcore.DebugLevel
			}
			return l
		},
	)

	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(zap, zap_opt),
	))
	user.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("Listening on %v", ":"+port)
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
