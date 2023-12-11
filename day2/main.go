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
	for i, row := range strings.Split(fileString, "\n") {
		game := strings.Split(row, ": ")       // Split game number from the rounds
		rounds := strings.Split(game[1], "; ") // Split each round from the other rounds
		num := map[string]int{}                // map each game color to a number
		for _, round := range rounds {
			colors := strings.Split(round, ", ") // Split each round from the other rounds
			for _, color := range colors {
				values := strings.Split(color, " ")
				number, _ := strconv.Atoi(values[0])
				num[values[1]] = max(num[values[1]], number) // get the max of each color for the game
			}
		}
		if num["red"] <= 12 && num["green"] <= 13 && num["blue"] <= 14 {
			part1 += i + 1
		}
		part2 += num["red"] * num["green"] * num["blue"]
	}
	fmt.Printf("Part 1 is %d\n", part1)
	fmt.Printf("Part 2 is %d\n", part2)
}
