package day_08

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day_08 struct{}

type instruction struct {
	opcode   string
	operand  int
	executed bool
}

func New() *Day_08 {
	return &Day_08{}
}

func getLines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func loadInstructions(input string) []instruction {
	lines := getLines(input)

	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		split := strings.Split(line, " ")
		opcode := split[0]
		operand, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		instructions[i] = instruction{opcode, operand, false}
	}

	return instructions
}

func (d *Day_08) PartA(input string) string {
	instructions := loadInstructions(input)

	acc := 0
	pc := 0

	for !instructions[pc].executed {
		instructions[pc].executed = true

		switch instructions[pc].opcode {
		case "jmp":
			pc += instructions[pc].operand
		case "acc":
			acc += instructions[pc].operand
			pc += 1
		case "nop":
			pc += 1
		}
	}

	return strconv.Itoa(acc)
}

func (d *Day_08) PartB(input string) string {
	instructions := loadInstructions(input)

	var acc int

	for i := 0; i < len(instructions); i++ {
		// Flip a single jmp/nop
		if instructions[i].opcode == "jmp" {
			instructions[i].opcode = "nop"
		} else if instructions[i].opcode == "nop" {
			instructions[i].opcode = "jmp"
		} else {
			continue
		}

		// Reset
		acc = 0
		for j := 0; j < len(instructions); j++ {
			instructions[j].executed = false
		}

		pc := 0
		for pc < len(instructions) && !instructions[pc].executed {
			instructions[pc].executed = true

			switch instructions[pc].opcode {
			case "jmp":
				pc += instructions[pc].operand
			case "acc":
				acc += instructions[pc].operand
				pc += 1
			case "nop":
				pc += 1
			}
		}

		if pc >= len(instructions) {
			break
		}

		// Return jmp/nop to original state
		if instructions[i].opcode == "jmp" {
			instructions[i].opcode = "nop"
		} else if instructions[i].opcode == "nop" {
			instructions[i].opcode = "jmp"
		}
	}

	return strconv.Itoa(acc)
}
