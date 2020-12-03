package day_03

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Day_03 struct{}

func New() *Day_03 {
	return &Day_03{}
}

func getLines(filepath string) [][]byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var lines [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	return lines
}

func (d *Day_03) PartA(input string) string {
	treeMap := getLines(input)
	width := len(treeMap[0])

	x := 0
	treesHit := 0

	for y := 0; y < len(treeMap); y++ {
		if treeMap[y][x] == '#' {
			treesHit++
		}
		x = (x + 3) % width
	}

	return strconv.Itoa(treesHit)
}

func (d *Day_03) PartB(input string) string {
	treeMap := getLines(input)
	width := len(treeMap[0])

	xSteps := []int{1, 3, 5, 7, 1}
	xPos := []int{0, 0, 0, 0, 0}
	treesHit := []int{0, 0, 0, 0, 0}

	for y := 0; y < len(treeMap); y++ {
		for i := 0; i < 4; i++ {
			if treeMap[y][xPos[i]] == '#' {
				treesHit[i]++
			}
			xPos[i] = (xPos[i] + xSteps[i]) % width
		}

		if y%2 == 0 {
			if treeMap[y][xPos[4]] == '#' {
				treesHit[4]++
			}
			xPos[4] = (xPos[4] + xSteps[4]) % width
		}

	}

	mulTreesHit := treesHit[0] * treesHit[1] * treesHit[2] * treesHit[3] * treesHit[4]

	return strconv.Itoa(mulTreesHit)
}
