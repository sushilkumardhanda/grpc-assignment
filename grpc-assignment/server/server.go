package server

import (
	"context"

	"grpc-user/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "grpc-user/proto"
)

type userServer struct {
	service service.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserServer(service service.UserService) pb.UserServiceServer {
	return &userServer{service: service}
}

func (s *userServer) GetUserByID(ctx context.Context, req *pb.UserIDs) (*pb.User, error) {
	user, err := s.service.GetUserByID(req.Ids[0])
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}
	return user, nil
}

func (s *userServer) GetUsersByIDs(req *pb.UserIDs, stream pb.UserService_GetUsersByIDsServer) error {
	users, err := s.service.GetUsersByIDs(req.Ids)
	if err != nil {
		return err
	}
	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (s *userServer) SearchUsers(req *pb.User, stream pb.UserService_SearchUsersServer) error {
	users, err := s.service.SearchUsers(req)
	if err != nil {
		return err
	}
	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}
