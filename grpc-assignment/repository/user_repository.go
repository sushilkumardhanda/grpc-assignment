package repository

import (
	"fmt"
	pb "grpc-user/proto"
)

type UserRepository interface {
	FindByID(id string) (*pb.User, error)
	FindByIDs(ids []string) ([]*pb.User, error)
	Search(user *pb.User) ([]*pb.User, error)
}

type inMemoryUserRepository struct {
	users []*pb.User
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: []*pb.User{
			{Id: "0", Name: "Sushil", Email: "abc@c.com", City: "Delhi", Phone: "7742311", Height: 5.9, Married: true},
			{Id: "1", Name: "Alice", Email: "alice@example.com", City: "Delhi", Phone: "7742311", Height: 5.9, Married: true},
			{Id: "2", Name: "Bob", Email: "bob@example.com", City: "Delhi", Phone: "7742311", Height: 5.9, Married: true},
			{Id: "3", Name: "Charlie", Email: "charlie@example.com", City: "Delhi", Phone: "7742311", Height: 5.9, Married: true},
		},
	}
}

func (r *inMemoryUserRepository) FindByID(id string) (*pb.User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

func (r *inMemoryUserRepository) FindByIDs(ids []string) ([]*pb.User, error) {
	var users []*pb.User
	for _, id := range ids {
		for _, user := range r.users {
			if user.Id == id {
				users = append(users, user)
			}
		}
	}
	return users, nil
}

func (r *inMemoryUserRepository) Search(user *pb.User) ([]*pb.User, error) {
	var users []*pb.User
	for _, u := range r.users {
		if (user.Name == "" || u.Name == user.Name) &&
			(user.Email == "" || u.Email == user.Email) &&
			(user.City == "" || u.City == user.City) &&
			(user.Phone == "" || u.Phone == user.Phone) {
			users = append(users, u)
		}
	}
	return users, nil
}
