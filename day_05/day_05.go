package day_05

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

type Day_05 struct{}

func New() *Day_05 {
	return &Day_05{}
}

func getLines(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, decode(scanner.Text()))
	}

	return lines
}

func decode(line string) int {
	total := 0

	for _, c := range line {
		total = total << 1
		if c == 'B' || c == 'R' {
			total += 1
		}
	}

	return total
}

func getHighestSeatId(input []int) int {
	max := 0
	for _, line := range input {

		if line > max {
			max = line
		}
	}
	return max
}

func getMissingSeat(input []int) int {
	max := 0
	min := math.MaxInt64
	total := 0

	for _, line := range input {

		if line > max {
			max = line
		}

		if line < min {
			min = line
		}

		total += line
	}

	sum := (max - min + 1) * (min + max) / 2
	return sum - total
}

func (d *Day_05) PartA(input string) string {
	return strconv.Itoa(getHighestSeatId(getLines(input)))
}

func (d *Day_05) PartB(input string) string {
	return strconv.Itoa(getMissingSeat(getLines(input)))
}
