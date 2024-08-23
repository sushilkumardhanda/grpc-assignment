package service

import (
	"testing"

	pb "grpc-user/proto"
	"grpc-user/repository"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := NewUserService(repo)

	user, err := service.GetUserByID("0")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Sushil", user.Name)

	_, err = service.GetUserByID("unknown")
	assert.Error(t, err)
}

func TestGetUsersByIDs(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := NewUserService(repo)

	users, err := service.GetUsersByIDs([]string{"0", "1"})
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	users, err = service.GetUsersByIDs([]string{"0", "unknown"})
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestSearchUsers(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := NewUserService(repo)

	searchCriteria := &pb.User{Name: "Alice"}
	users, err := service.SearchUsers(searchCriteria)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "Alice", users[0].Name)

	searchCriteria = &pb.User{City: "Delhi"}
	users, err = service.SearchUsers(searchCriteria)
	assert.NoError(t, err)
	assert.Len(t, users, 4)
}
