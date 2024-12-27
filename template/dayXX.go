package main

import (
	"fmt"
	"log"
	"os"

	"adventofcode2024/reader"
)

func data(file string) []string {
	lines, err := reader.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		log.Fatal(err)
	}
	return lines
}

func Puzzle1(data) int {
	result := 0
	return result
}
func Puzzle2(data) int {
	result := 0
	return result
}

func main() {
	// Check if a filename is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	// Get the filename from the command-line arguments
	filename := os.Args[1]

	// read input
	input := data(filename)

	fmt.Println("Hello, Advent of Code 2024!")

	// Solve puzzle 1
	result1 := Puzzle1(input)
	fmt.Println("Puzzle 1:", result1)

	// Solve puzzle 2
	result2 := Puzzle2(data(filename))
	fmt.Println("Puzzle 2:", result2)
}
