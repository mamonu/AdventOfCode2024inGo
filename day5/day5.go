package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) (map[int][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(map[int][]int)
	var updates [][]int

	// Parse rules
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		rules[from] = append(rules[from], to)
	}

	// Parse updates
	for scanner.Scan() {
		line := scanner.Text()
		update := []int{}
		for _, numStr := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(numStr)
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	return rules, updates, nil
}

func isValidUpdate(update []int, rules map[int][]int) bool {
	// Create a map of positions for quick lookup
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	// Validate rules
	for from, tos := range rules {
		if posFrom, ok := position[from]; ok {
			for _, to := range tos {
				if posTo, exists := position[to]; exists {
					// If "from" appears after "to", the rule is violated
					if posFrom > posTo {
						return false
					}
				}
			}
		}
	}

	return true
}

// Function to compute the middle page of a set of pages
func middlePage(update []int) int {
	return update[len(update)/2]
}

func main() {

	filename := "day5.txt"
	rules, updates, err := parseInput(filename)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	sumOfMiddlePages := 0

	// Process each update
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			sumOfMiddlePages += middlePage(update)
		}
	}

	fmt.Printf("The sum of the middle pages of correctly ordered updates is: %d", sumOfMiddlePages)
}
