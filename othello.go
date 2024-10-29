package main

import "fmt"

// Represents a possible configuration of Othello pieces
// 0 represents a blank space
// 1 represents a black piece
// 2 represents a white piece
//Todo: following states uncertain
// 3 represents a possible black move
// 4 represents a possible white move
type boardstate = [8][8]byte

//A way to represent the player without using magic numbers. Can be used to keep track of the current player.
type Player = byte

const (
	WHITE Player = 0
	BLACK Player = 1
)

// Stores a move, which contains the location that the player attempts to place a token
type Move struct {
	row    byte
	column byte
}

// Returns an initial starting state for the board
func initialState() boardstate {
	var blankBoard boardstate

	//Initializes the center four pieces in a checkerboard pattern
	blankBoard[4][4] = 1
	blankBoard[3][3] = 1
	blankBoard[3][4] = 2
	blankBoard[4][3] = 2

	return blankBoard

}

func displayBoardState(state boardstate) {

	// Values used to display pieces on a board. The characters correpond to the indexes used in the board state
	displayCharacters := []rune{'□', '●', '○', '◌', '◌'}

	//Print column labels
	fmt.Println("  A B C D E F G H")

	//Todo: Add row labels
	for rowIndex, row := range state {
		//Print the row labels
		fmt.Printf("%d ", rowIndex)
		for columnIndex, itemValue := range row {
			fmt.Printf("%s ", string(displayCharacters[itemValue]))
			// Todo: remove. If unused, ignore values
			_ = columnIndex
		}
		fmt.Println()
	}

}

//Todo: Function that computes possible moves, and their corresponding board states.
func getPossibleMoves(state boardstate, player Player) {

}

func main() {
	//Todo: When initializing, print rules
	fmt.Println("")

	board := initialState()
	displayBoardState(board)

}
