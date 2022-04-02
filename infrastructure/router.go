package infrastructure

import (
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/shota-aa/grpc-pr/interfaces/handler"
	pb "github.com/shota-aa/grpc-pr/pb/proto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func SetupGRPCServer(h handler.GRPC) *grpc.Server {
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
	s := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(zap, zap_opt),
	))
	pb.RegisterUserServiceServer(s, h.User)
	reflection.Register(s)
	
	return s
}
