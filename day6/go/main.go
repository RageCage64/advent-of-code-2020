package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readInput()
	countUniqueAnswers(input)
	countUnanimousAnswers(input)
}

func readInput() []string {
	// file, _ := os.Open("../input_fake.txt")
	file, _ := os.Open("../input.txt")
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func countUniqueAnswers(input []string) {
	answers := make(map[rune]struct{})
	groupAnswerCounts := make([]int, 0)
	for _, line := range input {
		if line == "" {
			groupAnswerCounts = append(groupAnswerCounts, len(answers))
			answers = make(map[rune]struct{})
		}

		for _, answer := range line {
			answers[answer] = struct{}{}
		}
	}

	var sum int
	for _, answers := range groupAnswerCounts {
		sum += answers
	}
	fmt.Println(sum)
}

func countUnanimousAnswers(input []string) {
	answers := make(map[rune]struct{})
	groupAnswerCounts := make([]int, 0)
	first := true
	for _, line := range input {
		if line == "" {
			groupAnswerCounts = append(groupAnswerCounts, len(answers))
			answers = make(map[rune]struct{})
			first = true
		} else {
			if first {
				for _, answer := range line {
					answers[answer] = struct{}{}
				}
				first = false
			} else {
				for answer, _ := range answers {
					if !strings.Contains(line, string(answer)) {
						delete(answers, answer)
					}
				}
			}
		}
	}

	var sum int
	for _, answers := range groupAnswerCounts {
		sum += answers
	}
	fmt.Println(sum)
}
