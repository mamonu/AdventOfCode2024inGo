package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(report []int) bool {
	if len(report) < 2 {
		return false // A single-level report is not valid
	}

	// Determine the direction of the first difference
	initialDiff := report[1] - report[0]
	if initialDiff == 0 || abs(initialDiff) > 3 {
		return false
	}
	increasing := initialDiff > 0

	// Check the rest of the report
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff == 0 || abs(diff) > 3 {
			return false // Adjacent levels differ by less than 1 or more than 3
		}
		if (increasing && diff < 0) || (!increasing && diff > 0) {
			return false // Direction changed
		}
	}

	return true
}

func isSafeWithDampener(report []int) bool {
	// First check if the report is already safe
	if isSafe(report) {
		return true
	}

	// Try removing each level one at a time
	for i := 0; i < len(report); i++ {
		// Create a new report excluding the i-th level
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)

		// Check if the modified report is safe
		if isSafe(modifiedReport) {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readReportsFromFile(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number in report: %s", part)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func main() {
	fileName := "day2.txt"

	// Read reports from file
	reports, err := readReportsFromFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Count safe reports
	safeCount := 0
	safeCountWithDampener := 0

	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
		if isSafeWithDampener(report) {
			safeCountWithDampener++
		}

	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
	fmt.Printf("Number of safe reports with dampener: %d\n", safeCountWithDampener)
}
