package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Types

type PassportFields map[string]string

func (p PassportFields) empty() {
	for field, _ := range p {
		p[field] = ""
	}
}

func (p PassportFields) addField(fieldStr string) {
	field := strings.Split(fieldStr, ":")
	if _, ok := p[field[0]]; ok {
		p[field[0]] = field[1]
	}
}

func (p PassportFields) validateFieldPresence() bool {
	for fieldName, value := range p {
		if value == "" && fieldName != "cid" {
			return false
		}
	}
	return true
}

func (p PassportFields) validateData() bool {
	// byr
	byrStr, ok := p["byr"]
	if !ok {
		return false
	}
	byr, err := strconv.Atoi(byrStr)
	if err != nil {
		return false
	}
	if byr < 1920 || 2002 < byr {
		return false
	}

	// iyr
	iyrStr, ok := p["iyr"]
	if !ok {
		return false
	}
	iyr, err := strconv.Atoi(iyrStr)
	if err != nil {
		return false
	}
	if iyr < 2010 || 2020 < iyr {
		return false
	}

	// eyr
	eyrStr, ok := p["eyr"]
	if !ok {
		return false
	}
	eyr, err := strconv.Atoi(eyrStr)
	if err != nil {
		return false
	}
	if eyr < 2020 || 2030 < eyr {
		return false
	}

	// hgt
	hgt, ok := p["hgt"]
	if !ok {
		return false
	}
	if strings.Contains(hgt, "cm") {
		height, err := strconv.Atoi(strings.Trim(hgt, "cm"))
		if err != nil {
			return false
		}
		if height < 150 || 193 < height {
			return false
		}
	} else if strings.Contains(hgt, "in") {
		height, err := strconv.Atoi(strings.Trim(hgt, "in"))
		if err != nil {
			return false
		}
		if height < 59 || 76 < height {
			return false
		}
	} else {
		return false
	}

	// hcl
	hcl, ok := p["hcl"]
	if !ok {
		return false
	}
	if len(hcl) != 7 || hcl[0] != '#' {
		return false
	}
	if matched, _ := regexp.Match("[0-9]|[a-f]", []byte(hcl[1:])); !matched {
		return false
	}

	// ecl
	ecl, ok := p["ecl"]
	if !ok {
		return false
	}
	eyeColours := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	if _, ok := eyeColours[ecl]; !ok {
		return false
	}

	// pid
	pid, ok := p["pid"]
	if !ok {
		return false
	}
	if len(pid) != 9 {
		return false
	}
	if matched, _ := regexp.Match("[0-9]", []byte(pid)); !matched {
		return false
	}

	// we made it through the whole disgusting function
	return true
}

type QuestionPart int

const (
	Part1 QuestionPart = iota
	Part2
)

// Procedures

func main() {
	input := readInput()
	solve(Part1, input)
	solve(Part2, input)
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

func solve(part QuestionPart, inputData []string) {
	var validPassports int
	var requiredFields PassportFields
	requiredFields = map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		"cid": "",
	}

	for _, line := range inputData {
		data := strings.Trim(line, "\n")

		if data == "" {
			var success bool

			switch part {
			case Part1:
				success = requiredFields.validateFieldPresence()
			case Part2:
				success = requiredFields.validateData()
			}

			if success {
				validPassports++
			}
			requiredFields.empty()
		} else {
			fields := strings.Split(data, " ")
			for _, field := range fields {
				requiredFields.addField(field)
			}
		}
	}

	var success bool
	switch part {
	case Part1:
		success = requiredFields.validateFieldPresence()
	case Part2:
		success = requiredFields.validateData()
	}
	if success {
		validPassports++
	}

	fmt.Println(validPassports)
}
