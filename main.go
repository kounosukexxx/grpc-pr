package main

import (
	"log"
	"net"

	"github.com/shota-aa/grpc-pr/config"
	"github.com/shota-aa/grpc-pr/infrastructure"
)

func main() {
	config := config.GetConfig()
	grpcHander, err := infrastructure.InjectGRPCServer(config)
	if err != nil {
		panic("failed to inject grpc server")
	}
	s := infrastructure.SetupGRPCServer(grpcHander)

	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %v", ":"+config.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}