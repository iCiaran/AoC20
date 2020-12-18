package day_17

import (
	"bufio"
	"log"
	"os"
)

type Day_17 struct{}

func New() *Day_17 {
	return &Day_17{}
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

func (d *Day_17) PartA(input string) string {
    return "A"
}

func (d *Day_17) PartB(input string) string {
	return "B"
}
