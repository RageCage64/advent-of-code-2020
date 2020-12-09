package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op        string
	argument  int
	callCount int
}

func newInstruction(instructionStr string) Instruction {
	instructionFields := strings.Fields(instructionStr)
	argumentInt, _ := strconv.Atoi(instructionFields[1])
	return Instruction{
		op:        instructionFields[0],
		argument:  argumentInt,
		callCount: 0,
	}
}

func (i *Instruction) reset() {
	i.callCount = 0
}

func (i *Instruction) potentialCorruption() bool {
	if i.op == "nop" || i.op == "jmp" {
		return true
	}
	return false
}

func (i *Instruction) swapOp() {
	switch i.op {
	case "nop":
		i.op = "jmp"
	case "jmp":
		i.op = "nop"
	}
}

func main() {
	program := readProgram()
	solvePart1(program)
	resetProgram(program)
	fmt.Println("----------")
	solvePart2(program)
}

func readProgram() []Instruction {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	var program []Instruction
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		instructionStr := scanner.Text()
		instruction := newInstruction(instructionStr)
		program = append(program, instruction)
	}

	return program
}

func solvePart1(program []Instruction) {
	_, accumulator := runProgram(program)
	fmt.Println(accumulator)
}

func solvePart2(program []Instruction) {
	for i := 0; i < len(program); i++ {
		instruction := &program[i]
		if instruction.potentialCorruption() {
			instruction.swapOp()
			if terminates, accumulator := runProgram(program); terminates {
				fmt.Println(accumulator)
				break
			}
			resetProgram(program)
			instruction.swapOp()
		}
	}
}

func runProgram(program []Instruction) (bool, int) {
	var programCounter, accumulator int
	terminates := true

	for programCounter < len(program) {
		instruction := &program[programCounter]
		// fmt.Println(*instruction)
		if (*instruction).callCount > 0 {
			terminates = false
			break
		}
		(*instruction).callCount++

		switch (*instruction).op {
		case "acc":
			accumulator += (*instruction).argument
			programCounter++
		case "jmp":
			programCounter += (*instruction).argument
		case "nop":
			programCounter++
		}
	}

	return terminates, accumulator
}

func resetProgram(program []Instruction) {
	for i := 0; i < len(program); i++ {
		instruction := &program[i]
		instruction.reset()
	}
}
