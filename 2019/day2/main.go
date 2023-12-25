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
	positions := strings.Split(fileString, ",")
	intPositions := make([]int, len(positions))
	// convert postions array of string to array of int
	for i := range positions {
		intPositions[i], _ = strconv.Atoi(positions[i])
	}
	// make a copy of intPositions array of int
	intPositionsCopy := make([]int, len(intPositions))
	copy(intPositionsCopy, intPositions)
	part1 = LoopIntCodes(intPositionsCopy, 12, 2)
	fmt.Printf("Part 1 is %d\n", part1)
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			noun := i
			verb := j
			intPositionsCopy = make([]int, len(intPositions))
			copy(intPositionsCopy, intPositions)
			part2 = LoopIntCodes(intPositionsCopy, noun, verb)
			if part2 == 19690720 {
				fmt.Printf("Part 2 is %d\n", 100*noun+verb)
				goto end
			}
		}
	}
end:
	return
}

func LoopIntCodes(positions []int, noun int, verb int) int {
	positions[1] = noun
	positions[2] = verb
	for i := 0; i < len(positions); i += 4 {
		switch positions[i] {
		case 1:
			positions[positions[i+3]] = positions[positions[i+1]] + positions[positions[i+2]]
		case 2:
			positions[positions[i+3]] = positions[positions[i+1]] * positions[positions[i+2]]
		case 99:
			return positions[0]
		default:
			return 0
		}
	}
	return 0
}
