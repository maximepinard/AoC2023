/**
Advent of Code 2023
Maxime PINARD
*/

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() string {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return string(fileContent)
}

type Matrix struct {
	Rows []Row
}

type Row struct {
	nums []int
}

/* for index, row := range rows {
	numbers := strings.Split(row, " ")
	if currentLength < len(numbers) || index == len(rows)-1 {
		if len(CurrentMatrix.Rows) > 0 {
			Matrixs = append(Matrixs, CurrentMatrix)
			CurrentMatrix = Matrix{}
		}
	}
	var nums []int
	for _, number := range numbers {
		intNumber, err := strconv.Atoi(number)
		if err == nil {
			nums = append(nums, intNumber)
		}
	}
	CurrentMatrix.Rows = append(CurrentMatrix.Rows, Row{nums: nums})
	currentLength = len(numbers)
} */

func partOne() {
	fileContent := readFile()

	// Split the lines
	re := regexp.MustCompile(`[ ]{2,}`)
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	modifiedString := re.ReplaceAllString(fileContent, " ")

	rows := strings.Split(modifiedString, "\n")
	var Matrixs []Matrix
	CurrentMatrix := Matrix{}
	for _, row := range rows {
		CurrentMatrix = Matrix{}
		numbers := strings.Split(row, " ")
		var nums []int
		for _, number := range numbers {
			intNumber, err := strconv.Atoi(number)
			if err == nil {
				nums = append(nums, intNumber)
			}
		}
		CurrentMatrix.Rows = append(CurrentMatrix.Rows, Row{nums: nums})
		Matrixs = append(Matrixs, CurrentMatrix)
	}
	var NewMatrixs []Matrix
	for _, CurrentMatrix := range Matrixs {
		index := 0
		for {
			all0 := true
			newRow := Row{}
			row := CurrentMatrix.Rows[index]
			for i := 0; i < len(row.nums)-1; i++ {
				if row.nums[i] != 0 {
					all0 = false
				}
				newRow.nums = append(newRow.nums, row.nums[i+1]-row.nums[i])
			}
			if all0 == false {
				CurrentMatrix.Rows = append(CurrentMatrix.Rows, newRow)
			} else {
				newRow = Row{}
				for i := 0; i < len(row.nums)-1; i++ {
					newRow.nums = append(newRow.nums, 0)
				}
				CurrentMatrix.Rows = append(CurrentMatrix.Rows, newRow)
				break
			}
			index++
		}
		NewMatrixs = append(NewMatrixs, CurrentMatrix)
	}
	fmt.Println(NewMatrixs)
	Total := 0
	for _, CurrentMatrix := range NewMatrixs {
		lastAdd := 0
		for _, row := range CurrentMatrix.Rows {
			lastIndex := len(row.nums) - 1
			lastAdd += row.nums[lastIndex]
		}
		fmt.Printf("last number is %d\n", lastAdd)
		Total += lastAdd
	}
	fmt.Printf("Part 1: The Total is %d\n", Total)
}

func main() {
	partOne()
	// partTwo()
	return
}
