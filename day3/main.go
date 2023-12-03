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
	total := 0

	// Split the lines
	rows := strings.Split(fileContent, `
`)

	// Loop through lines
	for index, row := range rows {

		//find number
		num := ""
		for c := 0; c < len(row); c++ {
			//46 is .
			//48 to 57 are digit
			if row[c] > 47 && row[c] < 58 {
				num += string(row[c])
			}

			// if not number or last character in the row
			if row[c] < 48 || row[c] > 57 || c == len(row)-1 {
				if num != "" {

					// get the start and end
					start := c - 1 - len(num) // is inclusive
					if start < 0 {
						start = 0
					}
					end := c + 1 // is exclusive
					if end == len(row) {
						end = len(row) - 1
					}

					numb := findInRow(rows[index], start, end)
					if index > 0 && numb == false {
						numb = findInRow(rows[index-1], start, end)
					}
					if index+1 < len(rows) && numb == false {
						numb = findInRow(rows[index+1], start, end)
					}

					if numb == true {
						number, err := strconv.Atoi(num)
						if err != nil {
							fmt.Println("my bad !")
						}
						// fmt.Println(number)
						total += number
					}
				}
				num = ""
			}
		}
		//when found check for near symbol

		//if not forget them
	}

	fmt.Printf("the total is: ")
	fmt.Println(total)
}

func main() {
	partOne()
}
