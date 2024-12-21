package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ReadFile reads the content of a text file line by line
func ReadFile(filename string) ([]int, []int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a slice to store lists
	var listA []int
	var listB []int

	// Read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Remove leading/trailing whitespace
		if line == "" {
			continue // Skip empty lines
		}

		// Split the line into two parts
		parts := strings.Fields(line) // Fields splits by whitespace
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Invalid line format: %s\n", line)
			continue
		}

		// Convert each part to an integer
		first, err1 := strconv.Atoi(parts[0])
		second, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Fprintf(os.Stderr, "Error parsing numbers in line: %s\n", line)
			continue
		}

		// Append the numbers to their respective slices
		listA = append(listA, first)
		listB = append(listB, second)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}
	sort.Ints(listA)
	sort.Ints(listB)
	return listA, listB, nil
}

// SolvePuzzle takes the parsed input and solves the puzzle
func SolvePuzzle(listA []int, listB []int) (int, error) {

	// Int to hold the sum of the distances
	result := 0

	// distance = abs(listA[i] - listB[i])
	for index := range listA {

		tempResult := listA[index] - listB[index]
		if tempResult < 0 {
			tempResult = -tempResult
		}
		result += tempResult
	}

	return result, nil

}

func main() {
	// Check if a filename is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}

	// Get the filename from the command-line arguments
	filename := os.Args[1]

	// Read the file
	listA, listB, err := ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	result, _ := SolvePuzzle(listA, listB)
	fmt.Println("Puzzle Solution", result)
}
