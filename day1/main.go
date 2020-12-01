package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := getNumbers()
	numberSet := makeNumberSet(numbers)

	for _, number := range numbers {
		complement := 2020 - number
		if _, ok := numberSet[complement]; ok {
			fmt.Println(number * complement)
			break
		}
	}
}

func getNumbers() []int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		newNum, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, newNum)
	}

	return numbers
}

func makeNumberSet(numbers []int) map[int]struct{} {
	numberSet := make(map[int]struct{})
	for _, number := range numbers {
		numberSet[number] = struct{}{}
	}
	return numberSet
}
