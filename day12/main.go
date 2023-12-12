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
	fileString = strings.ReplaceAll(fileString, "\r", "")

	numbers := [][]int{}
	tiles := [][]string{}
	for _, line := range strings.Split(fileString, "\n") {
		split := strings.Split(line, " ")
		tile := []string{}
		numbs := []int{}
		for _, c := range split[0] {
			tile = append(tile, string(c))
		}
		for _, n := range strings.Split(split[1], ",") {
			num, _ := strconv.Atoi(n)
			numbs = append(numbs, num)
		}
		tiles = append(tiles, tile)
		numbers = append(numbers, numbs)
		/* fmt.Printf("Tiles\n")
		fmt.Println(tile)
		fmt.Printf("Numbs\n")
		fmt.Println(numbs) */
	}
	//fmt.Println(numbers)
	total := 0
	for y, line := range tiles {
		total += countPossibilities(line, numbers[y])
	}
	fmt.Printf("Part 1 is %d", total)
}

func countPossibilities(line []string, numbs []int) int {
	possiblesString := []string{""}
	for _, char := range line {
		if char == "?" {
			if len(possiblesString) > 0 {
				for i, s := range possiblesString {
					possiblesString[i] += "#"
					possiblesString = append(possiblesString, s+".")
				}
			}
		} else {
			for i := range possiblesString {
				possiblesString[i] += char
			}
		}
	}
	possibles := 0
	for _, s := range possiblesString {
		groups := splitBySameSubstrings(s)
		//fmt.Println(groups)
		result := checkGoals(groups, numbs)
		if result == true {
			possibles++
		}
	}
	/* if len(possiblesString) == 8 {
		fmt.Println(possiblesString)
	} */
	//fmt.Printf("string %s have %d possible %d and %d good\n", line, len(possiblesString), numbs, possibles)
	return possibles
}

func splitBySameSubstrings(s string) []string {
	var substrings []string
	if len(s) == 0 {
		return substrings
	}
	start := 0
	current := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != current {
			if string(current) == "#" {
				substrings = append(substrings, s[start:i])
			}
			start = i
			current = s[i]
		}
	}
	if string(current) == "#" {
		substrings = append(substrings, s[start:])
	}
	return substrings
}

func checkGoals(substrings []string, goals []int) bool {
	if len(substrings) != len(goals) {
		return false
	}
	for i := 0; i < len(substrings); i++ {
		if len(substrings[i]) != goals[i] {
			return false
		}
	}
	return true
}
