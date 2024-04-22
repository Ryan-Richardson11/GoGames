package sudoku

import (
	"fmt"
	"math/rand"
)

func findValue(value int, arr []int) bool {
	for _, i := range arr {
		if i == value {
			return true
		}
	}
	return false
}

func generateRandomNumber() int {
	randomNumber := rand.Intn(10)
	return randomNumber
}

func DrawCells() {
	cellLines := [3]int{2, 5, 8}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(" ", generateRandomNumber(), " ")
			if findValue(j, cellLines[:]) {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if findValue(i, cellLines[:]) {
			fmt.Print("------------------------------")
		}
		fmt.Println()
	}
}

// Backtracking algorithm to generate random puzzle
// Gives user n number of starting digits depending on difficulty
// func GenerateSudokuPuzzle() {

// }
