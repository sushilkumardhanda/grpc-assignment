package repository

import (
	"testing"

	pb "grpc-user/proto"

	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
	repo := NewInMemoryUserRepository()

	user, err := repo.FindByID("0")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Sushil", user.Name)

	_, err = repo.FindByID("unknown")
	assert.Error(t, err)
}

func TestFindByIDs(t *testing.T) {
	repo := NewInMemoryUserRepository()

	users, err := repo.FindByIDs([]string{"0", "1"})
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	users, err = repo.FindByIDs([]string{"0", "unknown"})
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestSearch(t *testing.T) {
	repo := NewInMemoryUserRepository()

	searchCriteria := &pb.User{Name: "Alice"}
	users, err := repo.Search(searchCriteria)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "Alice", users[0].Name)

	searchCriteria = &pb.User{City: "Delhi"}
	users, err = repo.Search(searchCriteria)
	assert.NoError(t, err)
	assert.Len(t, users, 4)
}
