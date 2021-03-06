package day_02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day_02 struct{}

type record struct {
	min      int
	max      int
	letter   byte
	password string
}

func New() *Day_02 {
	return &Day_02{}
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

func parseRecord(line string) *record {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == ':' || r == '-'
	})
	min, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	max, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return &record{min, max, fields[2][0], fields[3]}
}

func checkValidMinMax(r *record) bool {
	count := 0
	for _, c := range []byte(r.password) {
		if c == r.letter {
			count++
		}
		if count > r.max {
			return false
		}
	}
	return count >= r.min
}

func checkValidIndex(r *record) bool {
	firstContains := r.password[r.min-1] == r.letter
	secondContains := r.password[r.max-1] == r.letter
	return firstContains != secondContains
}

func (d *Day_02) PartA(input string) string {
	validCount := 0
	for _, line := range getLines(input) {
		r := parseRecord(line)
		if checkValidMinMax(r) {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}

func (d *Day_02) PartB(input string) string {
	validCount := 0
	for _, line := range getLines(input) {
		r := parseRecord(line)
		if checkValidIndex(r) {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}
