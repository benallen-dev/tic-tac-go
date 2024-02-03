package main

import (
	"fmt"
)


func draw(board Board) {
	fmt.Print("\033[H\033[2J")
	fmt.Print(board.String())
	fmt.Print("\n\n")
}

func main() {

	// Initialize board
	board := NewBoard()
	currentPlayer := X


	// While there is no winner
	for board.HasWinner() == Empty { 
	
		// Draw board
		draw(board)

		// Move cursor down two lines
		fmt.Printf("%s's turn: ", currentPlayer)

		// Get input
		var input int
		foo, err := fmt.Scanln(&input)
		if err != nil || foo != 1 {
			fmt.Println("Error:", err)
		}

		// Check if input is valid
		if input < 0 || input > 8 {
			// Just ignore it
			continue
		}

		// Attempt to make move
		moveErr := board.SetSquare(input, currentPlayer) 
		if moveErr != nil {
			// Not a valid move either!
			continue
		}

		// Switch player
		if currentPlayer == X {
			currentPlayer = O
		} else {
			currentPlayer = X
		}

	}

	// Draw board one more time because draw() is performed at the top of the loop
	draw(board)

	var winner Square
	if currentPlayer == X {
		winner = O
	} else {
		winner = X
	}

	// We have a winner
	fmt.Printf("\n%s wins!\n", winner)

}
