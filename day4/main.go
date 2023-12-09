package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func readFile() string {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return string(fileContent)
}

func findInRow(row string, start int, end int) bool {
	subString := row[start:end]
	// check each string of substring for the previous row
	for sc := 0; sc < len(subString); sc++ {
		// found not a digit nor .
		if (subString[sc] < 48 || subString[sc] > 57) && subString[sc] != 46 {
			return true
		}
	}
	return false
}

func partOne(print bool) {
	fileContent := readFile()

	// Split the lines
	fc := strings.Replace(fileContent, "  ", " ", -1)
	rows := strings.Split(fc, "\n")
	total := 0
	for _, row := range rows {
		cards := strings.Split(row, ": ")
		card := strings.Split(cards[1], " | ")

		winningNumbers := strings.Fields(card[0])
		gotNumbers := strings.Fields(card[1])

		/* fmt.Printf("\nWinning Numbers: %v\n", winningNumbers)
		fmt.Printf("Got Numbers: %v\n", gotNumbers) */
		val := 0
		for _, wn := range winningNumbers {
			for _, gn := range gotNumbers {
				if gn == wn {
					if val == 0 {
						val = 1
					} else {
						val *= 2
					}
					// fmt.Printf("- %s - = %d \n", gn, val)
				}
			}
		}
		total += val
	}
	if print {
		fmt.Println(total)
	}
}

func partTwo(print bool) {
	fileContent := readFile()

	// Split the lines
	fc := strings.Replace(fileContent, "  ", " ", -1)
	rows := strings.Split(fc, "\n")
	stackCards := []int{}

	for i := 0; i < len(rows); i++ {
		stackCards = append(stackCards, 1)
	}

	for index, row := range rows {
		cards := strings.Split(row, ": ")
		card := strings.Split(cards[1], " | ")

		winningNumbers := strings.Fields(card[0])
		gotNumbers := strings.Fields(card[1])

		/* fmt.Printf("\nWinning Numbers: %v\n", winningNumbers)
		fmt.Printf("Got Numbers: %v\n", gotNumbers) */
		val := 0
		for _, wn := range winningNumbers {
			for _, gn := range gotNumbers {
				if gn == wn {
					val++
					//fmt.Printf("--- %s --- %d \n", gn, val)
				}
			}
		}
		for k := 1; k <= val; k++ {
			stackCards[index+k] += stackCards[index]
		}
	}
	total := 0
	//fmt.Println(stackCards)
	for i := 0; i < len(stackCards); i++ {
		total += stackCards[i]
	}
	if print {
		fmt.Println(total)
	}
}

func main() {
	partOne(true)
	partTwo(true)

	res1 := testing.Benchmark(BenchmarkPartOne)
	fmt.Printf("Memory allocations : %d \n", res1.MemAllocs)
	fmt.Printf("Number of bytes allocated: %d \n", res1.Bytes)
	fmt.Printf("Number of run: %d \n", res1.N)
	fmt.Printf("Time taken: %s \n", res1.T)
	fmt.Printf("Time taken per run: %f \n", res1.T.Seconds()/float64(res1.N))

	res2 := testing.Benchmark(BenchmarkPartTwo)
	fmt.Printf("Memory allocations : %d \n", res2.MemAllocs)
	fmt.Printf("Number of bytes allocated: %d \n", res2.Bytes)
	fmt.Printf("Number of run: %d \n", res2.N)
	fmt.Printf("Time taken: %s \n", res2.T)
	fmt.Printf("Time taken per run: %f \n", res2.T.Seconds()/float64(res2.N))
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partOne(false)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partTwo(false)
	}
}
