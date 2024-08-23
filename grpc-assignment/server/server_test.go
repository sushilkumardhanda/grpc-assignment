package server

import (
	"context"
	"net"
	"testing"

	"grpc-user/service"
	"grpc-user/repository"
	pb "grpc-user/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	// Dependency Injection
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userServer := NewUserServer(userService)

	pb.RegisterUserServiceServer(s, userServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	req := &pb.UserIDs{Ids: []string{"0"}}
	resp, err := client.GetUserByID(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Sushil", resp.Name)

	req = &pb.UserIDs{Ids: []string{"unknown"}}
	_, err = client.GetUserByID(ctx, req)
	assert.Error(t, err)
}

func TestGetUsersByIDs(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	stream, err := client.GetUsersByIDs(ctx, &pb.UserIDs{Ids: []string{"0", "1"}})
	assert.NoError(t, err)

	var users []*pb.User
	for {
		user, err := stream.Recv()
		if err != nil {
			break
		}
		users = append(users, user)
	}

	assert.Len(t, users, 2)
}

func TestSearchUsers(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	stream, err := client.SearchUsers(ctx, &pb.User{Name: "Alice"})
	assert.NoError(t, err)

	var users []*pb.User
	for {
		user, err := stream.Recv()
		if err != nil {
			break
		}
		users = append(users, user)
	}

	assert.Len(t, users, 1)
	assert.Equal(t, "Alice", users[0].Name)
}
