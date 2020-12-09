package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type BagContainRule struct {
// 	bag   string
// 	count int
// }

func main() {
	input := readInput()
	bagGraph := buildBagGraph(input)
	solvePart1(bagGraph)
	solvePart2(bagGraph)
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

func buildBagGraph(input []string) map[string]map[string]int {
	bagGraph := make(map[string]map[string]int)
	for _, line := range input {
		splitLine := strings.Fields(line)
		bag := splitLine[0] + " " + splitLine[1]
		bagGraph[bag] = getRules(splitLine)
	}
	return bagGraph
}

func solvePart1(bagGraph map[string]map[string]int) {
	checkedBags := make(map[string]bool)
	for _, bag := range allBagsThatCanDirectlyContain(bagGraph, "shiny gold") {
		checkedBags[bag] = false
	}

	for hasUncheckedBags(checkedBags) {
		for bag, _ := range checkedBags {
			newBags := allBagsThatCanDirectlyContain(bagGraph, bag)
			for _, newBag := range newBags {
				if _, ok := checkedBags[newBag]; !ok {
					checkedBags[newBag] = false
				}
			}
			checkedBags[bag] = true
		}
	}
	fmt.Println(len(checkedBags))
}

func getRules(splitLine []string) map[string]int {
	bagCountStr := splitLine[4]
	if bagCountStr == "no" {
		return nil
	}

	containBags := make(map[string]int)
	var bagType string
	var bagCount int
	containRulesStr := strings.Join(splitLine[4:], " ")
	containRules := strings.Split(containRulesStr, ",")
	for _, rule := range containRules {
		fields := strings.Fields(rule)
		bagCount, _ = strconv.Atoi(fields[0])
		bagType = fields[1] + " " + fields[2]
		containBags[bagType] = bagCount
	}
	return containBags
}

func allBagsThatCanDirectlyContain(
	bags map[string]map[string]int,
	checkBag string,
) []string {
	var containBags []string
	for bag, rules := range bags {
		if rules != nil {
			if _, ok := rules[checkBag]; ok {
				containBags = append(containBags, bag)
			}
		}
	}
	return containBags
}

func hasUncheckedBags(bags map[string]bool) bool {
	for _, checked := range bags {
		if !checked {
			return true
		}
	}
	return false
}

func solvePart2(bagGraph map[string]map[string]int) {
	fmt.Println(countContainedBags(bagGraph, "shiny gold"))
}

func countContainedBags(bagGraph map[string]map[string]int, checkBag string) int {
	var countBags int

	bagRules := bagGraph[checkBag]
	if bagRules == nil {
		return 0
	}

	for bag, count := range bagRules {
		countBags += count + (count * countContainedBags(bagGraph, bag))
	}

	return countBags
}
