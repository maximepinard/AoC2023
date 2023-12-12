package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type state struct {
	pattern string
	numbers string
}

var cache = make(map[state]uint64)

func setCache(pattern string, numbers []uint8, value uint64) uint64 {
	cache[state{pattern, string(numbers)}] = value
	return value
}

func count(pattern string, numbers []uint8) uint64 {
	if len(pattern) == 0 && len(numbers) == 0 {
		return 1
	}

	if len(pattern) == 0 {
		return 0
	}

	// test cache
	if value, ok := cache[state{pattern, string(numbers)}]; ok {
		return value
	}

	if pattern[0] == '.' {
		res := count(pattern[1:], numbers)
		return setCache(pattern, numbers, res)
	}

	// cut branches
	var sum uint64
	for _, n := range numbers {
		sum += uint64(n)
	}
	if uint64(len(pattern)) < sum {
		res := 0
		return setCache(pattern, numbers, uint64(res))
	}

	if pattern[0] == '?' {
		res := count(pattern[1:], numbers) + count("#"+pattern[1:], numbers)
		return setCache(pattern, numbers, res)
	}

	if pattern[0] == '#' {
		if len(numbers) == 0 {
			res := 0
			return setCache(pattern, numbers, uint64(res))
		}

		n := numbers[0]
		indexDot := strings.Index(pattern, ".")
		if indexDot == -1 {
			indexDot = len(pattern)
		}
		if uint64(indexDot) < uint64(n) {
			// not enough # or ?
			res := 0
			return setCache(pattern, numbers, uint64(res))
		}

		// eat n # or ?
		remaining := pattern[n:]
		if len(remaining) == 0 {
			res := count(remaining, numbers[1:])
			return setCache(pattern, numbers, res)
		}

		if remaining[0] == '#' {
			// fail
			res := 0
			return setCache(pattern, numbers, uint64(res))
		}
		// remaining[0] == '.' || remaining[0] == '?'
		// eat first ? since it should be a .
		res := count(remaining[1:], numbers[1:])
		return setCache(pattern, numbers, res)
	}
	panic("unreachable")
}

func unfoldPattern(pattern string) string {
	var res = pattern
	for i := 0; i < 4; i++ {
		res = res + "?" + pattern
	}
	return res
}

func unfoldNumbers(numbers []uint8) []uint8 {
	var res []uint8
	for i := 0; i < 5; i++ {
		res = append(res, numbers...)
	}
	return res
}

func ToInt(s string) uint64 {
	res, _ := strconv.Atoi(s)
	return uint64(res)
}

func solve(input string, unfold bool) uint64 {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var res uint64
	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ',' })
		var pattern = fields[0]
		var numbers []uint8
		for _, field := range fields[1:] {
			numbers = append(numbers, uint8(ToInt(field)))
		}
		if unfold {
			pattern = unfoldPattern(pattern)
			numbers = unfoldNumbers(numbers)
		}
		res += count(pattern, numbers)
	}

	return res
}

func Part1(input string) uint64 {
	return solve(input, false)
}

func Part2(input string) uint64 {
	return solve(input, true)
}

func main() {
	start := time.Now()
	fileContent, _ := os.ReadFile("input.txt")
	inputDay := string(fileContent)
	inputDay = strings.ReplaceAll(inputDay, "\r", "")
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
