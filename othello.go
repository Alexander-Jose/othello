package main

import "fmt"

//A way to represent the player without using magic numbers. Can be used to keep track of the current player.
type Color = byte

const (
	BLANK Color = 0
	BLACK Color = 1
	WHITE Color = 2
)

// Represents a possible configuration of Othello pieces
// 0 represents a blank space
// 1 represents a black piece
// 2 represents a white piece
type boardstate = [8][8]Color

// Stores a move, which contains the location that the player attempts to place a token
type Move struct {
	row    byte
	column byte
	color  Color
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

// Get the resulting state from trying to place a piece at this position.
// If the piece cannot be placed, isValid will be false, and resultingState has no guarantees of usefulness.
func getStateFromMove(currentMove Move, initialState boardstate) boardstate {
	//In golang, this will create a copy of the array rather than a reference
	newState := initialState

	//We will set this flag if any of the directions gives us valid changes
	isValid := false

	// We need to check the tiles in every direction for possible flips.
	// We do this by picking a direction, iterating over it, and making certain that we encounter only enemy pieces until we encounter one of our own.
	// This section concerns picking a direction.
	// We use rowStep and columnStep to be the amount we step in each direction, until we hit a final tile.
	//Todo: step backwards, flipping tiles?

	for direction := 0; direction < 9; direction++ {

		// The rowStep needs to alternate between -1, 0, and 1. If we take the modulo, we get a pattern of 0,1,2,0,1,2,0,1,2. Subtracting 1 allows us to center at 0, the row position of the move
		rowStep := -1 + direction%3
		// The columnStep needs to go between -1, 0, and 1, but we need to keep it at each value for 3 loops, so that we can get each row at each column value.
		columnStep := -1 + direction/3

		//Since the value at the move location + 0,0 is always empty, we do not need a special case, as we will immediately exit when the direction does not move.
		currentRow := currentMove.row
		currentColumn := currentMove.column

		//Initialize a slice which stores all the changes we need to make. We use a size of 6, as this is the maximum number of tiles in one direction that can be modified on an 8 by 8 board
		changes := make([][]byte, 6)

		// We loop until the current location values exceed the size of the array, which is the worst case. If they do, we know the direction does not have flippable tiles.
		// This loop should usually terminate early, however, either by encountering a blank spot, or a tile of the same color as the color of the tile placed.
		for (currentRow >= 0 && currentRow < byte(len(initialState))) && (currentColumn >= 0 && currentColumn < byte(len(initialState[0]))) {
			if newState[currentRow][currentColumn] == 0 {
				//We have hit a blank space, which means this is not a valid direction
				continue
			}
			if Color(newState[currentRow][currentColumn]) == currentMove.color {

			}

		}

		//Swap all of the changed tiles to the correct color. If there are none, such as when there is no final piece of a color, this loop will be skipped.

	}

}

//Todo: Function that computes possible moves, and their corresponding board states.
func getPossibleMoves(state boardstate, player Color) ([]Move, []boardstate) {

	// Create two slices, so that we can dynamically add possible moves and their resulting states
	// 60 is chosen, as there are only 60 empty spots on an othello board.
	// While the arrays can be smaller, due to the fact that not all empty spaces are usable, we do not know the precise maximum amount of moves possible on one turn
	possibleMoves := make([]Move, 60)
	resultingStates := make([]boardstate, 60)

	//We loop over each tile, checking if it is a possible move.
	for rowIndex, row := range state {
		for columnIndex, tileColor := range row {
			if tileColor != BLANK {
				//Since the tile is not empty, the piece cannot be placed here
				continue
			}
			currentMove := Move{byte(rowIndex), byte(columnIndex), player}

			// Check if we can place a piece at this position, and how that would change the board
			resultingState, isValid := getStateFromMove(currentMove)
			if isValid {
				possibleMoves = append(possibleMoves, currentMove)
				resultingStates = append(resultingStates, resultingState)
			}

		}
	}

	//Return slices of the arrays, so that the whole

}

func main() {
	//Todo: When initializing, print rules
	fmt.Println("")

	board := initialState()
	displayBoardState(board)

}
