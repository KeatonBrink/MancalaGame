// Contains the mancala board logic.
package main

import "fmt"

// MakeMove takes in a board, a player, and a move and returns the new board
// after the move has been made.

func MakeMove(board MancalaGameBoard, player string, move int) (MancalaGameBoard, error) {
	// Check if the move is valid
	if board.p1Name == player {
		if board.p1Row[move] == 0 {
			return board, fmt.Errorf("invalid move")
		}	
	} else if board.p2Name == player {
		if board.p2Row[move] == 0 {
			return board, fmt.Errorf("invalid move")
		}
	} else {
		return board, fmt.Errorf("invalid player")
	}
	
	// This logic is pretty bad, but I want to see how bad before I fix it
	// Make the move
	if board.p1Name == player {
		// Move the stones
		stones := board.p1Row[move]
		board.p1Row[move] = 0
		for i := 0; i < stones; i++ {
			if move + i + 1 > 6 {
				board.p2Row[(move + i + 1) % 7] += 1
			} else {
				board.p1Row[move + i + 1] += 1
			}
		}
		// Check if the last stone was in an empty pit
		// I think this is a bug, but I'm not sure
		// The move could have ended in the other 
		if board.p1Row[(move + stones) % 7] == 1 {
			board.p1Row[6] += board.p2Row[6 - ((move + stones) % 7)] + 1
			board.p2Row[6 - ((move + stones) % 7)] = 0
			board.p1Row[(move + stones) % 7] = 0
		}
	} else if board.p2Name == player {
		// Move the stones
		stones := board.p2Row[move]
		board.p2Row[move] = 0
		for i := 0; i < stones; i++ {
			if move + i + 1 > 6 {
				board.p1Row[(move + i + 1) % 7] += 1
			} else {
				board.p2Row[move + i + 1] += 1
			}
		}
		// Check if the last stone was in an empty pit
		if board.p2Row[(move + stones) % 7] == 1 {
			board.p2Row[6] += board.p1Row[6 - ((move + stones) % 7)] + 1
			board.p1Row[6 - ((move + stones) % 7)] = 0
			board.p2Row[(move + stones) % 7] = 0
		}
	} else {
		return board, fmt.Errorf("invalid player")
	}
	return board, nil
}

func isGameOver(board MancalaGameBoard) bool {
	// Check if the game is over
	p1Empty := true
	p2Empty := true
	for i := 0; i < 6; i++ {
		if board.p1Row[i] != 0 {
			p1Empty = false
		}
		if board.p2Row[i] != 0 {
			p2Empty = false
		}
	}
	if p1Empty || p2Empty {
		return true
	}
	return false
}