package main

import (
	"bufio"
	"fmt"
	"os"
)

// Types

const Tree = byte('#')

type Slope struct {
	x int
	y int
}

type TreeArea []string

func (a TreeArea) countTrees(slope Slope) int {
	var x, y, treeCount int
	patternWidth := len(a[0])
	for y < len(a) {
		// fmt.Printf("x: %d, y: %d \n", x, y)
		if a[y][x] == Tree {
			treeCount++
		}

		x += slope.x
		if x >= patternWidth {
			x -= patternWidth
		}
		y += slope.y
	}
	return treeCount
}

func (a TreeArea) Print() {
	for _, line := range a {
		fmt.Println(line)
	}
}

func readArea() TreeArea {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	var area TreeArea
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		area = append(area, scanner.Text())
	}

	return area
}

// Main procedures

func main() {
	area := readArea()
	solvePart1(area)
	solvePart2(area)
}

func solvePart1(area TreeArea) {
	slope := Slope{x: 3, y: 1}

	fmt.Println(area.countTrees(slope))
}

func solvePart2(area TreeArea) {
	slopes := [5]Slope{
		Slope{x: 1, y: 1},
		Slope{x: 3, y: 1},
		Slope{x: 5, y: 1},
		Slope{x: 7, y: 1},
		Slope{x: 1, y: 2},
	}
	treesPerSlope := make([]int, 5)

	for i, slope := range slopes {
		treesPerSlope[i] = area.countTrees(slope)
	}

	result := treesPerSlope[0]
	for _, treeCount := range treesPerSlope[1:] {
		result = result * treeCount
	}

	fmt.Println(result)
}
