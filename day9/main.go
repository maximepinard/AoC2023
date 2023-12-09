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

func partOne(reverse bool) {
	fileContent := readFile()

	// Split the lines
	re := regexp.MustCompile(`[ ]{2,}`)
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	modifiedString := re.ReplaceAllString(fileContent, " ")

	rows := strings.Split(modifiedString, "\n")
	var Matrixs []Matrix
	CurrentMatrix := Matrix{}
	// Read and parse
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
	// Recreate history
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
				break
			}
			index++
		}
		NewMatrixs = append(NewMatrixs, CurrentMatrix)
	}
	// Calculate new column and Add them up
	//fmt.Println(NewMatrixs)
	Total := 0
	for _, CurrentMatrix := range NewMatrixs {
		lastAdd := 0
		for in := len(CurrentMatrix.Rows) - 1; in >= 0; in-- {
			row := CurrentMatrix.Rows[in]
			lastIndex := len(row.nums) - 1
			if reverse == true {
				lastIndex = 0
				lastAdd = row.nums[lastIndex] - lastAdd
			} else {
				lastAdd += row.nums[lastIndex]
			}
			//fmt.Printf("row %d is num %d equal %d\n", in, row.nums[lastIndex], lastAdd)
		}
		//fmt.Printf("last number is %d\n", lastAdd)
		Total += lastAdd
	}
	if reverse == true {
		fmt.Printf("Part 2: The Total is %d\n", Total)
	} else {
		fmt.Printf("Part 1: The Total is %d\n", Total)
	}
}

func partTwo() {
	reverse := true
	partOne(reverse)
}

func main() {
	reverse := false
	partOne(reverse)
	partTwo()
	return
}
