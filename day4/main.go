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
	"strings"
)

func readFile() string {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return string(fileContent)
}

func findInRow(row string, start int, end int) bool {
	subString := row[start:end]
	// check each string of substring for the previous row
	for sc := 0; sc < len(subString); sc++ {
		// found not a digit nor .
		if (subString[sc] < 48 || subString[sc] > 57) && subString[sc] != 46 {
			return true
		}
	}
	return false
}

func partOne() {
	fileContent := readFile()

	// Split the lines
	fc := strings.Replace(fileContent, "  ", " ", -1)
	rows := strings.Split(fc, "\n")
	total := 0
	for _, row := range rows {
		cards := strings.Split(row, ": ")
		card := strings.Split(cards[1], " | ")

		winningNumbers := strings.Fields(card[0])
		gotNumbers := strings.Fields(card[1])

		/* fmt.Printf("\nWinning Numbers: %v\n", winningNumbers)
		fmt.Printf("Got Numbers: %v\n", gotNumbers) */
		val := 0
		for _, wn := range winningNumbers {
			for _, gn := range gotNumbers {
				if gn == wn {
					if val == 0 {
						val = 1
					} else {
						val *= 2
					}
					// fmt.Printf("- %s - = %d \n", gn, val)
				}
			}
		}
		total += val
	}
	fmt.Println(total)
}

func partTwo() {

}

func main() {
	partOne()
	partTwo()
}
