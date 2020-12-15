package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func main() {
	adapters := readAdapters()
	solve(adapters)
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

func solve(adapters []int) {
	currJolt := 0
	phoneJolt := getPhoneJoltage(adapters)

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
