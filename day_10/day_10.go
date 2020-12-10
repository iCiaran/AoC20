package day_10

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type Day_10 struct{}

func New() *Day_10 {
	return &Day_10{}
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

func (d *Day_10) PartA(input string) string {
	joltages := getLines(input)
	sort.Ints(joltages)

	// Map of gap between 2 joltages -> number of times it occurs
	gaps := make(map[int]int)

	// Always have to make the jump from source to first joltage
	gaps[joltages[0]] = 1

	// Last jump is always 3
	gaps[3] = 1

	// Include every adapter so always just take the next joltage
	// (can't go backwards)
	for i := 1; i < len(joltages); i++ {
		gap := joltages[i] - joltages[i-1]
		if v, ok := gaps[gap]; ok {
			gaps[gap] = v + 1
		} else {
			gaps[gap] = 1
		}
	}

	return strconv.Itoa(gaps[1] * gaps[3])
}

func (d *Day_10) PartB(input string) string {
	joltages := getLines(input)
	sort.Ints(joltages)

	// Map of joltage -> number of paths to reach it
	combinations := make(map[int]int)
	// Starts only containing the source
	combinations[0] = 1

	for _, joltage := range joltages {
		// Possible number of ways to get to each joltage is the sum
		// of the ways to get to joltage - 1, joltage - 2 and joltage -3
		possible := 0
		for i := 1; i <= 3; i++ {
			if v, ok := combinations[joltage-i]; ok {
				possible += v
			}
		}
		combinations[joltage] = possible
	}

	// Return the number of paths to the final joltage
	lastJoltage := joltages[len(joltages)-1]
	return strconv.Itoa(combinations[lastJoltage])
}
