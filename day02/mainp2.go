package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads the content of a text file line by line
func ReadFile(filename string) ([][]int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a slice of slices to store reports
	var reports [][]int

	// Read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var report []int

		line := strings.TrimSpace(scanner.Text()) // Remove leading/trailing whitespace
		if line == "" {
			continue // Skip empty lines
		}

		// Split the line into two parts
		parts := strings.Fields(line) // Fields splits by whitespace

		// Convert each part to an integer and store in report
		for _, value := range parts {
			num, err := strconv.Atoi(value)

			if err != nil {
				fmt.Println("Error converting:", err)
				return nil, err
			}

			report = append(report, num)
		}

		reports = append(reports, report)

	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return reports, nil
}

func CheckSafety(report []int) bool {

	var lastDirection int = 0
	var direction int = 0

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		absDiff := int(math.Abs(float64(diff)))
		if !(absDiff >= 1 && absDiff <= 3) { // outside safe range
			return false
		}

		if diff > 0 {
			direction = 1
		} else {
			direction = -1
		}

		if lastDirection != 0 && direction != lastDirection {
			return false // unsafe
		}

		lastDirection = direction
	}
	return true
}

func CheckDampenedSafety(report []int) bool {
	if CheckSafety(report) {
		return true
	}

	for i, _ := range report {

		dampenedReport := append([]int{}, report[:i]...)
		dampenedReport = append(dampenedReport, report[i+1:]...)

		if CheckSafety(dampenedReport) {
			return true
		}
	}
	return false
}

// SolvePuzzle takes the parsed input and solves the puzzle
func SolvePuzzle1(reports [][]int) (int, error) {

	// Int to hold the number of safe reports
	safeCount := 0

	// for each report in reports,
	for _, report := range reports {

		if len(report) < 2 {
			return 0, fmt.Errorf("Not enough Elements to determine safety")
		}

		if CheckSafety(report) {
			safeCount++
		}
	}
	return safeCount, nil

}

// SolvePuzzle takes the parsed input and solves the puzzle
func SolvePuzzle2(reports [][]int) (int, error) {

	// Int to hold the number of safe reports
	safeCount := 0

	// for each report in reports,
	for _, report := range reports {

		if len(report) < 2 {
			return 0, fmt.Errorf("Not enough Elements to determine safety")
		}

		if CheckDampenedSafety(report) {
			safeCount++
		}
	}
	return safeCount, nil

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
	reports, err := ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	result, _ := SolvePuzzle1(reports)
	fmt.Println("Puzzle one Solution", result)

	result2, _ := SolvePuzzle2(reports)
	fmt.Println("Puzzle two Solution", result2)
}
