/**
Advent of Code 2023
Maxime PINARD
using go
I am late to the party - made the 03/12/2023
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

/**
  PART 1
*/

func firstPart() {

	fileContent, err := os.ReadFile("puzzle.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var myArray []string

	err = json.Unmarshal(fileContent, &myArray)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	total := 0
	for i := 0; i < len(myArray); i++ {
		// fmt.Println(myArray[i])
		s := myArray[i]
		first := "0"
		last := "0"
		for j := 0; j < len(s); j++ {
			if s[j] > 48 && s[j] < 58 {
				if first == "0" {
					first = string(s[j])
					// fmt.Printf(first)
				}
				last = string(s[j])
			}
		}
		// fmt.Println(last)
		number, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Println("my bad !")
		}
		total = total + number
	}
	fmt.Println("Part 1")
	fmt.Println(total)
}

/**
  PART 2
*/

func numberToString(number string) string {
	val := "0"
	switch number {
	case "one":
		val = "1"
		break
	case "two":
		val = "2"
		break
	case "three":
		val = "3"
		break
	case "four":
		val = "4"
		break
	case "five":
		val = "5"
		break
	case "six":
		val = "6"
		break
	case "seven":
		val = "7"
		break
	case "eight":
		val = "8"
		break
	case "nine":
		val = "9"
		break
	}
	return val
}

func readFile() []string {
	fileContent, err := os.ReadFile("puzzle.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var myArray []string

	err = json.Unmarshal(fileContent, &myArray)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}
	return myArray
}

func secondPart() {

	total := 0

	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	myArray := readFile()

	for i := 0; i < len(myArray); i++ {
		// fmt.Println(myArray[i])
		row := myArray[i]

		// each row should have a first and last value
		first := "0"
		last := "0"

		// each character of row is looped through to search for a number
		for j := 0; j < len(row); j++ {

			// check the byte value of the char to see if it is a digit !
			if row[j] > 48 && row[j] < 58 {
				if first == "0" {
					first = string(row[j])
				}
				last = string(row[j])
			}

			// loop for each possible written digit: [one, two, ..., nine]
			for g := 0; g < len(numbers); g++ {
				// if the current index of the char of the row is big enough to search for that written digit, if so then do it
				if len(numbers[g])-1 <= j {
					test := myArray[i][j-(len(numbers[g])-1) : j+1] // [0:3] to get 012 because start is inclusive but end is exclusive
					val := numberToString(test)
					// if the value is found then go places
					if val != "0" {
						// if first was not overwritten then save that value in it
						if first == "0" {
							first = val
						}
						// replace last anyway
						last = val
					}
				}
			}
		}

		// fmt.Println(first + last)
		number, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Println("my bad !")
		}
		total = total + number
	}
	fmt.Println("Part 2")
	fmt.Println(total)
}

func main() {
	firstPart()
	secondPart()
}

/**
Advent of Code 2023

First day coding in GO

Things i have learned in GO

loop through each char of a string give the byte value of the char
bye value of digits are from 49 (0) to 57 (9)

string() to convert byte to string
strconv.Atoi() to convert string to int
...
os.readFile to read a file
json.Unmarshal() to readJson into something that match
no need for parenthesis in if and for statements

myString[0:3] is getting the first substring corresponding the first 3 char but the start 0 is inclusive for the end 3 is exclusive !

*/
