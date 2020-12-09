package day_09

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Day_09 struct {
	preamble int
}

func New() *Day_09 {
	return &Day_09{25}
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
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		lines = append(lines, val)
	}

	return lines
}

func isNextValid(next, index, distance int, lastSeen map[int]int, recent []int) bool {
	for _, n := range recent {
		if v, ok := lastSeen[next-n]; ok && index-v < distance {
			return true
		}
	}

	return false
}

func findInvalid(values []int, preamble int) int {
	lastSeen := make(map[int]int)
	recent := make([]int, preamble)

	for i := 0; i < preamble; i++ {
		recent[i] = values[i]
		lastSeen[values[i]] = i
	}

	for i := preamble; i < len(values); i++ {
		if !isNextValid(values[i], i, preamble, lastSeen, recent) {
			return values[i]
		}
		recent = append(recent[1:], values[i])
		lastSeen[values[i]] = i
	}

	return -1
}

func findContiguousSum(values []int, target int) (int, int) {
	currentSum := values[0]
	startIndex := 0
	endIndex := 1
	for endIndex < len(values) {
		// If current sum > target remove values from start of range
		for currentSum > target && startIndex < endIndex-1 {
			currentSum -= values[startIndex]
			startIndex++
		}

		// Target found
		if currentSum == target {
			return startIndex, endIndex - 1
		}

		// Add values from end of range if we can
		if endIndex < len(values) {
			currentSum += values[endIndex]
		}

		endIndex++
	}
	return -1, -1
}

func findMinMax(values []int) (int, int) {
	lowest := values[0]
	highest := values[0]

	for _, n := range values {
		if n < lowest {
			lowest = n
		}
		if n > highest {
			highest = n
		}
	}

	return highest, lowest
}

func (d *Day_09) PartA(input string) string {
	values := getLines(input)
	invalid := findInvalid(values, d.preamble)
	return strconv.Itoa(invalid)
}

func (d *Day_09) PartB(input string) string {
	values := getLines(input)
	target := findInvalid(values, d.preamble)
	start, end := findContiguousSum(values, target)
	highest, lowest := findMinMax(values[start:end])
	return strconv.Itoa(highest + lowest)
}
