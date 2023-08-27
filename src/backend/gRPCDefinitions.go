package main

import (
	"context"
	"log"

	man "github.com/KeatonBrink/MancalaGame/src/protos/generated"
)

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
	s.userHashes[newUserHash] = req.GetUserName()
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
	log.Printf("Received Move Request: %v %v %v", userName, req.GetUserHash(), req.GetPitIndex())
	// Find the game board by cycling through boards in progress
	var gameBoard MancalaGameBoard
	player := 0
	gameInProgressIndex := -1
	for i, v := range s.gamesInProgress {
		if v.p1Name == userName || v.p2Name == userName {
			if v.p1Name == userName {
				player = 1
			} else {
				player = 2
			}
			gameBoard = s.gamesInProgress[i]
			gameInProgressIndex = i
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
	// Check if it is the player's turn
	if player != gameBoard.playerTurn {
		return &man.MoveResponse{
			ErrorCode:    2,
			ErrorMessage: "Not your turn",
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
	// Change player turn
	if returnBoard.playerTurn == 1 {
		returnBoard.playerTurn = 2
	} else {
		returnBoard.playerTurn = 1
	}
	s.gamesInProgress[gameInProgressIndex] = returnBoard
	log.Printf("Valid Move, Sending New Board: %v", returnBoard.String(player))
	// Send the updated board
	return &man.MoveResponse{
		ErrorCode:    0,
		ErrorMessage: "",
		Message:      "New Board",
		Board:        returnBoard.String(player),
	}, nil
}