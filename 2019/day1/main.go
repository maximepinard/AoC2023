package main

import (
	"fmt"
	"os"

	"strconv"
	"strings"
)

func main() {
	fileContent, _ := os.ReadFile("input.txt")
	fileString := string(fileContent)

	part1, part2 := 0, 0
	for _, row := range strings.Split(fileString, "\n") {
		number, _ := strconv.Atoi(row)
		part1 += number/3 - 2
		part2 += calculateFuel(number)
	}
	fmt.Printf("Part 1 is %d\n", part1)
	fmt.Printf("Part 2 is %d\n", part2)
}

func calculateFuel(a int) int {
	b := (a / 3) - 2
	if b > 0 {
		b += calculateFuel(b)
	} else {
		b = 0
	}
	return b
}
