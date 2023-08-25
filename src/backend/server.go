package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	man "github.com/KeatonBrink/MancalaGame/src/protos"
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
	s.waitingRoom = append(s.waitingRoom, req.GetUserName())
	if len(s.waitingRoom) > 1 {
		return &man.HandshakeResponse{
			ErrorCode:              0,
			ErrorMessage:           "",
			Message:                "Game Starting Soon",
			ServerWebSocketAddress: "",
		}, nil
	} else {
		return &man.HandshakeResponse{
			ErrorCode:              0,
			ErrorMessage:           "",
			Message:                "Waiting for Second Player",
			ServerWebSocketAddress: "8081",
		}, nil
	}
}

func handleWebSocket(serverObj *server, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		serverObj.mu.Lock()
		if len(serverObj.waitingRoom) > 1 {
			// Send a message to the client
			err = conn.WriteMessage(websocket.TextMessage, []byte("Opponent Found"))
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
			serverObj.mu.Unlock()
			newGameBoard := MancalaGameBoard{
				p1Name: serverObj.waitingRoom[0],
				p2Name: serverObj.waitingRoom[1],
			}
			newGameBoard.resetBoard()
			serverObj.gameInProgress = append(serverObj.gameInProgress, newGameBoard)
			if len(serverObj.waitingRoom) > 2 {
				serverObj.waitingRoom = serverObj.waitingRoom[2:]
			} else {
				serverObj.waitingRoom = []string{}
			}

			return
		}
		serverObj.mu.Unlock()
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	var gRPCServer server
	go websocketSetup(&gRPCServer)
	man.RegisterMancalaServiceServer(s, &gRPCServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func websocketSetup(serverObj *server) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(serverObj, w, r)
	})
	log.Printf("Websocket listening at %v", *wsport)
	http.ListenAndServe(":"+strconv.Itoa(*wsport), nil)
}
