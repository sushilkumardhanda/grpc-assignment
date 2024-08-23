package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "grpc-user/proto" // Import your protobuf package
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Example: Calling GetUserByID
	resp, err := client.GetUserByID(context.Background(), &pb.UserIDs{Ids: []string{"1"}})
	if err != nil {
		log.Fatalf("Error calling GetUserByID: %v", err)
	}
	log.Printf("User found: %v", resp)

	users, err := client.GetUsersByIDs(context.Background(), &pb.UserIDs{Ids: []string{"1", "2"}})
	if err != nil {
		log.Fatalf("Error calling GetUsersByIDs: %v", err)
	}
	for {
		user, err := users.Recv()
		if err != nil {
			break
		}
		log.Printf("User found: %v", user)
	}
	// Example: Calling SearchUsers
	searchResp, err := client.SearchUsers(context.Background(), &pb.User{Name: "Sushil", City: "Delhi"})
	if err != nil {
		log.Fatalf("Error calling SearchUsers: %v", err)
	}
	for {
		user, err := searchResp.Recv()
		if err != nil {
			break
		}
		log.Printf("User found: %v", user)
	}
}
