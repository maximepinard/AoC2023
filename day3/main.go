package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	x, y   int
	number int
}
type Symbol struct {
	x, y int
}

func isNumberTouchingSymbol(num Number, gear Symbol) bool {
	strNumber := strconv.Itoa(num.number)
	if math.Abs(float64(num.y-gear.y)) < 2 {
		if gear.x >= num.x-1 && gear.x <= num.x+len(strNumber) {
			return true
		}
	}
	return false
}

func main() {
	fileContent, _ := os.ReadFile("input.txt")
	fileString := string(fileContent)

	symbols := []Symbol{}
	numbers := []Number{}
	rows := strings.Split(fileString, "\n")
	for y, row := range rows {
		num := ""
		for x, char := range row {
			if strings.ContainsAny(string(char), "0123456789") {
				num += string(char)
			} else if string(char) != "." {
				symbol := Symbol{x, y}
				symbols = append(symbols, symbol)
			}
			if strings.ContainsAny(string(char), "0123456789") == false || x == len(row)-1 {
				if num != "" {
					numb, _ := strconv.Atoi(num)
					number := Number{y: y, x: x - len(num), number: numb}
					numbers = append(numbers, number)
					num = ""
				}
			}
		}
	}

	part1, part2 := 0, 0
	for _, symbol := range symbols {
		selectedNumbers := []int{}
		for _, num := range numbers {
			if isNumberTouchingSymbol(num, symbol) {
				part1 += num.number
				if string(rows[symbol.y][symbol.x]) == "*" {
					selectedNumbers = append(selectedNumbers, num.number)

				}
			}
		}
		multiply := 1
		if len(selectedNumbers) > 1 {
			for _, numToMultiply := range selectedNumbers {
				multiply *= numToMultiply
			}
			part2 += multiply
		}
	}

	fmt.Printf("Part 1 is %d\n", part1)
	fmt.Printf("Part 2 is %d\n", part2)
}
