package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	man "github.com/KeatonBrink/MancalaGame/src/protos"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	man.UnimplementedMancalaServiceServer
	waitingRoom    []string
	allUserNames   []string
	gameInProgress []MancalaGameBoard
	mu             sync.Mutex
}

type MancalaGameBoard struct {
	p1Name string
	p1Row  [7]int
	p2Name string
	p2Row  [7]int
}

func (mgb *MancalaGameBoard) resetBoard() {
	for i := 0; i < 6; i++ {
		mgb.p1Row[i] = 4
		mgb.p2Row[i] = 4
	}
	mgb.p1Row[6] = 0
	mgb.p2Row[6] = 0
}

func (s *server) GameHandshake(ctx context.Context, req *man.HandshakeRequest) (*man.HandshakeResponse, error) {
	log.Printf("Received: %v", req.GetUserName())
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.allUserNames {
		if v == req.GetUserName() {
			return &man.HandshakeResponse{
				ErrorCode:              1,
				ErrorMessage:           "Name already in use",
				Message:                "",
				ServerWebSocketAddress: "",
			}, nil
		}
	}
	s.allUserNames = append(s.allUserNames, req.GetUserName())
	if len(s.waitingRoom) > 0 {
		return &man.HandshakeResponse{
			ErrorCode:              0,
			ErrorMessage:           "",
			Message:                "Game Starting Soon",
			ServerWebSocketAddress: "",
		}, nil
	} else {
		s.waitingRoom = append(s.waitingRoom, req.GetUserName())
		return &man.HandshakeResponse{
			ErrorCode:              0,
			ErrorMessage:           "",
			Message:                "Waiting for Second Player",
			ServerWebSocketAddress: "8081",
		}, nil
	}
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
