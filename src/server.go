package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	man "github.com/KeatonBrink/MancalaGame/src/protos"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	man.UnimplementedMancalaServiceServer
}

func (s *server) GameHandshake(ctx context.Context, in *man.HandshakeRequest) (*man.HandshakeResponse, error) {
	log.Printf("Received: %v", in.GetUserName())
	return &man.HandshakeResponse{Message: "Hello " + in.GetUserName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	man.RegisterMancalaServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
