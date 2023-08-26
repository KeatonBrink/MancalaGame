package main

import (
	"context"
	"flag"
	"fmt"
	"hash/fnv"
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
	userHashes     map[string]string
	gameInProgress []MancalaGameBoard
	mu             sync.Mutex
}

type MancalaGameBoard struct {
	p1Name string
	p1Row  [7]int
	p2Name string
	p2Row  [7]int
	playerTurn int
}

func (mgb *MancalaGameBoard) String(player int) string {
	if player == 1 {
		return fmt.Sprintf("%v;%v", mgb.p1Row, mgb.p2Row)
	} else if player == 2 {
		return fmt.Sprintf("%v;%v", mgb.p2Row, mgb.p1Row)
	} else {
		return ""
	}
}

func (mgb *MancalaGameBoard) resetBoard() {
	for i := 0; i < 6; i++ {
		mgb.p1Row[i] = 4
		mgb.p2Row[i] = 4
	}
	mgb.p1Row[6] = 0
	mgb.p2Row[6] = 0
	mgb.playerTurn = 1
}

func (s *server) GameHandshake(ctx context.Context, req *man.HandshakeRequest) (*man.HandshakeResponse, error) {
	log.Printf("Received Game Handshake: %v", req.GetUserName())
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.allUserNames {
		if v == req.GetUserName() {
			return &man.HandshakeResponse{
				ErrorCode:              1,
				ErrorMessage:           "Name already in use",
				Message:                "",
				UserHash:               "",
				ServerWebSocketAddress: "",
			}, nil
		}
	}
	s.allUserNames = append(s.allUserNames, req.GetUserName())
	s.waitingRoom = append(s.waitingRoom, req.GetUserName())
	// Hash the username and store it
	newUserHash := fnvHash(req.GetUserName())
	s.userHashes[req.GetUserName()] = newUserHash
	if len(s.waitingRoom) > 1 {
		return &man.HandshakeResponse{
			ErrorCode:              0,
			ErrorMessage:           "",
			Message:                "Game Starting Soon",
			UserHash:               newUserHash,
			ServerWebSocketAddress: "",
		}, nil
	} else {
		return &man.HandshakeResponse{
			ErrorCode:              0,
			ErrorMessage:           "",
			Message:                "Waiting for Second Player",
			UserHash:               newUserHash,
			ServerWebSocketAddress: "8081",
		}, nil
	}
}

func (s *server) MakeMove(ctx context.Context, req *man.MoveRequest) (*man.MoveResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Get the user's name from user hash
	userName := s.userHashes[req.GetUserHash()]
	log.Printf("Received Move Request: %v", userName)
	// Find the game board by cycling through boards in progress
	var gameBoard MancalaGameBoard
	player := 0
	for i, v := range s.gameInProgress {
		if v.p1Name == userName || v.p2Name == userName {
			if v.p1Name == userName {
				player = 1
			} else {
				player = 2
			}
			gameBoard = s.gameInProgress[i]
			break
		}
	}
	// Check if the game board was found
	if gameBoard.p1Name == "" && gameBoard.p2Name == "" {
		return &man.MoveResponse{
			ErrorCode:    1,
			ErrorMessage: "Game not found",
			Message:      "",
			Board:        "",
		}, nil
	}
	// Check if the move is valid
	// Convert int32 to int
	move := int(req.GetPitIndex())
	returnBoard, err := MakeMove(gameBoard, userName, move)
	if err != nil {
		return &man.MoveResponse{
			ErrorCode:    3,
			ErrorMessage: err.Error(),
			Message:      "",
			Board:        "",
		}, nil
	}
	// Check if the game is over
	if isGameOver(returnBoard) {
		// Send the game over message
		return &man.MoveResponse{
			ErrorCode:    0,
			ErrorMessage: "",
			Message:      "Game Over",
			Board:        returnBoard.String(player),
		}, nil
	}
	// Send the updated board
	return &man.MoveResponse{
		ErrorCode:    0,
		ErrorMessage: "",
		Message:      "New Board",
		Board:        returnBoard.String(player),
	}, nil
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

func fnvHash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.Itoa(int(h.Sum32()))
}
