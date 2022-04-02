package handler

type GRPC struct {
	User *UserHandler
	Room *RoomHandler
}

func NewGRPC(user *UserHandler, room *RoomHandler) GRPC {
	return GRPC{
		User: user,
		Room: room,
	}
}