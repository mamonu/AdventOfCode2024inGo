package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Part 1 - Total sum of valid multiplications: 181345830
// Part 2 - Total sum of enabled multiplications: 98729041
// readFile reads the input file line by line and returns the lines as a slice of strings.
func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// extractAndSumMultiplications scans lines for valid `mul(X,Y)` instructions,
// calculates their results, and returns the total sum.
func extractAndSumMultiplications(lines []string) (int, error) {
	// Regular expression to match valid mul(X,Y)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	total := 0

	for _, line := range lines {
		// Find all matches in the line
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// Parse the numbers X and Y
			x, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, fmt.Errorf("invalid number X: %s", match[1])
			}
			y, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, fmt.Errorf("invalid number Y: %s", match[2])
			}
			// Add the multiplication result
			total += x * y
		}
	}

	return total, nil
}

// extractAndSumConditionalMultiplications processes the lines considering `do()` and `don't()` instructions
func extractAndSumConditionalMultiplications(lines []string) (int, error) {
	// Regular expressions for the control instructions

	enabled := true // Initial state: `mul` instructions are enabled

	// Combine all lines into a single stream
	stream := strings.Join(lines, "")

	// Split the stream into chunks by instruction boundaries
	chunks := strings.Split(stream, ")")

	var enabledChunks []string // To collect only enabled chunks

	for _, chunk := range chunks {
		chunk = strings.TrimSpace(chunk) // Remove leading/trailing whitespace
		if strings.HasSuffix(chunk, "do(") {
			enabled = true
			continue
		}
		if strings.HasSuffix(chunk, "don't(") {
			enabled = false
			continue
		}

		// Add the chunk to enabledChunks if `enabled` is true
		if enabled {
			enabledChunks = append(enabledChunks, chunk+")") // Add back the ")"
		}
	}

	// Use `extractAndSumMultiplications` on the filtered chunks
	return extractAndSumMultiplications(enabledChunks)
}

func main() {

	fileName := "day3.txt"

	// Read the input file
	lines, err := readFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// day3 part 1: Process the lines to compute the sum of all multiplications
	total, err := extractAndSumMultiplications(lines)
	if err != nil {
		fmt.Printf("Error processing multiplications: %v\n", err)
		return
	}
	fmt.Printf("Part 1 - Total sum of valid multiplications: %d\n", total)

	// day3 part 2: Process the lines considering `do()` and `don't()` instructions
	conditionalTotal, err := extractAndSumConditionalMultiplications(lines)
	if err != nil {
		fmt.Printf("Error processing conditional multiplications: %v\n", err)
		return
	}
	fmt.Printf("Part 2 - Total sum of enabled multiplications: %d\n", conditionalTotal)
}
