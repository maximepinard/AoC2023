/**
Advent of Code 2023
Maxime PINARD
using go
I am late to the party - made the 03/12/2023
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() string {
	fileContent, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return string(fileContent)
}

func partOne() {
	fileContent := readFile()
	total := 0
	// colorsStrings := [3]string{"red", "green", "blue"}

	// Split the lines
	result := strings.Split(fileContent, `
`)

	// Loop through 100 lines
	for i := 0; i < len(result); i++ {
		// Game init at sucess
		isOkay := 1
		// fmt.Println(result[i])
		game := strings.Split(result[i], `: `) // Split game number from the rounds
		rounds := strings.Split(game[1], `; `) // Split each round from the other rounds

		for _, round := range rounds {
			colors := strings.Split(round, `, `) // Split each round from the other rounds
			// val:= ""

			for _, color := range colors {

				values := strings.Split(color, " ")
				//fmt.Println(color)
				number, err := strconv.Atoi(values[0])
				if err != nil {
					fmt.Println("my bad !")
				}

				if values[1] == "red" && number > 12 {
					isOkay = 0
				}
				if values[1] == "green" && number > 13 {
					isOkay = 0
				}
				if values[1] == "blue" && number > 14 {
					isOkay = 0
				}
			}

		}
		if isOkay == 1 {
			total += i + 1 // Game 1 is index 0, so need to add 1 to index
		}
	}
	fmt.Printf("the total is: ")
	fmt.Println(total)
}

func partTwo() {
	fileContent := readFile()
	total := 0
	// colorsStrings := [3]string{"red", "green", "blue"}

	// Split the lines
	result := strings.Split(fileContent, `
`)

	// Loop through 100 lines
	for i := 0; i < len(result); i++ {
		// fmt.Println(result[i])
		game := strings.Split(result[i], `: `) // Split game number from the rounds
		rounds := strings.Split(game[1], `; `) // Split each round from the other rounds

		// find max value for each game for each color
		red := 0
		green := 0
		blue := 0
		for _, round := range rounds {
			colors := strings.Split(round, `, `) // Split each round from the other rounds
			// val:= ""

			for _, color := range colors {

				values := strings.Split(color, " ")
				// fmt.Println(color)
				number, err := strconv.Atoi(values[0])
				if err != nil {
					fmt.Println("my bad !")
				}

				if values[1] == "red" && number > red {
					red = number
				}
				if values[1] == "green" && number > green {
					green = number
				}
				if values[1] == "blue" && number > blue {
					blue = number
				}
			}

		}
		total += (red * green * blue) // Game 1 is index 0, so need to add 1 to index
	}
	fmt.Printf("the total is: ")
	fmt.Println(total)
}

func main() {
	partOne()
	partTwo()
}
