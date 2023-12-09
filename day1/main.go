package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddFirstAndLastDigit(fileString string, part int) {
	result := 0
	for _, word := range strings.Fields(fileString) {
		indexFirst := strings.IndexAny(word, "123456789")
		indexLast := strings.LastIndexAny(word, "123456789")
		first := string(word[indexFirst])
		last := string(word[indexLast])
		num, _ := strconv.Atoi(first + last)
		result += num
	}
	fmt.Printf("Part %d is %d\n", part, result)
}

func ReplaceDigitStringByInt(fileString string) {
	stringToReplace := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for index, s := range stringToReplace {
		convert := strconv.Itoa(index + 1) // convert the index [1] of two => [1] + 1 to string
		firstChar := string(s[0])          // seven = s7n so that sevenine = s7nine don't break when two digit strings are merged
		lastChar := string(s[len(s)-1])
		replace := firstChar + convert + lastChar
		fileString = strings.ReplaceAll(fileString, s, replace)
	}
	AddFirstAndLastDigit(fileString, 2)
}

func main() {
	fileContent, _ := os.ReadFile("input.txt")
	fileString := string(fileContent)
	// Part 1
	AddFirstAndLastDigit(fileString, 1)
	// Part 2
	ReplaceDigitStringByInt(fileString)
}
