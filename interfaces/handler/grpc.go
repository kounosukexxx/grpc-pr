package handler

type GRPC struct {
	User *UserHandler
}

func NewGRPC(user *UserHandler) GRPC {
	return GRPC{
		User: user,
	}
}