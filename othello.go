///Author: Alexander Jose
///SID: 10388912
///10/29/2024
///Assignment 3: Othello.
///A program that implements the game of othello, playable by two players, or one player, via the minimax algorithm
///

package main

import (
	"fmt"
	"strconv"
)

// A way to represent the player without using magic numbers. Can be used to keep track of the current player.
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

// Settings
var debugMode bool = false
var multiplayer bool = true

//Todo: a way to keep track of which players are human and which are AI

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
	displayCharacters := []rune{'□', '○', '●'}

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

func checkVictory(state boardstate) {
	//Todo: test

}

// Get the resulting state from trying to place a piece at this position.
// If the piece cannot be placed, isValid will be false, and resultingState has no guarantees of usefulness.
func getStateFromMove(currentMove Move, initialState boardstate) (newState boardstate, isValid bool) {
	//In golang, this will create a copy of the array rather than a reference
	newState = initialState

	//We will set this flag if any of the directions gives us valid changes
	isValid = false

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
		// This stores the location of each tile iterated over before we want to change its color, in the format {{x,y},{x,y}}
		// The [:0] makes the length 0, so that we can check if changes is empty.
		changes := make([][]byte, 6)[:0]

		//Make an initial step in the correct direction. We make our initial step outside the loop so it can check if the location is inside the board, and immediately end if it is not.
		currentRow += byte(rowStep)
		currentColumn += byte(columnStep)

		// We loop until the current location values exceed the size of the array, which is the worst case. If they do, we know the direction does not have flippable tiles.
		// This loop should usually terminate early, however, either by encountering a blank spot, or a tile of the same color as the color of the tile placed.
		for (currentRow >= 0 && currentRow < byte(len(initialState))) && (currentColumn >= 0 && currentColumn < byte(len(initialState[0]))) {

			//Deal with the tile we hit
			if newState[currentRow][currentColumn] == BLANK {
				//We have hit a blank space, which means this is not a valid direction
				break
			}
			if newState[currentRow][currentColumn] == currentMove.color {
				//This ends our direction

				//If the direction made no changes, this direction is a failure
				if len(changes) == 0 {
					break
				}
				//Otherwise, apply the changes to the new board, and mark the move as valid
				for _, changeLocation := range changes {
					//Set each tile to the color of the move
					newState[changeLocation[0]][changeLocation[1]] = currentMove.color
				}
				isValid = true
			}

			//In this case, we can assume we have hit the opposite color
			//We should add the current location to the list of locations that could possibly be changed, and continue
			changes = append(changes, []byte{currentRow, currentColumn})

			//Make a step in the correct direction
			currentRow += byte(rowStep)
			currentColumn += byte(columnStep)

		}

	}
	//As we have handled the changes made by placing a piece with this move, we must now place the piece
	newState[currentMove.row][currentMove.column] = currentMove.color

	return newState, isValid

}

// Todo: Function that computes possible moves, and their corresponding board states.
func getPossibleMoves(state boardstate, player Color) ([]Move, []boardstate) {

	// Create two slices, so that we can dynamically add possible moves and their resulting states
	// We initialize with a size of 60, as there are only 60 empty spots on an othello board.
	// We make these into 0 length slices, so that we append into allocated space, rather than past it.
	// While the arrays can be smaller, due to the fact that not all empty spaces are usable, we do not know the precise maximum amount of moves possible on one turn
	possibleMoves := make([]Move, 60)[:0]
	resultingStates := make([]boardstate, 60)[:0]

	//We loop over each tile, checking if it is a possible move.
	for rowIndex, row := range state {
		for columnIndex, tileColor := range row {
			if tileColor != BLANK {
				//Since the tile is not empty, the piece cannot be placed here
				continue
			}
			currentMove := Move{byte(rowIndex), byte(columnIndex), player}

			// Check if we can place a piece at this position, and how that would change the board
			resultingState, isValid := getStateFromMove(currentMove, state)
			if isValid {
				possibleMoves = append(possibleMoves, currentMove)
				resultingStates = append(resultingStates, resultingState)
				//TOdo: unused debug, remove
				//fmt.Println("Adding to possible moves:")
				//displayBoardState(resultingState)
			}

		}
	}

	return possibleMoves, resultingStates

}

func handleTurn(board boardstate, color Color, isAI bool) (resultingBoard boardstate) {

	var input string
	var colorName = "BLACK"
	if color == WHITE {
		colorName = "WHITE"
	}
	possibleMoves, resultingStates := getPossibleMoves(board, color)

	//We label the outer loop so that once we make a move, we can end the turn
endTurn:
	for {
		//Display instructions
		if isAI {
			fmt.Println("Type enter to allow the CPU opponent to make a move. Enter 1 to toggle debug mode. Enter 2 to toggle AI. \n%s AI:", colorName)
		} else {
			fmt.Printf("Enter a valid tile, in the format (1A) to make a move. Enter 1 to toggle debug mode. Enter 2 to toggle AI.\n%s:", colorName)
		}
		if len(possibleMoves) == 0 {
			//User can only forfeit
			fmt.Println("No moves possible. Press enter to forfeit turn.")

		}
		//Wait for user input
		fmt.Scanln(&input)
		//Handle settings used in both modes.
		switch input {
		case "1":
			//Invert the current value, print the result to the user, and start from the top.
			debugMode = !debugMode
			fmt.Printf("Debug mode: %t\n", debugMode)
			continue
		case "2":
			//Todo: AI setup. Should be one for each color.
			debugMode = !debugMode
			fmt.Printf("Debug mode: %t\n", debugMode)
			continue
		default:
			break
		}
		if len(possibleMoves) == 0 {
			//User forfeits turn
			break endTurn
		}

		if isAI {
			//Todo: AI player
			break
		} else {
			//Human player

			//Handle any moves

			for moveIndex, move := range possibleMoves {
				/// This string is what the user should input if they want to make the move we are currently testing
				var moveAsString string = strconv.Itoa(int(move.row)) + string(65+move.column)

				if input == moveAsString {
					// We have found the move the player made, so we set the state of the board, and exit the loop.
					resultingBoard = resultingStates[moveIndex]
					if debugMode {
						//Todo: pretty-print, prevent color from being printed
						fmt.Printf("Possible moves: %v\n", possibleMoves[0:moveIndex])
					}
					break endTurn
				}
			}
			//Attempted move is not in the list of moves
			fmt.Println("Invalid move!")

		}

	}

	return resultingBoard

}

func main() {
	//Todo: When initializing, print rules
	fmt.Println("")

	board := initialState()
	currentPlayer := BLACK
	for {
		displayBoardState(board)

		board = handleTurn(board, currentPlayer, false)

		// Switch player
		if currentPlayer == BLACK {
			currentPlayer = WHITE
		} else {
			currentPlayer = BLACK
		}

		checkVictory(board)

	}
	//"You are %d\nEnter a valid tile, in the format (1A) to make a move. Enter 1 to enable debug mode. Enter 2 to enable AI."

}
