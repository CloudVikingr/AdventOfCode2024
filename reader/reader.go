package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadFile reads the content of a text file line by line
func ReadFile(filepath string) ([]string, error) {

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Declare a slice to store the lines
	var lines []string

	// Read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Remove leading/trailing whitespace
		if line == "" {
			continue // Skip empty lines
		}
		lines = append(lines, line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return lines, nil
}
