package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	Front rune = 'F'
	Back  rune = 'B'
	Left  rune = 'L'
	Right rune = 'R'

	Rows    int = 128
	Columns int = 8
)

func main() {
	input := readInput()
	solvePart1(input)
}

func readInput() []string {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	var inputLines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	return inputLines
}

func solvePart1(inputData []string) {
	seatIds := make([]int, 0)

	for _, line := range inputData {
		seatIds = append(seatIds, getSeatId(line))
	}

	fmt.Println(seatIds)

	var maxId int
	for _, seatId := range seatIds {
		if seatId > maxId {
			maxId = seatId
		}
	}
	fmt.Println(maxId)
}

func getSeatId(line string) int {
	row := BSP(line[:7], 0, Rows-1)
	row = row * 8

	col := BSP(line[7:], 0, Columns-1)
	// fmt.Println(r)
	// fmt.Println(r)
	// fmt.Println(r)

	return row + col
}

func BSP(rules string, l int, r int) int {
	n := len(rules)
	for _, bspRune := range rules[:n-1] {
		differential := float64(r - l)
		switch bspRune {
		case Front, Left:
			r -= int(math.Ceil(differential / 2.0))
		case Back, Right:
			l += int(math.Ceil(differential / 2.0))
		}
	}
	switch rune(rules[n-1]) {
	case Front, Left:
		return l
	case Back, Right:
		return r
	}

	// if this was real code this would be error handling
	return -1
}
