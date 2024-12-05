package main

import (
	"bufio"
	"fmt"
	"os"
)

// Directions for movement (row, col)
var directions = [8][2]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Down-right
	{1, -1},  // Down-left
	{-1, 1},  // Up-right
	{-1, -1}, // Up-left
}

// Function to load grid from a file
func loadGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func countOccurrences(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)
	count := 0

	// Helper function to check if a word exists
	// starting from (row, col) in a given direction
	isValid := func(row, col, dir int) bool {
		for i := 0; i < wordLen; i++ {
			newRow := row + i*directions[dir][0]
			newCol := col + i*directions[dir][1]
			// Check bounds
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				return false
			}
			// Check character match
			if grid[newRow][newCol] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	// Iterate through each cell in the grid
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Check in all 8 directions
			for dir := 0; dir < 8; dir++ {
				if isValid(row, col, dir) {
					count++
				}
			}
		}
	}

	return count
}

// Function to find all coordinates of 'A'
func findACoordinates(grid [][]rune) [][2]int {
	var coords [][2]int
	rows := len(grid)
	cols := len(grid[0])

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 'A' {
				coords = append(coords, [2]int{y, x})
			}
		}
	}
	return coords
}

// Function to count diagonal-cross patterns
func countDiagonalCrossPatterns(grid [][]rune) int {
	coords := findACoordinates(grid)
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for _, coord := range coords {
		y, x := coord[0], coord[1]

		// Skip if too close to the edge
		if y == 0 || x == 0 || y == rows-1 || x == cols-1 {
			continue
		}

		// Check diagonals
		diag1 := string(grid[y-1][x-1]) + string(grid[y+1][x+1]) // Top-left and bottom-right
		diag2 := string(grid[y-1][x+1]) + string(grid[y+1][x-1]) // Top-right and bottom-left

		// Increment count if both diagonals form a valid cross
		if (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM") {
			count++
		}
	}

	return count
}
func main() {

	// Load grid from file
	filename := "day4.txt"
	grid, err := loadGrid(filename)
	if err != nil {
		fmt.Printf("Error loading grid: %v\n", err)
		return
	}

	word := "XMAS"
	task1result := countOccurrences(grid, word)
	fmt.Printf("The word '%s' appears %d times in the grid.\n", word, task1result)

	task2result := countDiagonalCrossPatterns(grid)
	fmt.Printf("The number of X-MAS patterns in the grid is: %d\n", task2result)

}
