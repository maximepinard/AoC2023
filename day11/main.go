package main

import (
	"fmt"
	"math"
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

func expandGalaxies(Galaxies [][]string) [][]string {
	rowIndex := make([]int, len(Galaxies[0]))
	colIndex := make([]int, len(Galaxies))
	for i, row := range Galaxies {
		for j, cell := range row {
			if cell == "#" {
				rowIndex[i]++
				colIndex[j]++
			}
		}
	}
	fmt.Printf("Expanded row\n")
	shift := 0
	for i, row := range rowIndex {
		if row == 0 {
			fmt.Println(i + shift)
			toInsert := createStringArray(len(Galaxies[0]), ".")
			Galaxies = insertRow(Galaxies, i+shift, toInsert)
			shift++
		}
	}
	fmt.Printf("Expanded Col\n")
	shift = 0
	for j, col := range colIndex {
		if col == 0 {
			fmt.Println(j + shift)
			toInsert := createStringArray(len(Galaxies), ".")
			Galaxies = insertColumn(Galaxies, j+shift, toInsert)
			shift++
		}
	}
	return Galaxies
}

func createStringArray(size int, defaultValue string) []string {
	result := make([]string, size)
	for i := range result {
		result[i] = defaultValue
	}
	return result
}

func insertColumn(matrix [][]string, columnIndex int, column []string) [][]string {
	result := make([][]string, len(matrix))
	for i, row := range matrix {
		newRow := make([]string, len(row)+1)
		copy(newRow[:columnIndex], row[:columnIndex])
		newRow[columnIndex] = column[i]
		copy(newRow[columnIndex+1:], row[columnIndex:])
		result[i] = newRow
	}
	return result
}

func insertRow(matrix [][]string, rowIndex int, row []string) [][]string {
	result := make([][]string, len(matrix)+1)
	copy(result[:rowIndex], matrix[:rowIndex])
	result[rowIndex] = row
	copy(result[rowIndex+1:], matrix[rowIndex:])
	return result
}

type Galaxy struct {
	x, y int
}

func CalculateShortestPaths(Galaxies [][]string) {
	TheGalaxies := []Galaxy{}
	for y, row := range Galaxies {
		for x, cell := range row {
			if cell == "#" {
				TheGalaxies = append(TheGalaxies, Galaxy{x, y})
			}
		}
	}
	total := 0
	nbG := len(TheGalaxies)
	for i := 0; i < nbG; i++ {
		for j := i + 1; j < nbG; j++ {
			total += shortestPath(TheGalaxies[i], TheGalaxies[j])
		}
	}
	fmt.Printf("Part 1 is %d\n", total)
}

func shortestPath(start, end Galaxy) int {
	deltaX := math.Abs(float64(start.x - end.x))
	deltaY := math.Abs(float64(start.y - end.y))
	return int(deltaX + deltaY)
}

func main() {
	fileContent, _ := os.ReadFile("input.txt")
	fileString := string(fileContent)

	Galaxies := appendGalaxies(fileString)
	fmt.Printf("Original Galaxy\n")
	for _, row := range Galaxies {
		for _, cell := range row {
			fmt.Printf("%s", cell)
		}
		fmt.Printf("\n")
	}
	Galaxies = expandGalaxies(Galaxies)
	fmt.Printf("Expanded Galaxy\n")
	for _, row := range Galaxies {
		for _, cell := range row {
			fmt.Printf("%s", cell)
		}
		fmt.Printf("\n")
	}
	CalculateShortestPaths(Galaxies)

	// fmt.Println(Galaxies)
}
