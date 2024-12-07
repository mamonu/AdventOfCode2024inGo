package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Concatenates two numbers / added for part 2
func concatenateNums(left, right int) int {
	leftStr := strconv.Itoa(left)
	rightStr := strconv.Itoa(right)
	concatenated, _ := strconv.Atoi(leftStr + rightStr)
	return concatenated
}

func evaluate(numbers []int, operators []string) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		switch operators[i-1] {
		case "+":
			result += numbers[i]
		case "*":
			result *= numbers[i]
		case "||":
			// Concatenate digits of the two numbers
			result = concatenateNums(result, numbers[i])
		}
	}
	return result
}

func generateOpCombinations(numbers []int, target int, operators []string, currentIndex int, includeConcatenation bool, isMatchFound *bool) {
	if currentIndex == len(numbers)-1 {
		if evaluate(numbers, operators) == target {
			*isMatchFound = true
		}
		return
	}

	// Try "+" operator
	operators[currentIndex] = "+"
	generateOpCombinations(numbers, target, operators, currentIndex+1, includeConcatenation, isMatchFound)

	// Try "*" operator
	operators[currentIndex] = "*"
	generateOpCombinations(numbers, target, operators, currentIndex+1, includeConcatenation, isMatchFound)

	// Try "||" operator only if allowed/ added for part 2
	if includeConcatenation {
		operators[currentIndex] = "||"
		generateOpCombinations(numbers, target, operators, currentIndex+1, includeConcatenation, isMatchFound)
	}
}

func isValid(numbers []int, target int, includeConcatenation bool) bool {
	operators := make([]string, len(numbers)-1) // operators for the current combination
	matchFound := false                         // Flag if a valid combination was found
	generateOpCombinations(numbers, target, operators, 0, includeConcatenation, &matchFound)
	return matchFound
}

func main() {
	file, err := os.Open("day7.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	// Define the command-line flag
	includeConcatenation := flag.Bool("includeConcatenation", false, "Include concatenation (||) operator in the calculations")
	flag.Parse() // Parse the command-line flags

	scanner := bufio.NewScanner(file)
	totalCalibrationSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			continue
		}

		numberStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numberStrings))
		for i, numStr := range numberStrings {
			numbers[i], _ = strconv.Atoi(numStr)
		}

		// for part 1: go run day7.go -includeConcatenation=false or go run day7.go (false is the default value)
		// for part 2: go run day7.go -includeConcatenation=true

		if isValid(numbers, target, *includeConcatenation) {
			totalCalibrationSum += target
		}
	}

	fmt.Println("Total Calibration Result:", totalCalibrationSum)
}
