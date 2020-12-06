package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	solve(input)
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
	seatIds := make([]int, len(inputData))
	for i, line := range inputData {
		seatIds[i] = getSeatId(line)
	}

	var maxId int
	for _, seatId := range seatIds {
		if seatId > maxId {
			maxId = seatId
		}
	}
	fmt.Println(maxId)

	sort.Ints(seatIds)
	for i, seatId := range seatIds {
		if i == 0 {
			continue
		}

		if seatId-seatIds[i-1] > 1 {
			fmt.Println(seatIds[i-1], seatId)
		}
	}
}

func getSeatId(line string) int {
	row := BSP(line[:7], 0, Rows-1)
	row = row * 8

	col := BSP(line[7:], 0, Columns-1)

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
