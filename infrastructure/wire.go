//go:generate go run github.com/google/wire/cmd/wire@v0.5.0
//go:build wireinject

package infrastructure

import (
	"github.com/google/wire"
	"github.com/shota-aa/grpc-pr/config"
	"github.com/shota-aa/grpc-pr/interfaces/handler"
	impl "github.com/shota-aa/grpc-pr/interfaces/repository"
	"github.com/shota-aa/grpc-pr/usecase/service"
)

var firestoreSet = wire.NewSet(NewFirestoreClient)

var userSet = wire.NewSet(
	impl.NewUserRepository,
	service.NewUserService,
	handler.NewUserHandler,
)

var roomSet = wire.NewSet(
	impl.NewRoomRepository,
	service.NewRoomService,
	handler.NewRoomHandler,
)

var grpcSet = wire.NewSet(handler.NewGRPC)

func InjectGRPCServer(c *config.Config) (handler.GRPC, error) {
	wire.Build(
		firestoreSet,
		userSet,
		roomSet,
		grpcSet,
	)
	return handler.GRPC{}, nil
}
