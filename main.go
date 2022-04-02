package main

import (
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/shota-aa/grpc-pr/config"
	"github.com/shota-aa/grpc-pr/infrastructure"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
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
				l = zapcore.ErrorLevel
			}
			return l
		},
	)

	config := config.GetConfig()
	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcHander, err := infrastructure.InjectGRPCServer(config)
	if err != nil {
		panic("failed to inject grpc server")
	}
	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(zap, zap_opt),
	))
	infrastructure.Setup(s, grpcHander)
	log.Printf("Listening on %v", ":"+config.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
