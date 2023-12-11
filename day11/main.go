package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func appendGalaxies(fileString string) [][]string {
	galaxies := [][]string{}
	for _, line := range strings.Split(fileString, "\n") {
		galaxy := []string{}
		for _, c := range line {
			if string(c) == "." || string(c) == "#" {
				galaxy = append(galaxy, string(c))
			}
		}
		galaxies = append(galaxies, galaxy)
	}
	return galaxies
}

func expandGalaxies(Galaxies [][]string) ([]int, []int) {
	rowIndex := make([]int, len(Galaxies[0]))
	colIndex := make([]int, len(Galaxies))
	for i, row := range Galaxies {
		for j, cell := range row {
			if cell == "#" {
				colIndex[i]++
				rowIndex[j]++
			}
		}
	}
	fmt.Println("colIndex")
	fmt.Println(colIndex)
	fmt.Println("rowIndex")
	fmt.Println(rowIndex)
	return colIndex, rowIndex
}

type Galaxy struct {
	x, y *big.Int
}

func CalculateShortestPaths(Galaxies [][]string, colIndex []int, rowIndex []int) {
	TheGalaxiesPart1 := []Galaxy{}
	TheGalaxiesPart2 := []Galaxy{}
	shiftY := big.NewInt(0)
	shiftY_2 := big.NewInt(0)
	for y, row := range Galaxies {
		shiftX := big.NewInt(0)
		shiftX_2 := big.NewInt(0)
		if colIndex[y] == 0 {
			shiftY.Add(shiftY, big.NewInt(2-1))
			shiftY_2.Add(shiftY_2, big.NewInt(1000000-1))
		}
		for x, cell := range row {
			if rowIndex[x] == 0 {
				shiftX.Add(shiftX, big.NewInt(2-1))
				shiftX_2.Add(shiftX_2, big.NewInt(1000000-1))
			}
			if cell == "#" {
				galaxyX := new(big.Int).Add(big.NewInt(int64(x)), shiftX)
				galaxyY := new(big.Int).Add(big.NewInt(int64(y)), shiftY)
				galaxyX_2 := new(big.Int).Add(big.NewInt(int64(x)), shiftX_2)
				galaxyY_2 := new(big.Int).Add(big.NewInt(int64(y)), shiftY_2)
				TheGalaxiesPart1 = append(TheGalaxiesPart1, Galaxy{galaxyX, galaxyY})
				TheGalaxiesPart2 = append(TheGalaxiesPart2, Galaxy{galaxyX_2, galaxyY_2})
			}
		}
	}
	totalPart1 := big.NewInt(0)
	totalPart2 := big.NewInt(0)
	nbG := len(TheGalaxiesPart2)
	for i := 0; i < nbG; i++ {
		for j := i + 1; j < nbG; j++ {
			distance1 := shortestPath(TheGalaxiesPart1[i], TheGalaxiesPart1[j])
			distance2 := shortestPath(TheGalaxiesPart2[i], TheGalaxiesPart2[j])
			totalPart1.Add(totalPart1, distance1)
			totalPart2.Add(totalPart2, distance2)
		}
	}
	fmt.Printf("Part 1 is %d\n", totalPart1)
	fmt.Printf("Part 2 is %d\n", totalPart2)
}

func shortestPath(start, end Galaxy) *big.Int {
	deltaX := new(big.Int).Abs(new(big.Int).Sub(start.x, end.x))
	deltaY := new(big.Int).Abs(new(big.Int).Sub(start.y, end.y))
	return new(big.Int).Add(deltaX, deltaY)
}

func main() {
	fileContent, _ := os.ReadFile("input.txt")
	fileString := string(fileContent)
	Galaxies := appendGalaxies(fileString)
	colIndex, rowIndex := expandGalaxies(Galaxies)
	CalculateShortestPaths(Galaxies, colIndex, rowIndex)
}
