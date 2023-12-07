/**
Advent of Code 2023
Maxime PINARD
*/

package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readFile() string {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return string(fileContent)
}

// storing each car value as a hex
// 0023456789 TJQKA
// 0123456789 ABCDE
// and combo before
// 1 = high card, 2 = one pair, 3 = two pair, 4 = Three, 5 = full, 6 = four, 7 = five
// AA2456 = 1EE2456
// then convert to decimal to check strength

type Hand struct {
	value   int64
	initStr string
	detail  string
	bet     int
}

func valToHex(str string, bet string, part int) Hand {
	// fmt.Printf("my input is %s\n", str)
	hexa := ""
	for c := 0; c < len(str); c++ {
		if str[c] > 47 && str[c] < 58 { //if 0 to 9
			hexa += string(str[c])
		} else if string(str[c]) == "T" {
			hexa += "A"
		} else if string(str[c]) == "J" {
			if part == 1 {
				hexa += "B"
			} else {
				hexa += "1"
			}
		} else if string(str[c]) == "Q" {
			hexa += "C"
		} else if string(str[c]) == "K" {
			hexa += "D"
		} else if string(str[c]) == "A" {
			hexa += "E"
		}
	}
	/* if str == "JJJJJ" && part == 2 { // J is weakest alone
		hexa = "11111"
	} */
	print := 0
	/* if str == "AAAJA" || str == "AAAAT" || str == "KJKKJ" {
		print = 1
	} */
	hexa = evaluateStrength(hexa, part, str)
	if print == 1 {
		fmt.Printf("%s = %s\n", str, hexa)
	}
	// fmt.Printf("my result is %s\n", hexa)
	value, err := strconv.ParseInt(hexa, 16, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	num, err := strconv.Atoi(bet)
	if err != nil {
		fmt.Println(err.Error())
	}
	hand := Hand{value: value, initStr: str, detail: hexa, bet: num}
	if print == 1 {
		fmt.Println(hand)
	}
	return hand
}

func sortStringDescending(input string) string {
	runes := []rune(input)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] > runes[j]
	})
	return string(runes)
}

func evaluateStrength(hexa string, part int, str string) string {
	// 0 to 15
	// 0 to F
	// hexadecimal

	print := 0
	/* if str == "9TT7J" {
		print = 1
	} */

	count := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	hexaS := sortStringDescending(hexa)
	Joker := 0
	for c := 0; c < len(hexaS); c++ {
		value, err := strconv.ParseInt(string(hexaS[c]), 16, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		if string(hexaS[c]) == "1" && part == 2 {
			Joker++
		} else {
			count[value]++
		}
	}
	HighestCount := 0
	CountTwo := 0
	CountThree := 0
	for i := 0; i < len(count); i++ {
		if count[i] > HighestCount {
			HighestCount = count[i]
		}
		if count[i] == 3 {
			CountThree++
		}
		if count[i] == 2 {
			CountTwo++
		}
	}
	val := "0"
	if HighestCount == 5 || HighestCount+Joker == 5 {
		val = "7"
	} else if HighestCount == 4 || HighestCount+Joker == 4 {
		val = "6"
	} else if (Joker == 1 && CountTwo == 2) || (CountThree == 1 && CountTwo == 1) {
		val = "5"
	} else if CountThree == 1 || (CountTwo == 1 && Joker == 1) || Joker >= 2 {
		val = "4"
	} else if CountTwo == 2 {
		val = "3"
	} else if CountTwo == 1 || Joker >= 1 {
		val = "2"
	} else {
		val = "1"
	}
	if print == 1 {
		fmt.Printf("%s is ranked %s, 3: %d, 2: %d, h: %d\n", str, val, CountThree, CountTwo, HighestCount)
	}
	//
	return val + hexa
}

func sortByValue(handList []Hand) []Hand {
	sort.Slice(handList, func(i, j int) bool {
		return handList[i].value < handList[j].value
	})
	return handList
}

func partOne() {
	fileContent := readFile()

	// Split the lines
	re := regexp.MustCompile(`[ ]{2,}`)
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	modifiedString := re.ReplaceAllString(fileContent, " ")

	// Split the string by line breaks
	rows := strings.Split(modifiedString, "\n")
	hands := []Hand{}
	for _, row := range rows {
		split := strings.Split(row, " ")
		if len(split) < 2 {
			fmt.Printf("my row is %s and my split is %s \n", row, split)
		} else {
			newHand := valToHex(split[0], split[1], 1)
			hands = append(hands, newHand)
		}
	}
	hands = sortByValue(hands)
	//fmt.Println(hands)
	total := 0
	var previousValue int64 = 0
	for index, hand := range hands {
		if int64(previousValue) > hand.value {
			fmt.Printf("Erreur ")
			fmt.Println(hands[index-1])
			fmt.Println(hands[index])
			return
		}
		previousValue = hand.value
		value := (1 + index) * hand.bet
		// fmt.Printf("hand %s, %d * %d = %d\n", hand.initStr, 1+index, hand.bet, value)
		total += value
	}
	fmt.Printf("Part1 the total is : %d\n", total)
	fmt.Println(len(hands))
	return
}

func partTwo() {
	fileContent := readFile()

	// Split the lines
	re := regexp.MustCompile(`[ ]{2,}`)
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	modifiedString := re.ReplaceAllString(fileContent, " ")

	// Split the string by line breaks
	rows := strings.Split(modifiedString, "\n")
	hands := []Hand{}
	for _, row := range rows {
		split := strings.Split(row, " ")
		if len(split) < 2 {
			fmt.Printf("my row is %s and my split is %s \n", row, split)
		} else {
			newHand := valToHex(split[0], split[1], 2)
			hands = append(hands, newHand)
		}
	}
	hands = sortByValue(hands)
	//fmt.Println(hands)
	total := 0
	var previousValue int64 = 0
	for index, hand := range hands {
		if int64(previousValue) > hand.value {
			fmt.Printf("Erreur ")
			fmt.Println(hands[index-1])
			fmt.Println(hands[index])
			return
		}
		previousValue = hand.value
		value := (1 + index) * hand.bet
		//fmt.Printf("%s\n", hand.initStr)
		//fmt.Printf("hand %s, %d * %d = %d\n", hand.initStr, 1+index, hand.bet, value)
		total += value
	}
	fmt.Printf("Part2 the total is : %d\n", total)
	//fmt.Println(len(hands))
	return
}

func main() {
	partOne()
	partTwo()
	return
}
