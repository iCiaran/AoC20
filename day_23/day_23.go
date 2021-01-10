package day_23

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Day_23 struct{}

func New() *Day_23 {
	return &Day_23{}
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

func inNextThree(nextThree []int, destination int) bool {
	return nextThree[0] == destination || nextThree[1] == destination || nextThree[2] == destination
}

func playGame(cups map[int]int, highestCup, startingCup, iterations int) {
	currentCup := startingCup
	nextThree := make([]int, 3)

	for i := 0; i < iterations; i++ {
		// Find the 3 cups following the current cup
		nextThree[0] = cups[currentCup]
		nextThree[1] = cups[nextThree[0]]
		nextThree[2] = cups[nextThree[1]]

		destination := currentCup - 1
		if destination < 1 {
			destination = highestCup
		}

		// Keep subtracting 1 from the destination until it isn't
		// one of the 3 cups to be removed
		for inNextThree(nextThree, destination) {
			destination--
			if destination < 1 {
				destination = highestCup
			}
		}

		// Remove the 3 cups
		cups[currentCup] = cups[nextThree[2]]

		// Insert the 3 cups after the destination
		afterDestination := cups[destination]
		cups[destination] = nextThree[0]
		cups[nextThree[2]] = afterDestination

		// Move current cup clockwise
		currentCup = cups[currentCup]
	}
}

func (d *Day_23) PartA(input string) string {
	cups := make(map[int]int)
	label := []byte(getLines(input)[0])

	highestCup := 0
	for i := range label {
		if int(label[i]-48) > highestCup {
			highestCup = int(label[i] - 48)
		}
		cups[int(label[i]-48)] = int(label[(i+1)%len(label)] - 48)
	}

	playGame(cups, highestCup, int(label[0]-48), 100)

	var res string

	current := 1
	for i := 1; i < 9; i++ {
		res = res + strconv.Itoa(cups[current])
		current = cups[current]
	}

	return res
}

func (d *Day_23) PartB(input string) string {
	cups := make(map[int]int)
	label := []byte(getLines(input)[0])

	for i := 0; i < len(label)-1; i++ {
		cups[int(label[i]-48)] = int(label[i+1] - 48)
	}

	cups[int(label[len(label)-1]-48)] = 10

	for cupCount := 10; cupCount < 1_000_000; cupCount++ {
		cups[cupCount] = cupCount + 1
	}

	cups[1_000_000] = int(label[0] - 48)

	playGame(cups, 1_000_000, int(label[0]-48), 10_000_000)

	return strconv.Itoa(cups[1] * cups[cups[1]])
}
