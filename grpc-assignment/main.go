package main

import (
	"log"
	"net"

	"grpc-user/repository"
	"grpc-user/server"
	"grpc-user/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "grpc-user/proto" // Import your protobuf package
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Dependency Injection
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userServer := server.NewUserServer(userService)

	pb.RegisterUserServiceServer(s, userServer)
	reflection.Register(s)
	log.Println("gRPC server started on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
