package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

type Point struct {
	x, y int
}

// ReadGrid reads the map from the input file and returns the grid, the guard's starting position, and the initial direction.
func ReadGrid(filename string) ([][]rune, Point, Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, Point{}, Point{}, err
	}
	defer file.Close()

	var grid [][]rune
	var guardPosition Point
	var direction Point
	scanner := bufio.NewScanner(file)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := []rune(line)
		for x, char := range row {
			switch char {
			case '^':
				guardPosition = Point{x, y}
				direction = Point{0, -1} // Up
				row[x] = '.'             // Clear guard's initial position
			case 'v':
				guardPosition = Point{x, y}
				direction = Point{0, 1} // Down
				row[x] = '.'
			case '<':
				guardPosition = Point{x, y}
				direction = Point{-1, 0} // Left
				row[x] = '.'
			case '>':
				guardPosition = Point{x, y}
				direction = Point{1, 0} // Right
				row[x] = '.'
			}
		}
		grid = append(grid, row)
	}

	return grid, guardPosition, direction, nil
}

// DisplayGrid displays the current state of the grid with the guard's position and colored squares.
func DisplayGrid(grid [][]rune, guardPosition Point) {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	// Define colors for the grid. Lets make its a Zelda theme map :)
	const green = "\033[42m"  // Green background for `.`
	const brown = "\033[43m"  // Brown background for `#`
	const yellow = "\033[33m" // Yellow text for the guard
	const reset = "\033[0m"   // Reset to default colors

	// Render the grid with colors
	for y, row := range grid {
		for x, cell := range row {
			if guardPosition.x == x && guardPosition.y == y {
				// Guard's position
				fmt.Print(yellow + "G" + reset)
			} else if cell == '.' {
				// Green square for `.`
				fmt.Print(green + " " + reset)
			} else if cell == '#' {
				// Brown square for `#`
				fmt.Print(brown + " " + reset)
			} else {
				// Default display for other characters
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
}

func SimulatePatrol(grid [][]rune, guardPosition Point, direction Point, displayGrid bool) int {
	directions := []Point{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}

	// Find the index of the current direction in the directions slice.
	getDirectionIndex := func(d Point) int {
		for i, dir := range directions {
			if dir == d {
				return i
			}
		}
		return -1
	}

	visited := make(map[Point]bool)
	visited[guardPosition] = true

	// Simulate the guard's movement.
	for {

		// Only display the grid if the displayGrid flag is true
		if displayGrid {
			DisplayGrid(grid, guardPosition)
			time.Sleep(200 * time.Millisecond) // Add a delay for animation effect
		}

		nextPosition := Point{guardPosition.x + direction.x, guardPosition.y + direction.y}

		// Check if the guard is leaving the grid.
		if nextPosition.y < 0 || nextPosition.y >= len(grid) ||
			nextPosition.x < 0 || nextPosition.x >= len(grid[0]) {
			break
		}

		// Check for obstacles.
		if grid[nextPosition.y][nextPosition.x] == '#' {
			// Turn right.
			currentDirIndex := getDirectionIndex(direction)
			direction = directions[(currentDirIndex+1)%len(directions)]
		} else {
			// Move forward.
			guardPosition = nextPosition
			visited[guardPosition] = true
		}
	}

	return len(visited)
}

func main() {

	displayGrid := flag.Bool("display", false, "Toggle display grid animation on or off")
	flag.Parse()

	grid, guardPosition, direction, err := ReadGrid("day6.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// go run day6.go --display     : Display the grid animation (slower)
	// go run day6.go       : Just solve the pt1 puzzle and print the result
	result := SimulatePatrol(grid, guardPosition, direction, *displayGrid)
	fmt.Println("Distinct positions visited:", result)
}
