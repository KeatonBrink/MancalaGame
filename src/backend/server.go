package main

import (
	"flag"
	"fmt"
	"hash/fnv"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	man "github.com/KeatonBrink/MancalaGame/src/protos/generated"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all cross-origin requests
	},
}

var (
	port   = flag.Int("port", 50051, "The server port")
	wsport = flag.Int("wsport", 8081, "The websocket port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	man.UnimplementedMancalaServiceServer
	waitingRoom    []string
	allUserNames   []string
	userHashes     map[string]string
	gamesInProgress []MancalaGameBoard
	mu             sync.Mutex
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	var gRPCServer server
	gRPCServer.userHashes = make(map[string]string)
	go websocketSetup(&gRPCServer)
	man.RegisterMancalaServiceServer(s, &gRPCServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func fnvHash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32()))
}
