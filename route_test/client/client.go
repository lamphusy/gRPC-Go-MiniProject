// client.go

package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "grpc-test/route_test/route/routeUser/pb"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewRouteUserClient(conn)

	// Create a request (assuming Profile struct is properly populated)
	profile := &pb.Profile{}

	// Call the GetUser RPC
	reply, err := client.GetUser(context.Background(), profile)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	// Print the response
	fmt.Printf("Message: %s\n", reply.Message)
	fmt.Printf("User Information: %+v\n", reply.Information)
}
