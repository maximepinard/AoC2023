/**
Advent of Code 2023
Maxime PINARD
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	source      int
	destination int
	length      int
}

func readFile() string {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return string(fileContent)
}

func whichType(s string) {
}

func partOne() {
	fileContent := readFile()

	// Split the lines
	fc := strings.Replace(fileContent, "  ", " ", -1)
	rows := strings.Split(fc, "\n")
	whichType := "seed"
	seedList := []int{}
	soilList := []Map{}
	fertilizerList := []Map{}
	waterList := []Map{}
	lightList := []Map{}
	temperatureList := []Map{}
	humidityList := []Map{}
	locationList := []Map{}
	// actual planted things
	soilPlantedList := []int{}
	fertilizerPlantedList := []int{}
	waterPlantedList := []int{}
	lightPlantedList := []int{}
	temperaturePlantedList := []int{}
	humidityPlantedList := []int{}
	locationPlantedList := []int{}
	for index, row := range rows {
		if strings.Contains(row, "map") {
			whichType = strings.Split(row, " ")[0]
		} else if strings.ContainsAny(row, "0123456789") {
			if index == 0 {
				temp := strings.Split(row, ":")[1]
				vals := strings.Split(temp, " ")
				for _, val := range vals {
					numb, err := strconv.Atoi(val)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						seedList = append(seedList, numb)
					}
				}
			} else {
				properties := strings.Split(row, " ")
				dest, err2 := strconv.Atoi(properties[0])
				if err2 != nil {
					fmt.Println(err2.Error())
				}
				source, err1 := strconv.Atoi(properties[1])
				if err1 != nil {
					fmt.Println(err1.Error())
				}
				length, err3 := strconv.Atoi(properties[2])
				if err3 != nil {
					fmt.Println(err3.Error())
				}
				if err1 != nil || err2 != nil || err3 != nil {

				} else {
					switch whichType {
					case "seed-to-soil":
						soilList = append(soilList, Map{source: source, destination: dest, length: length})
						break
					case "soil-to-fertilizer":
						fertilizerList = append(fertilizerList, Map{source: source, destination: dest, length: length})
						break
					case "fertilizer-to-water":
						waterList = append(waterList, Map{source: source, destination: dest, length: length})
						break
					case "water-to-light":
						lightList = append(lightList, Map{source: source, destination: dest, length: length})
						break
					case "light-to-temperature":
						temperatureList = append(temperatureList, Map{source: source, destination: dest, length: length})
						break
					case "temperature-to-humidity":
						humidityList = append(humidityList, Map{source: source, destination: dest, length: length})
						break
					case "humidity-to-location":
						locationList = append(locationList, Map{source: source, destination: dest, length: length})
						break
					}
				}
			}
		}
	}

	fmt.Printf("seed: ")
	/* fmt.Println(seedList)
	fmt.Printf("seed-to-soil: ")
	fmt.Println(soilList)
	fmt.Printf("soil-to-fertilizer: ")
	fmt.Println(fertilizerList)
	fmt.Printf("fertilizer-to-water: ")
	fmt.Println(waterList)
	fmt.Printf("water-to-light: ")
	fmt.Println(lightList)
	fmt.Printf("light-to-temperature: ")
	fmt.Println(temperatureList)
	fmt.Printf("temperature-to-humidity: ")
	fmt.Println(humidityList)
	fmt.Printf("humidity-to-location: ")
	fmt.Println(locationList) */

	// seed to soil
	soilPlantedList = getSeedNumbers(seedList, soilList)

	fmt.Printf("soilPlantedList: ")
	fmt.Println(soilPlantedList)

	// soil to fertilize
	fertilizerPlantedList = getSeedNumbers(soilPlantedList, fertilizerList)
	fmt.Printf("fertilizerPlantedList: ")
	fmt.Println(fertilizerPlantedList)

	// fertilize to water
	waterPlantedList = getSeedNumbers(fertilizerPlantedList, waterList)
	fmt.Printf("waterPlantedList: ")
	fmt.Println(waterPlantedList)

	// water to light
	lightPlantedList = getSeedNumbers(waterPlantedList, lightList)
	fmt.Printf("lightPlantedList: ")
	fmt.Println(lightPlantedList)

	// light to temp
	temperaturePlantedList = getSeedNumbers(lightPlantedList, temperatureList)
	fmt.Printf("temperaturePlantedList: ")
	fmt.Println(temperaturePlantedList)

	// temp to humidity
	humidityPlantedList = getSeedNumbers(temperaturePlantedList, humidityList)
	fmt.Printf("humidityPlantedList: ")
	fmt.Println(humidityPlantedList)

	// humidity to location
	locationPlantedList = getSeedNumbers(humidityPlantedList, locationList)
	fmt.Printf("locationPlantedList: ")
	fmt.Println(locationPlantedList)

	// Initialize min with the first element of the array
	min := locationPlantedList[0]

	// Iterate through the array and update min if a smaller element is found
	for _, value := range locationPlantedList {
		if value < min {
			min = value
		}
	}
	fmt.Println(min)
}

func getSeedNumbers(seedList []int, mapList []Map) []int {
	newIntList := []int{}
	for i := 0; i < len(seedList); i++ {
		planted := false
		for j := 0; j < len(mapList); j++ {
			startRange := mapList[j].source
			endRange := mapList[j].source + mapList[j].length
			if seedList[i] >= startRange && seedList[i] < endRange {
				shift := seedList[i] - startRange
				final := mapList[j].destination + shift
				planted = true
				newIntList = append(newIntList, final)
				// fmt.Printf("%d -> %d (%d, %d, %d, %d)\n", seedList[i], final, startRange, endRange, mapList[j].destination, mapList[j].length)
			}
		}
		if planted == false {
			newIntList = append(newIntList, seedList[i])
			fmt.Printf("%d -> %d\n", seedList[i], seedList[i])
		}
	}
	return newIntList
}

func partTwo() {

}

func main() {
	partOne()
	partTwo()
}
