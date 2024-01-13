// Contains the mancala board logic.
package main

import "fmt"

type MancalaGameBoard struct {
	p1Name     string
	p1Row      [7]int
	p2Name     string
	p2Row      [7]int
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

func (board *MancalaGameBoard) MakeMove(player string, move int) error {
	if board.playerTurn == -1 {
		return fmt.Errorf("game over, invalid move")
	}

	// Try to correctly get the pits pointers
	var curPlayerPits, otherPlayerPits *[7]int
	if board.playerTurn == 1 {
		curPlayerPits = &board.p1Row
		otherPlayerPits = &board.p2Row
	} else if board.playerTurn == 2 {
		curPlayerPits = &board.p2Row
		otherPlayerPits = &board.p1Row
	} else {
		return fmt.Errorf("invalid player")
	}

	// Check if the move is valid
	if curPlayerPits[move] == 0 {
		return fmt.Errorf("invalid move")
	}

	// Move the stones
	stones := curPlayerPits[move]
	curPlayerPits[move] = 0
	// Simple way to check if the last pit was 0 after adding a stone
	finalPitStones := -1
	finalInd := -1
	for i := 1; i <= stones; i++ {
		pitIndex := (move + i) % 13

		if pitIndex < 7 {
			curPlayerPits[pitIndex] += 1
			finalPitStones = curPlayerPits[pitIndex]
			finalInd = pitIndex
		} else {
			pitIndex = pitIndex - 7
			otherPlayerPits[pitIndex] += 1
			finalPitStones = -1
			finalInd = -1
		}
	}

	crossPit := crossPitConversion(finalInd)
	if finalPitStones == 1 && finalInd != -1 && finalInd != 6 && otherPlayerPits[crossPit] > 0 {
		curPlayerPits[6] += otherPlayerPits[crossPit] + 1
		otherPlayerPits[crossPit] = 0
		curPlayerPits[finalInd] = 0
	}

	if finalInd != 6 {
		board.playerTurn = (board.playerTurn % 2) + 1
	}

	if board.isGameOver() {
		for i := 0; i < 6; i++ {
			curPlayerPits[6] += curPlayerPits[i]
			curPlayerPits[i] = 0
			otherPlayerPits[6] += otherPlayerPits[i]
			otherPlayerPits[i] = 0
		}
		board.playerTurn = -1
	}

	return nil
}

func crossPitConversion(index int) int {
	return 6 - index
}

func (board *MancalaGameBoard) isGameOver() bool {
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

func (board *MancalaGameBoard) GetWinner() (int, error) {
	if !board.isGameOver() {
		return -1, fmt.Errorf("game not over")
	}
	if board.p1Row[6] > board.p2Row[6] {
		return 1, nil
	}
	if board.p1Row[6] < board.p2Row[6] {
		return 2, nil
	}
	return 0, nil
}
