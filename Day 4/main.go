package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var inputData [][]string
	xMasFoundCount := 0
	xMasCrossFoundCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		var row []string
		for _, value := range line {
			row = append(row, value)
		}
		inputData = append(inputData, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Part One
	for rowIndex := range len(inputData) {
		for colIndex := range len(inputData[rowIndex]) {
			if inputData[rowIndex][colIndex] != "X" {
				continue
			}
			for _, deltaRow := range []int{-1, 0, 1} {
				for _, deltaCol := range []int{-1, 0, 1} {
					if deltaRow == 0 && deltaCol == 0 {
						continue
					}
					if (0 <= rowIndex+3*deltaRow && rowIndex+3*deltaRow < len(inputData) && 0 <= colIndex+3*deltaCol && colIndex+3*deltaCol < len(inputData[0])) == false {
						continue
					}
					if inputData[rowIndex+deltaRow][colIndex+deltaCol] == "M" && inputData[rowIndex+2*deltaRow][colIndex+2*deltaCol] == "A" && inputData[rowIndex+3*deltaRow][colIndex+3*deltaCol] == "S" {
						xMasFoundCount += 1
					}
				}
			}
		}
	}

	// Part Two

	var theCorners [4]string
	for rowIndex := range (len(inputData)) - 1 {
		for colIndex := range (len(inputData[rowIndex])) - 1 {
			if rowIndex == 0 || colIndex == 0 {
				continue
			}
			if inputData[rowIndex][colIndex] != "A" {
				continue
			}
			theCorners[0] = inputData[rowIndex-1][colIndex-1]
			theCorners[1] = inputData[rowIndex-1][colIndex+1]
			theCorners[2] = inputData[rowIndex+1][colIndex+1]
			theCorners[3] = inputData[rowIndex+1][colIndex-1]
			theCornersString := strings.Join(theCorners[:], "")
			if theCornersString == "MMSS" || theCornersString == "MSSM" || theCornersString == "SSMM" || theCornersString == "SMMS" {
				xMasCrossFoundCount += 1
			}
		}
	}

	fmt.Println(xMasFoundCount)
	fmt.Println(xMasCrossFoundCount)
}
