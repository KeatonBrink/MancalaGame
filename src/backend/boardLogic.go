// Contains the mancala board logic.
package main

import "fmt"

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
	
	// Try to correctly get the pits pointers
	var curPlayerPits, otherPlayerPits *[7]int
	if board.p1Name == player {
		curPlayerPits = &board.p1Row
		otherPlayerPits = &board.p2Row
	} else if board.p2Name == player {
		curPlayerPits = &board.p2Row
		otherPlayerPits = &board.p1Row
	} else {
		return board, fmt.Errorf("invalid player")
	}
	// Move the stones
	stones := curPlayerPits[move]
	curPlayerPits[move] = 0
	// Simple way to check if the last pit was 0 after adding a stone
	finalPit := -1
	finalInd := -1
	for i := 1; i <= stones; i++ {
		pitIndex := (move + i) % 13
		if pitIndex < 7 {
			curPlayerPits[pitIndex] += 1
			finalPit = curPlayerPits[pitIndex]
			finalInd = pitIndex
		} else {
			pitIndex = pitIndex - 7
			otherPlayerPits[pitIndex] += 1
			finalPit = -1
			finalInd = -1
		}
	}

	if finalPit == 1 && finalInd != -1 && finalInd != 6 {
		crossPit := crossPitConversion(finalInd)
		curPlayerPits[6] += otherPlayerPits[crossPit]
		otherPlayerPits[crossPit] = 0
	}
	return board, nil
}

func crossPitConversion(index int) int {
	return 6 - index
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