package main

import (
	"fmt"
)

var skylineDefinitions = map[rune][]int{
	'n': {4, -1, 3, 3, -1, -1},
	's': {-1, -1, -1, 4, -1, -1},
	'e': {-1, 4, -1, -1, 5, 3},
	'w': {3, -1, -1, 6, -1, 3},
}

var gameBoard [36]int

func solveGameBoard() bool {
	nextEmptyIndex := getNextEmptyIndex()

	if nextEmptyIndex == -1 {
		return true
	}

	if nextEmptyIndex == 3 { // Check if index is the 3rd index
		printBoardState(gameBoard)
	}

	for _, i := range getValidOptions(nextEmptyIndex) {
		gameBoard[nextEmptyIndex] = i
		if verifyGameBoard() && solveGameBoard() {
			return true
		}
		gameBoard[nextEmptyIndex] = 0
	}
	return false
}

func printBoardState(boardCopy [36]int) {
	for i, val := range boardCopy {
		if i%6 == 0 {
			fmt.Println()
		}
		fmt.Print(val, " ")
	}
	fmt.Println("\n---------------------")
}

func getNextEmptyIndex() int {
	for i, val := range gameBoard {
		if val == 0 {
			return i
		}
	}
	return -1
}

func getValidOptions(index int) []int {
	validOptions := []int{1, 2, 3, 4, 5, 6}

	// if index == 0 {
	// 	return []int{3, 4, 5, 6}
	// }

	x := index % 6
	y := index / 6

	for i := 0; i < 6; i++ {
		if gameBoard[x+i*6] != 0 {
			validOptions = removeElement(validOptions, gameBoard[x+i*6])
		}
	}

	for i := 0; i < 6; i++ {
		if gameBoard[i+y*6] != 0 {
			validOptions = removeElement(validOptions, gameBoard[i+y*6])
		}
	}

	return validOptions
}

func removeElement(slice []int, elem int) []int {
	for i := range slice {
		if slice[i] == elem {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func verifyGameBoard() bool {
	return checkRowsUnique() && checkColumnsUnique() && checkRowsSkylines() && checkColumnsSkylines()
}

func checkRowsUnique() bool {
	for y := 0; y < 6; y++ {
		if !checkRowUnique(y) {
			return false
		}
	}
	return true
}

func checkColumnsUnique() bool {
	for x := 0; x < 6; x++ {
		if !checkColumnUnique(x) {
			return false
		}
	}
	return true
}

func checkRowUnique(y int) bool {
	usedNumbers := make(map[int]bool)

	for x := 0; x < 6; x++ {
		cell := gameBoard[x+y*6]
		if cell != 0 {
			if usedNumbers[cell] {
				return false
			}
			usedNumbers[cell] = true
		}
	}
	return true
}
func checkColumnsSkylines() bool {
	for x := 0; x < 6; x++ {
		if skylineDefinitions['n'][x] != -1 {
			currentMax := 0
			visibleCount := 0
			emptyCells := 0

			for y := 0; y < 6; y++ {
				cell := gameBoard[x+y*6]
				if cell > currentMax {
					currentMax = cell
					visibleCount++
				}
				if cell == 0 {
					emptyCells++
				}
			}

			if visibleCount > skylineDefinitions['n'][x] || (emptyCells == 0 && visibleCount < skylineDefinitions['n'][x]) {
				return false
			}
		}

		if skylineDefinitions['s'][x] != -1 {
			currentMax := 0
			visibleCount := 0
			emptyCells := 0

			for y := 5; y >= 0; y-- {
				cell := gameBoard[x+y*6]
				if cell > currentMax {
					currentMax = cell
					visibleCount++
				}
				if cell == 0 {
					emptyCells++
				}
			}

			if visibleCount > skylineDefinitions['s'][x] || (emptyCells == 0 && visibleCount < skylineDefinitions['s'][x]) {
				return false
			}
		}
	}

	return true
}

func checkRowsSkylines() bool {
	for y := 0; y < 6; y++ {
		if skylineDefinitions['e'][y] != -1 {
			currentMax := 0
			visibleCount := 0
			emptyCells := 0

			for x := 5; x >= 0; x-- {
				cell := gameBoard[x+y*6]
				if cell > currentMax {
					currentMax = cell
					visibleCount++
				}
				if cell == 0 {
					emptyCells++
				}
			}

			if visibleCount > skylineDefinitions['e'][y] || (emptyCells == 0 && visibleCount < skylineDefinitions['e'][y]) {
				return false
			}
		}

		if skylineDefinitions['w'][y] != -1 {
			currentMax := 0
			visibleCount := 0
			emptyCells := 0

			for x := 0; x < 6; x++ {
				cell := gameBoard[x+y*6]
				if cell > currentMax {
					currentMax = cell
					visibleCount++
				}
				if cell == 0 {
					emptyCells++
				}
			}

			if visibleCount > skylineDefinitions['w'][y] || (emptyCells == 0 && visibleCount < skylineDefinitions['w'][y]) {
				return false
			}
		}
	}

	return true
}

func checkColumnUnique(x int) bool {
	usedNumbers := make(map[int]bool)

	for y := 0; y < 6; y++ {
		cell := gameBoard[x+y*6]
		if cell != 0 {
			if usedNumbers[cell] {
				return false
			}
			usedNumbers[cell] = true
		}
	}

	return true
}

func main() {
	// this row is easy to calculate
	// prefilling values makes backtracking much faster
	gameBoard[18] = 1
	gameBoard[19] = 2
	gameBoard[20] = 3
	gameBoard[21] = 4
	gameBoard[22] = 5
	gameBoard[23] = 6

	solveGameBoard()

	for i, val := range gameBoard {
		if i%6 == 0 {
			fmt.Println()
		}
		fmt.Print(val, " ")
	}
}
