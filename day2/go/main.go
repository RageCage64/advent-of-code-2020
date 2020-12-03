package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Types

type Policy struct {
	min       int
	max       int
	character string
}

type Password struct {
	policy   Policy
	password string
}

func newPassword(entry string) *Password {
	policyParts := strings.Split(entry, " ")
	occurrenceRange := strings.Split(policyParts[0], "-")
	occurCharacter := string(policyParts[1][0])
	minOccur, _ := strconv.Atoi(occurrenceRange[0])
	maxOccur, _ := strconv.Atoi(occurrenceRange[1])
	newPolicy := Policy{
		min:       minOccur,
		max:       maxOccur,
		character: occurCharacter,
	}
	password := Password{
		policy:   newPolicy,
		password: policyParts[2],
	}
	return &password
}

func (p *Password) verifyPart1Policy() bool {
	occurrences := strings.Count(p.password, p.policy.character)
	if p.policy.min <= occurrences && occurrences <= p.policy.max {
		return true
	}
	return false
}

func (p *Password) verifyPart2Policy() bool {
	var minContains bool
	var maxContains bool
	if string(p.password[p.policy.min-1]) == p.policy.character {
		minContains = true
	}
	if string(p.password[p.policy.max-1]) == p.policy.character {
		maxContains = true
	}

	if minContains != maxContains {
		return true
	}
	return false
}

type QuestionPart int

const (
	Part1 QuestionPart = iota
	Part2
)

// Main procedures

func main() {
	solve(Part1)
	solve(Part2)
}

func solve(part QuestionPart) {
	passwords := getPasswords()

	validPasswordCount := 0
	for _, password := range passwords {
		switch part {
		case Part1:
			if password.verifyPart1Policy() {
				validPasswordCount++
			}
		case Part2:
			if password.verifyPart2Policy() {
				validPasswordCount++
			}
		}
	}
	fmt.Println(validPasswordCount)
}

// Utility

func getPasswords() []*Password {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	var entries []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}

	var passwords []*Password
	for _, entry := range entries {
		newPass := newPassword(entry)
		passwords = append(passwords, newPass)
	}

	return passwords
}
