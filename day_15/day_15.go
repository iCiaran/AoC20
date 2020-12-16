package day_15

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day_15 struct{}

func New() *Day_15 {
	return &Day_15{}
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
		for _, str := range strings.Split(scanner.Text(), ",") {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			lines = append(lines, num)
		}
	}

	return lines
}

func (d *Day_15) PartA(input string) string {
	lastSeen := make(map[int]int)

	turn := 1
	lastSpoken := 0
	startingNumbers := getLines(input)

	for turn <= 2020 {
		next := 0
		if turn <= len(startingNumbers) {
			next = startingNumbers[turn-1]
		} else if t, ok := lastSeen[lastSpoken]; ok {
			next = turn - t
		}

		lastSeen[lastSpoken] = turn
		lastSpoken = next
		turn++
	}

	return strconv.Itoa(lastSpoken)
}

func (d *Day_15) PartB(input string) string {
	lastSeen := make(map[int]int)

	turn := 1
	lastSpoken := 0
	startingNumbers := getLines(input)

	for turn <= 30000000 {
		next := 0
		if turn <= len(startingNumbers) {
			next = startingNumbers[turn-1]
		} else if t, ok := lastSeen[lastSpoken]; ok {
			next = turn - t
		}

		lastSeen[lastSpoken] = turn
		lastSpoken = next
		turn++
	}

	return strconv.Itoa(lastSpoken)
}
