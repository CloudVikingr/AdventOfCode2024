package main

import (
	"fmt"
	"log"
	"os"

	"adventofcode2024/reader"
)

func data(file string) [][]rune {
	lines, err := reader.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		log.Fatal(err)
	}
	var wsPuzzle [][]rune
	for _, line := range lines {
		runeline := []rune(line)
		wsPuzzle = append(wsPuzzle, runeline)
	}
	return wsPuzzle
}

func Puzzle1(data [][]rune) int {
	// Implement the logic for Puzzle1 here
	word := "XMAS"
	return SearchWord(data, word)
}
func Puzzle2(data [][]rune) int {
	return SearchXMas(data)
}

func SearchXMas(wsPuzzle [][]rune) int {
	rows := len(wsPuzzle)
	cols := len(wsPuzzle[0])
	crossCount := 0

	word := "MAS"

	checkCross := func(row, col int) bool {
		// Check if the cross is within bounds and return if its not
		if row-1 < 0 || row+1 >= rows || col-1 < 0 || col+1 >= cols {
			return false
		}

		positions := [4][2]int{
			{row - 1, col - 1}, // top left 0
			{row - 1, col + 1}, // top right 1
			{row + 1, col - 1}, // bottom left 2
			{row + 1, col + 1}, // bottom right 3
		}

		// M * S   M * M   S * M   S * S
		// * A *   * A *   * A *   * A *
		// M * S   S * S   S * M   M * M

		patterns := [4][4]int{
			{0, 3, 2, 1},
			{0, 3, 1, 2},
			{1, 2, 3, 0},
			{3, 0, 2, 1},
		}

		for _, pattern := range patterns {
			if wsPuzzle[positions[pattern[0]][0]][positions[pattern[0]][1]] == rune(word[0]) &&
				wsPuzzle[positions[pattern[1]][0]][positions[pattern[1]][1]] == rune(word[2]) &&
				wsPuzzle[positions[pattern[2]][0]][positions[pattern[2]][1]] == rune(word[0]) &&
				wsPuzzle[positions[pattern[3]][0]][positions[pattern[3]][1]] == rune(word[2]) {
				return true
			}

		}
		return false
	}

	// traverse grid and search for the cross centering on A
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if wsPuzzle[row][col] == 'A' {
				if checkCross(row, col) {
					crossCount++
				}
			}
		}
	}
	return crossCount

}

func SearchWord(wsPuzzle [][]rune, word string) int {
	rows := len(wsPuzzle)
	cols := len(wsPuzzle[0])
	wordLen := len(word)
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	matchCount := 0

	// anonymous function to search the grid in a direction
	searchDirection := func(row, col, dRow, dCol int) bool {
		for i := 0; i < wordLen; i++ {
			// i = 0 nrow = 0 + 0 * 0  nCol = 3 + 0 * 11    (0,3)
			// i = 1 nrow = 0 + 1 * 0  nCol = 3 + 1 * 1  (0,4)
			nRow, nCol := row+i*dRow, col+i*dCol
			// check if the next letter is out of bounds or not in the word
			if nRow < 0 || nRow >= rows || nCol < 0 || nCol >= cols || wsPuzzle[nRow][nCol] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	// search for the word in all directions
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, direction := range directions {
				if searchDirection(row, col, direction[0], direction[1]) {
					//return row, col, direction[0], direction[1], true
					matchCount++
				}
			}
		}
	}
	return matchCount
}

func main() {

	// Check if a filename is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	// Get the filename from the command-line arguments
	filename := os.Args[1]

	fmt.Println("Hello, Advent of Code 2024!")
	result1 := Puzzle1(data(filename))
	fmt.Println("Puzzle 1:", result1)

	result2 := Puzzle2(data(filename))
	fmt.Println("Puzzle 1:", result2)
}
