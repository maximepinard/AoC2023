/**
Advent of Code 2023
Maxime PINARD
*/

package main

import (
	"fmt"
	"math/big"
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

func partOne() {
	fileContent := readFile()

	// Split the lines
	re := regexp.MustCompile(`[ ]{2,}`)
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	modifiedString := re.ReplaceAllString(fileContent, " ")

	// Split the string by line breaks
	rows := strings.Split(modifiedString, "\n")
	distances := []int{}
	times := []int{}
	fmt.Printf("'%s'\n", rows[0])
	fmt.Printf("'%s'\n", rows[1])
	for index, row := range rows {
		if index == 0 {
			temp := strings.Split(row, ": ")[1]
			stringArray := strings.Split(temp, " ")
			for _, str := range stringArray {
				// Attempt to convert string to int
				num, err := strconv.Atoi(str)
				if err != nil {
					fmt.Printf("Error converting string '%s' to int: %s\n", str, err)
					// Handle error (if needed)
					continue // Skip to the next string if conversion fails
				}
				// Append the parsed int to the int array
				times = append(times, num)
			}
		} else if index == 1 {
			temp := strings.Split(row, ": ")[1]
			stringArray := strings.Split(temp, " ")
			for _, str := range stringArray {
				// Attempt to convert string to int
				num, err := strconv.Atoi(str)
				if err != nil {
					fmt.Printf("Error converting string '%s' to int: %s\n", str, err)
					// Handle error (if needed)
					continue // Skip to the next string if conversion fails
				}
				// Append the parsed int to the int array
				distances = append(distances, num)
			}
		}
	}
	fmt.Printf("times ")
	fmt.Printf("'%d'\n", times)
	fmt.Printf("distances ")
	fmt.Printf("'%d'\n", distances)
	total := 1
	if len(distances) != len(times) {
		fmt.Println("length do not match")
	} else {
		for i := 0; i < len(distances); i++ {
			distanceToBeat := distances[i]
			raceTime := times[i]
			speed := 0
			val := 0
			for t := 0; t < raceTime; t++ {
				if speed*(raceTime-t) > distanceToBeat {
					val++
				}
				speed++
			}
			if val > 0 {
				fmt.Printf("val for %d is : %d\n", i, val)
				total *= val
			}
		}
	}
	fmt.Printf("total is : %d", total)
	return
}

func partTwo() {
	fileContent := readFile()

	// Split the lines
	re := regexp.MustCompile(`[ ]{2,}`)
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	modifiedString := re.ReplaceAllString(fileContent, " ")

	// Split the string by line breaks
	rows := strings.Split(modifiedString, "\n")
	distances := []*big.Int{}
	times := []int{}
	fmt.Printf("'%s'\n", rows[0])
	fmt.Printf("'%s'\n", rows[1])
	for index, row := range rows {
		if index == 0 {
			temp := strings.Split(row, ": ")[1]
			str := strings.ReplaceAll(temp, " ", "")
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting string '%s' to int: %s\n", str, err)
				// Handle error (if needed)
				continue // Skip to the next string if conversion fails
			}
			// Append the parsed int to the int array
			times = append(times, num)
		} else if index == 1 {
			temp := strings.Split(row, ": ")[1]
			str := strings.ReplaceAll(temp, " ", "")
			num := new(big.Int)
			num.SetString(str, 10)
			// Append the parsed int to the int array
			distances = append(distances, num)
		}
	}
	fmt.Printf("times ")
	fmt.Printf("'%d'\n", times)
	fmt.Printf("distances ")
	fmt.Printf("'%d'\n", distances)
	total := 1
	if len(distances) != len(times) {
		fmt.Println("length do not match")
	} else {
		for i := 0; i < len(distances); i++ {
			distanceToBeat := distances[i]
			raceTime := times[i]
			speed := 0
			val := 0
			for t := 1; t <= raceTime; t++ {
				BigTime := new(big.Int)
				BigTime.SetInt64(int64(raceTime - t))
				speed++
				result := new(big.Int).Mul(new(big.Int).SetInt64(int64(speed)), new(big.Int).SetInt64(int64(raceTime-t)))
				if result.Cmp(distanceToBeat) > 0 {
					val++
				}
			}
			if val > 0 {
				fmt.Printf("val for %d is : %d\n", i, val)
				total *= val
			}
		}
	}
	fmt.Printf("total is : %d", total)
	return
}

func main() {
	// partOne()
	partTwo()
}
