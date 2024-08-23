package service

import (
	"grpc-user/repository"
	pb "grpc-user/proto"
)

type UserService interface {
	GetUserByID(id string) (*pb.User, error)
	GetUsersByIDs(ids []string) ([]*pb.User, error)
	SearchUsers(user *pb.User) ([]*pb.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserByID(id string) (*pb.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetUsersByIDs(ids []string) ([]*pb.User, error) {
	return s.repo.FindByIDs(ids)
}

func (s *userService) SearchUsers(user *pb.User) ([]*pb.User, error) {
	return s.repo.Search(user)
}
