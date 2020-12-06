package day_06

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Day_06 struct{}

func New() *Day_06 {
	return &Day_06{}
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

func getGroups(input []string) ([]map[rune]int, []int) {
	groups := make([]map[rune]int, 0)
	groupSize := make([]int, 0)

	current := make(map[rune]int)
	currentSize := 0

	for _, line := range input {
		if len(line) > 0 {
			currentSize++
			for _, c := range line {
				if v, ok := current[c]; ok {
					current[c] = v + 1
				} else {
					current[c] = 1

				}
			}
		} else {
			groups = append(groups, current)
			current = make(map[rune]int)
			groupSize = append(groupSize, currentSize)
			currentSize = 0
		}
	}

	groups = append(groups, current)
	groupSize = append(groupSize, currentSize)

	return groups, groupSize
}

func (d *Day_06) PartA(input string) string {
	groupSets, _ := getGroups(getLines(input))
	total := 0

	for _, g := range groupSets {
		total += len(g)
	}

	return strconv.Itoa(total)
}

func (d *Day_06) PartB(input string) string {
	groupSets, grouSizes := getGroups(getLines(input))
	total := 0

	for i, g := range groupSets {
		for _, v := range g {
			if v == grouSizes[i] {
				total++
			}
		}
	}

	return strconv.Itoa(total)
}
