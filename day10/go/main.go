package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters := readAdapters()
	sort.Ints(adapters)
	phoneJolt := getPhoneJoltage(adapters)
	solve(adapters, phoneJolt)
	solvePart2(adapters, phoneJolt)
}

func readAdapters() []int {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	var adapters []int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		adapter, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, adapter)
	}

	return adapters
}

func solve(adapters []int, phoneJolt int) {
	var lastJoltage, onesDiff, threesDiff int
	for _, adapter := range adapters {
		diff := adapter - lastJoltage
		if diff == 1 {
			onesDiff++
		} else if diff == 3 {
			threesDiff++
		}
		lastJoltage = adapter
	}
	diff := phoneJolt - lastJoltage
	if diff == 1 {
		onesDiff++
	} else if diff == 3 {
		threesDiff++
	}

	fmt.Println(onesDiff * threesDiff)
}

func solvePart2(adapters []int, phoneJolt int) {
	adjList := make(map[int][]int)
	adjList[0] = make([]int, 0)

}

func getPhoneJoltage(adapters []int) int {
	maxAdapter := 0
	for _, adapter := range adapters {
		if adapter > maxAdapter {
			maxAdapter = adapter
		}
	}
	return maxAdapter + 3
}
