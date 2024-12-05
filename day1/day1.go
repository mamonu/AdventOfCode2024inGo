package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid number: %s", s))
	}
	return n
}

func readListsFromFile(fileName string) ([]int, []int, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}

	var leftList, rightList []int
	for _, line := range strings.Split(strings.TrimSpace(string(content)), "\n") {
		nums := strings.Fields(line)
		l, r := stringToInt(nums[0]), stringToInt(nums[1])
		leftList = append(leftList, l)
		rightList = append(rightList, r)
	}
	return leftList, rightList, nil
}

func calculateTotalDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)
	total := 0
	for i := range leftList {
		total += int(math.Abs(float64(leftList[i] - rightList[i])))
	}
	return total
}

func calculateSimilarityScore(leftList, rightList []int) int {
	// Create a map to count occurrences of each number in the right list
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	// Calculate the similarity score
	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}

	return similarityScore
}

func main() {
	fileName := "day1.txt"
	leftList, rightList, err := readListsFromFile(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total distance: %d\n", calculateTotalDistance(leftList, rightList))

	similarityScore := calculateSimilarityScore(leftList, rightList)
	fmt.Printf("Similarity score: %d\n", similarityScore)
}
