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

func getHighestSeatId(input []string) int {
	max := 0
	for _, line := range input {
		total := decode(line)

		if total > max {
			max = total
		}
	}
	return max
}

func getTakenSeats(input []string) (map[int]bool, int) {
	max := 0
	seats := make(map[int]bool)
	for _, line := range input {

		total := decode(line)

		if total > max {
			max = total
		}

		seats[total] = true
	}

	return seats, max
}

func getMissingSeat(input []string) int {
	max := 0
	min := math.MaxInt64
	total := 0

	for _, line := range input {

		ref := decode(line)

		if ref > max {
			max = ref
		}

		if ref < min {
			min = ref
		}

		total += ref
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
