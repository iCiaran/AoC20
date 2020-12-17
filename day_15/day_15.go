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

func playGame(turns int, startingNumbers []int) int {
	lastSeen := make([]int, turns)

	turn := 1
	lastSpoken := 0

	for turn <= turns {
		next := 0
		if turn <= len(startingNumbers) {
			next = startingNumbers[turn-1]
		} else if lastSeen[lastSpoken] > 0 {
			next = turn - lastSeen[lastSpoken]
		}

		lastSeen[lastSpoken] = turn
		lastSpoken = next
		turn++
	}

	return lastSpoken
}

func (d *Day_15) PartA(input string) string {
	startingNumbers := getLines(input)

	return strconv.Itoa(playGame(2020, startingNumbers))
}

func (d *Day_15) PartB(input string) string {
	startingNumbers := getLines(input)

	return strconv.Itoa(playGame(30000000, startingNumbers))
}
