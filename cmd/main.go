package main

import (
	"fmt"

	"github.com/benallen-dev/tic-tac-go/pkg/color"
)

func draw(board Board, currentPlayer Square, cursor int) {

	// Clear the screen
	fmt.Print("\033[H\033[2J")
	// fmt.Print(board.String())

	var padding string = "  "

	output := "\n" + padding
	for i, square := range board.GetSquares() {
		// Left padding
		output += " "

		if square == Empty && i == cursor {
			output += color.LightGray + string(currentPlayer) + color.Reset
		} else if i == cursor {
			output += color.Red + string(square) + color.Reset
		} else {
			output += string(square)
		}

		// Right padding
		output += " "

		// Draw frame
		if i%3 == 2 {
			if i != 8 { // Don't add a line after the last row
				output += "\n" + padding + "───┼───┼───\n" + padding
			}
		} else {
			output += "│"
		}
	}

	fmt.Print(output)

	fmt.Print("\n\n")
}

func main() {
	// Initialize board
	board := NewBoard()
	currentPlayer := X
	cursor := 4

	gameover := false
	winner := Empty

	for gameover == false {

		// Draw board
		draw(board, currentPlayer, cursor)

		// Move cursor down two lines
		fmt.Printf("%s's turn: ", currentPlayer)

		// Get input
		var input int
		foo, err := fmt.Scanln(&input)
		if err != nil || foo != 1 {
			fmt.Println("Error:", err)
		}

		// This is where the logic diverges from the original
		// if the input is hjkl, move the cursor
		// if the input is <CR>, make the move

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

		// Ok now we're done, we can switch players

		// Switch player
		if currentPlayer == X {
			currentPlayer = O
		} else {
			currentPlayer = X
		}

		// Check for a winner
		gameover, winner = board.HasWinner()

	}

	// Draw board one more time because draw() is performed at the top of the loop
	// Cursor -1 is a special case that tells draw() to not draw the cursor
	draw(board, currentPlayer, -1)

	fmt.Printf("\n%s wins!\n", winner)

}
