package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-test/route_test/routetest/routeUser/pb"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

type serverUser struct {
	pb.UnimplementedRouteUserServer
	savedReply []*pb.Reply
}

// GetUser return User's Information
func (s *serverUser) GetUser(ctx context.Context, profile *pb.Profile) (*pb.Reply, error) {

	// for _, reply := range s.savedReply {
	// 	if proto.Equal(reply.Information, profile) {
	// 		return reply, nil
	// 	}
	// }
	// return &pb.Reply{Information: profile}, nil

	reply := &pb.Reply{
		Message: "Userdata retrieve successfully",
		Information: &pb.Profile{
			Id:       12356,
			Name:     "Erling Haaland",
			Email:    "haaland@gmail.com",
			Mobile:   "0971820510",
			Password: "haaland123",
		},
	}
	return reply, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteUserServer(grpcServer, &serverUser{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
