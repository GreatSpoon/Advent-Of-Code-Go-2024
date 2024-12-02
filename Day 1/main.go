package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateTotalDistance(leftList []int, rightList []int) int {
	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0.0

	for i := 0; i < len(leftList); i++ {
		distance := math.Abs(float64(leftList[i]) - float64(rightList[i]))
		totalDistance += distance
	}

	return int(totalDistance)
}

func calculateSimilarityScore(leftList []int, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	similarityScore := 0

	for i := 0; i < len(leftList); i++ {
		itemInRightListCount := 0
		for _, item := range rightList {
			if item == leftList[i] {
				itemInRightListCount += 1
			}
		}
		similarityScore += leftList[i] * itemInRightListCount
	}

	return similarityScore
}

func readListsFromFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var leftList, rightList []int

	scanner := bufio.NewScanner(file)

	for lineNum := 0; scanner.Scan(); lineNum++ {
		line := scanner.Text()
		// Split the line into two numbers
		items := strings.Fields(line)
		if len(items) != 2 {
			return nil, nil, fmt.Errorf("invalid format on line %d: expected 2 numbers, got %d", lineNum+1, len(items))
		}

		leftNum, err := strconv.Atoi(items[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number on line %d: %v", lineNum+1, err)
		}

		rightNum, err := strconv.Atoi(items[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number on line %d: %v", lineNum+1, err)
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return leftList, rightList, nil
}

func main() {
	// File containing the lists
	filename := "input.txt"

	// Read the lists from the file
	leftList, rightList, err := readListsFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading lists from file: %v\n", err)
		return
	}

	// Check if both lists have the same length
	if len(leftList) != len(rightList) {
		fmt.Println("The two lists must have the same length.")
		return
	}

	// Calculate the total distance
	totalDistance := calculateTotalDistance(leftList, rightList)
	similarityScore := calculateSimilarityScore(leftList, rightList)

	// Output the result
	fmt.Printf("The total distance between the two lists is: %d\n", totalDistance)
	fmt.Printf("The total similarity score is : %d\n", similarityScore)
}
