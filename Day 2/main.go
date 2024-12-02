package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func safetyCheck(numOne int, numTwo int, rowType string) bool {
	if numOne > numTwo && rowType == "decreasing" {
		return false
	}
	if numOne < numTwo && rowType == "increasing" {
		return false
	}
	diffInInts := absInt(numTwo - numOne)
	if diffInInts < 1 || diffInInts > 3 {
		return false
	}
	return true
}

func checkValid(row []int) bool {
	rowType := ""
	previousNum := row[0]
	for i := 1; i < len(row); i++ {
		if row[i] < previousNum && rowType == "" {
			rowType = "decreasing"
		} else if row[i] > previousNum && rowType == "" {
			rowType = "increasing"
		}

		if safetyCheck(row[i], previousNum, rowType) == false {
			return false
		}

		previousNum = row[i]
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validRows := 0
	validRowsPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		stringRow := strings.Split(line, " ") // Replace ' ' with your delimiter
		intRow := make([]int, len(stringRow))

		for i, s := range stringRow {
			val, err := strconv.Atoi(s) // Convert string to int
			if err != nil {
				fmt.Printf("Error converting %s to int: %v\n", s, err)
				return
			}
			intRow[i] = val
		}
		validRow := checkValid(intRow)
		if validRow {
			validRows += 1
			validRowsPart2 += 1
		} else {
			for i := 0; i < len(intRow); i++ {
				newIntRow := make([]int, len(intRow)-1)
				for index, value := range intRow {
					if i != index {
						if index > i {
							index -= 1
						}
						newIntRow[index] = value
					}
				}
				if checkValid(newIntRow) {
					validRowsPart2 += 1
					break
				}
			}
		}

	}

	fmt.Printf("Part 1: %d\n", validRows)
	fmt.Printf("Part 2: %d\n", validRowsPart2)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
