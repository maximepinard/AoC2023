package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fileContent, _ := os.ReadFile("input.txt")
	fileString := strings.ReplaceAll(string(fileContent), "\r", "")
	var part1, part2 uint64
	for _, line := range strings.Split(fileString, "\n") {
		split := strings.Fields(line)
		goal := strings.Split(split[1], ",")
		numbs := make([]int, len(goal))
		for i, n := range goal {
			num, _ := strconv.Atoi(n)
			numbs[i] = num
		}
		res1, res2 := countPossibilities(split[0], numbs)
		part1 += res1
		part2 += res2
	}
	fmt.Printf("Part 1 is %d\n", part1)
	fmt.Printf("Part 2 is %d\n", part2)
	fmt.Println(time.Since(start))
}

var cache = make(map[string]uint64)

func generateCombinations(s []rune, index int, numbs []int) uint64 {
	var count uint64
	key := string(s) + fmt.Sprintf("%d", numbs)
	if n, ok := cache[key]; ok {
		return n
	}

	if index == len(s) {
		var hashLengths []int
		currentLength := 0
		for _, r := range s {
			if r == '#' {
				currentLength++
			} else if currentLength > 0 {
				hashLengths = append(hashLengths, currentLength)
				currentLength = 0
			}
		}
		if currentLength > 0 {
			hashLengths = append(hashLengths, currentLength)
		}
		if len(hashLengths) == len(numbs) && reflect.DeepEqual(hashLengths, numbs) {
			count++
		}
		cache[key] = count
		return count
	}

	if s[index] == '?' {
		for _, r := range []rune{'.', '#'} {
			s[index] = r
			count += generateCombinations(s, index+1, numbs)
			cache[key] = count
			s[index] = '?'
		}
	} else {
		count += generateCombinations(s, index+1, numbs)
		cache[key] = count
	}
	cache[key] = count
	return count
}

/*
func mergeSlices(slices ...[]int) []int {
	var merged []int
	for _, s := range slices {
		merged = append(merged, s...)
	}
	return merged
} */

func countPossibilities(line string, numbs []int) (uint64, uint64) {
	strRunes := []rune(line)
	count := generateCombinations(strRunes, 0, numbs)
	// Part 2 run forever
	/* strRunes2 := []rune(line + "?" + line + "?" + line + "?" + line + "?" + line)
	numbs2 := mergeSlices(numbs, numbs, numbs, numbs, numbs)
	count2 := generateCombinations(strRunes2, 0, numbs2) */
	return count, 0
}
