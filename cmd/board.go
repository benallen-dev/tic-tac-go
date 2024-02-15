package main

import (
	"math"
	"errors"
)

var comboToSquares = map[int][3]int{
	7: [3]int{0, 1, 2},
	56: [3]int{3, 4, 5},
	73: [3]int{0, 3, 6},
	84: [3]int{2, 4, 6},
	146: [3]int{1, 4, 7},
	273: [3]int{0, 4, 8},
	292: [3]int{2, 5, 8},
	448: [3]int{6, 7, 8},
}

type Square string

const (
	Empty Square = " "
	X     Square = "X"
	O     Square = "O"
)

type Board struct {
	squares [9]Square
}

func NewBoard() Board {
	return Board{[9]Square{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}
}

func (b *Board) SetSquare(index int, value Square) error{
	// Check if the square is already taken
	if b.GetSquare(index) != Empty {
		return errors.New("Square is already taken")
	}

	b.squares[index] = value

	return nil
}

func (b *Board) GetSquare(index int) Square {
	return b.squares[index]
}

func (b *Board) GetSquares() [9]Square {
	return b.squares
}

func (b *Board) String() string {

	var padding string = "  "

	output := "\n" + padding
	for i, square := range b.GetSquares() {
		output += " " + string(square) + " "
		if i % 3 == 2 {
			if i != 8 { // Don't add a line after the last row
				output += "\n" + padding + "───┼───┼───\n" + padding
			}
		} else {
			output += "│"
		}
	}

	return output
}

// getSquaresForValue returns the indices of the squares that have the given value
func (b *Board) getSquaresForValue(value Square) []int {
	output := []int{}
	for i, square := range b.GetSquares() {
		if square == value {
			output = append(output, i)
		}
	}

	return output
}

func exponentiate(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

func exponentiateSquares(squares []int) int {
	output := 0
	for _, square := range squares {
		output += exponentiate(2, square)
	}

	return output
}

func (b *Board) HasWinner() ( gameover bool, winner Square) {
	// A winner exists if there are three in a row
	// square numbers are exponentiated to make the combinations unique
	// 0, 1, 2 => 1 + 2 + 4 = 7
	// 0, 3, 6 => 1 + 8 + 64 = 73
	// 0, 4, 8 => 1 + 16 + 256 = 273
	// 1, 4, 7 => 2 + 16 + 128 = 146
	// 2, 5, 8 => 4 + 32 + 256 = 292
	// 2, 4, 6 => 4 + 16 + 64 = 84
	// 3, 4, 5 => 8 + 16 + 32 = 56
	// 6, 7, 8 => 64 + 128 + 256 = 448
	winningCombinations := [8]int{7, 56, 73, 84, 146, 273, 292, 448}

	xSquares := b.getSquaresForValue(X)
	oSquares := b.getSquaresForValue(O)

	// Check for a winner
	xExponentiated := exponentiateSquares(xSquares)	
	oExponentiated := exponentiateSquares(oSquares)

	for _, combination := range winningCombinations {
		if xExponentiated & combination == combination {
			return true, X
		}

		if oExponentiated & combination == combination {
			return true, O
		}
	}

	// Now we have to check if all 9 squares are taken, if so, it's a draw
	if len(xSquares) + len(oSquares) == 9 {
		return true, Empty
	}

	// No winner and no draw
	return false, Empty
}
