/**
Advent of Code 2023
Maxime PINARD
*/

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func readFile() string {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return string(fileContent)
}

type Node struct {
	name        string
	Left, Right *Node
	left, right string
}

func partOne() {
	fileContent := readFile()

	// Split the lines
	re := regexp.MustCompile(`[ ]{2,}`)
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	modifiedString := re.ReplaceAllString(fileContent, " ")

	// Split the string by line breaks
	rows := strings.Split(modifiedString, "\n")
	nodes := []*Node{}
	instructions := ""
	var currentNode *Node
	for index, row := range rows {
		if index == 0 {
			instructions = row
		} else if index > 1 {
			split := strings.Split(row, " = ")
			// fmt.Println(split)
			parts := strings.Split(split[1], ", ")
			left := strings.Replace(parts[0], "(", "", 1)
			right := strings.Replace(parts[1], ")", "", 1)
			node := &Node{name: split[0], left: left, right: right, Left: nil, Right: nil}
			if node.name == "AAA" {
				currentNode = node
			}
			nodes = append(nodes, node)
		}
	}
	linkNodes(nodes)

	found := 0
	loop := 1
	for {
		for c := 0; c < len(instructions); c++ {
			if string(instructions[c]) == "L" {
				// fmt.Println("L")
				if currentNode.Left != nil {
					currentNode = currentNode.Left
				} else {
					fmt.Println("Left node is nil", currentNode)
					break
				}
			} else if string(instructions[c]) == "R" {
				// fmt.Println("R")
				if currentNode.Right != nil {
					currentNode = currentNode.Right
				} else {
					fmt.Println("Right node is nil", currentNode)
					break
				}
			} else {
				fmt.Println("Error: Invalid instruction")
			}
			if currentNode.name == "ZZZ" {
				found = (c + 1) * loop
				fmt.Printf("reached ZZZ in %d\n", found)
				break
			}
		}
		if found > 0 || currentNode == nil || loop*len(instructions) > 40000 {
			break
		}
		loop++
	}

	fmt.Printf("Part 1: currentNode is %s in %d, %d\n", currentNode.name, found, loop)

	return
}

func linkNodes(nodeList []*Node) *Node {
	nodeMap := make(map[string]*Node)

	// Step 1: Create nodes and store them in the map
	for _, n := range nodeList {
		nodeMap[n.name] = n
	}

	// Step 2: Link nodes using left and right names
	for _, n := range nodeList {
		if n.left != "" {
			n.Left = nodeMap[n.left]
		}
		if n.right != "" {
			n.Right = nodeMap[n.right]
		}
	}

	// Return the root node assuming the first node is the root
	return nodeList[0]
}

func findName(root *Node, target string) *Node {
	if root == nil {
		return nil
	}

	if root.name == target {
		return root
	}

	leftResult := findName(root.Left, target)
	if leftResult != nil {
		return leftResult
	}

	rightResult := findName(root.Right, target)
	if rightResult != nil {
		return rightResult
	}

	return nil
}

func findNodeByName(nodes []Node, nodeName string) *Node {
	for _, node := range nodes {
		if node.name == nodeName {
			return &node
		}
	}
	fmt.Println("BIG error")
	return nil
}

func findNode(root *Node, target string) *Node {
	if root == nil {
		return nil
	}

	if root.name == target {
		return root
	}

	leftResult := findNode(root.Left, target)
	if leftResult != nil {
		return leftResult
	}

	rightResult := findNode(root.Right, target)
	if rightResult != nil {
		return rightResult
	}

	return nil
}

func partTwo() {
	fileContent := readFile()

	// Split the lines
	re := regexp.MustCompile(`[ ]{2,}`)
	fileContent = strings.ReplaceAll(fileContent, "\r", "")
	modifiedString := re.ReplaceAllString(fileContent, " ")

	// Split the string by line breaks
	rows := strings.Split(modifiedString, "\n")
	nodes := []*Node{}
	instructions := ""
	var startNode []*Node
	for index, row := range rows {
		if index == 0 {
			instructions = row
		} else if index > 1 {
			split := strings.Split(row, " = ")
			// fmt.Println(split)
			parts := strings.Split(split[1], ", ")
			left := strings.Replace(parts[0], "(", "", 1)
			right := strings.Replace(parts[1], ")", "", 1)
			node := &Node{name: split[0], left: left, right: right, Left: nil, Right: nil}
			if string(node.name[2]) == "A" {
				startNode = append(startNode, node)
			}
			nodes = append(nodes, node)
		}
	}
	linkNodes(nodes)

	fmt.Println(startNode)

	loop := 1
	numbers := []int{}
	for i := 0; i < len(startNode); i++ {
		fmt.Println(startNode[i].name)
		numbers = append(numbers, -1)
	}
	for {
		for c := 0; c < len(instructions); c++ {
			if string(instructions[c]) == "L" {
				// fmt.Println("L")
				for index, node := range startNode {
					if startNode[index].Left != nil {
						startNode[index] = node.Left
					} else {
						fmt.Println("Left node is nil", startNode[index])
						break
					}
				}
			} else if string(instructions[c]) == "R" {
				// fmt.Println("R")
				for index, node := range startNode {
					if startNode[index].Right != nil {
						startNode[index] = node.Right
					} else {
						fmt.Println("Right node is nil", startNode[index])
						break
					}
				}
			} else {
				fmt.Println("Error: Invalid instruction")
			}
			for y, node := range startNode {
				if string(node.name[2]) == "Z" {
					if numbers[y] < 0 {
						fmt.Println(node.name)
						numbers[y] = (c + 1) * loop
					}
					// fmt.Printf("reached Z (%s) in %d\n", node.name, allFound)
				}
			}
		}
		allFound := 1
		for _, number := range numbers {
			if number < 0 {
				allFound = 0
			}
		}
		if allFound > 0 || loop > 2100000000 {
			break
		}
		loop++
	}

	fmt.Println(loop)
	fmt.Println(numbers)
	result := lcm(numbers[0], numbers[1])

	for i := 2; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	fmt.Printf("Part 2: step %d\n", result)

	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	partOne()
	partTwo()
	return
}
