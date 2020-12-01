package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	problemOne()

	problemTwo()
}

func problemOne() {
	numbers := getNumbers()
	numberSet := makeNumberSet(numbers)

	for _, number := range numbers {
		complement := 2020 - number
		if _, ok := numberSet[complement]; ok {
			fmt.Println(number * complement)
			return
		}
	}
}

func problemTwo() {
	numbers := getNumbers()
	numberSet := makeNumberSet(numbers)

	for i, numA := range numbers {
		complement2sum := 2020 - numA
		for _, numB := range numbers[i:] {
			complement := complement2sum - numB
			if _, ok := numberSet[complement]; ok {
				fmt.Println(numA * numB * complement)
				return
			}
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
