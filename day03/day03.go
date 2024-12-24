package main

import (
	"adventofcode2024/reader"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Puzzle1(program string) int {
	fmt.Println("Puzzle 1")

	// regex pattern to match the mul function
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(program, -1)

	result := 0

	for _, match := range matches {
		val := mulValue(match)
		//fmt.Println("Match:", match, "Value:", val)
		result += val
	}

	return result
}
func Puzzle2(program string) int {
	fmt.Println("Puzzle 2")
	// regex pattern to match the mul function
	pattern := `mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(program, -1)

	result := 0
	mulEnabled := true //  At the beginning of the program, mul instructions are enabled.

	for _, match := range matches {
		val := 0
		if match[0] == "do()" {
			mulEnabled = true
			continue
		}
		if match[0] == "don't()" {
			mulEnabled = false
			continue
		}
		if mulEnabled {
			val = mulValue(match)
		}

		result += val
	}

	return result
}

func mulValue(match []string) int {
	a, _ := strconv.Atoi(match[1])
	b, _ := strconv.Atoi(match[2])
	return a * b
}

func data(file string) string {
	lines, err := reader.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		log.Fatal(err)
	}

	return strings.Join(lines[:], "")
}

func main() {
	// Check if a filename is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	// Get the filename from the command-line arguments
	filename := os.Args[1]

	fmt.Println("Hello, Advent of Code 2024 - Day 3!")

	result1 := Puzzle1(data(filename))
	fmt.Println("Puzzle 1:", result1)
	result2 := Puzzle2(data(filename))
	fmt.Println("Puzzle 2:", result2)
}
