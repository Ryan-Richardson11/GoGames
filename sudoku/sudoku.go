package sudoku

import "fmt"

func DrawCells() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(" ", 0, " ")
		}
		fmt.Println()
	}
}

// Backtracking algorithm to generate random puzzle
// Gives user n number of starting digits depending on difficulty
// func GenerateSudokuPuzzle() {

// }
