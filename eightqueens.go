package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	for i := range board {
		board[i] = -1
	}

	isSafe := func(col, row int) bool {
		for prevCol := 0; prevCol < col; prevCol++ {
			prevRow := board[prevCol]
			if prevRow == row || prevRow-row == col-prevCol || row-prevRow == col-prevCol {
				return false
			}
		}
		return true
	}

	var solve func(int)
	solve = func(col int) {
		if col == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune('1' + board[i]))
			}
			z01.PrintRune('\n')
			return
		}

		for row := 0; row < 8; row++ {
			if isSafe(col, row) {
				board[col] = row
				solve(col + 1)
				board[col] = -1
			}
		}
	}

	solve(0)
}
