package main

import (
	"fmt"
	"os"
	"strings"
)

type Pipe struct {
	x, y int
}

var seen [][]int

func appendTiles(fileString string) ([][]string, Pipe) {
	tiles := [][]string{}
	start := Pipe{}
	for y, line := range strings.Split(fileString, "\n") {
		pipes := []string{}
		for x, c := range line {
			if string(c) == "S" {
				start = Pipe{x, y}
			}
			pipes = append(pipes, string(c))
		}
		tiles = append(tiles, pipes)
		seen = append(seen, make([]int, len(line)))
	}
	return tiles, start
}

func findLoop(tiles [][]string, p Pipe) {
	if p.x < 0 || p.y < 0 || p.x >= len(tiles[0]) || p.y >= len(tiles) || seen[p.y][p.x] > 0 {
		return
	}
	seen[p.y][p.x] = 1
	direction := ""
	switch tiles[p.y][p.x] {
	case "|":
		direction = "ns"
		break
	case "-":
		direction = "ew"
		break
	case "7":
		direction = "ws"
		break
	case "F":
		direction = "es"
		break
	case "J":
		direction = "wn"
		break
	case "L":
		direction = "en"
		break
	case "S":
		direction = "nsew"
		break
	}
	for _, char := range direction {
		pipe := Pipe{}
		switch string(char) {
		case "n":
			pipe = Pipe{p.x, p.y - 1}
			break
		case "s":
			pipe = Pipe{p.x, p.y + 1}
			break
		case "e":
			pipe = Pipe{p.x + 1, p.y}
			break
		case "w":
			pipe = Pipe{p.x - 1, p.y}
			break
		}
		findLoop(tiles, pipe)
	}
	return
}

func CountLoop(tiles [][]string) (int, int) {
	loopLength := 0
	for y := range seen {
		findOutside(tiles, Pipe{y, 0}, "")
	}
	for x := range seen[0] {
		findOutside(tiles, Pipe{0, x}, "")
	}
	for y := range tiles {
		for x := range tiles[y] {
			switch tiles[y][x] {
			case "J":
				tiles[y][x] = "┘"
				break
			case "L":
				tiles[y][x] = "└"
				break
			case "7":
				tiles[y][x] = "┐"
				break
			case "F":
				tiles[y][x] = "┌"
				break
			case "|":
				tiles[y][x] = "│"
				break
			case "-":
				tiles[y][x] = "─"
				break
			}
		}
	}
	enclosedlength := 0
	for y, row := range seen {
		for x, value := range row {
			if value == 1 || value == 3 {
				loopLength++
				fmt.Printf("%s", tiles[y][x])
			} else if value == 2 {
				fmt.Printf("%s", " ")
			} else {
				enclosedlength++
				fmt.Printf("%s", "\033[0;31m█\033[0m")
			}
		}
		fmt.Printf("\n")
	}
	return loopLength / 2, enclosedlength
}

func findOutside(tiles [][]string, p Pipe, move string) {
	if p.x < 0 || p.y < 0 || p.y >= len(seen) || p.x >= len(seen[0]) || seen[p.y][p.x] >= 1 {
		return
	}
	seen[p.y][p.x] = 2
	direction := "nsew"
	for _, char := range direction {
		pipe := Pipe{}
		switch string(char) {
		case "n":
			pipe = Pipe{p.x, p.y - 1}
			break
		case "s":
			pipe = Pipe{p.x, p.y + 1}
			break
		case "e":
			pipe = Pipe{p.x + 1, p.y}
			break
		case "w":
			pipe = Pipe{p.x - 1, p.y}
			break
		}
		findOutside(tiles, pipe, string(char))
	}
	return
}

func main() {
	fileContent, _ := os.ReadFile("input.txt")
	fileString := string(fileContent)

	tiles, start := appendTiles(fileString)
	findLoop(tiles, Pipe{start.x, start.y})
	part1, part2 := CountLoop(tiles)
	fmt.Printf("Part 1 is %d\n", part1)
	fmt.Printf("Part 2 is %d\n", part2)
}
