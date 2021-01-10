package day_25

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Day_25 struct{}

func New() *Day_25 {
	return &Day_25{}
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

func findLoopSize(publicKey int) int {
	value := 1
	loops := 0

	for value != publicKey {
		value = value * 7
		value = value % 20201227
		loops++
	}
	return loops
}

func transform(key, loop int) int {
	encryption := 1
	for i := 0; i < loop; i++ {
		encryption = encryption * key
		encryption = encryption % 20201227
	}
	return encryption
}

func (d *Day_25) PartA(input string) string {
	lines := getLines(input)
	cardPublicKey, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	doorPublicKey, err := strconv.Atoi(lines[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cardLoopSize := findLoopSize(cardPublicKey)
	encryption := transform(doorPublicKey, cardLoopSize)

	return strconv.Itoa(encryption)
}

func (d *Day_25) PartB(input string) string {
	return "Merry Christmas!"
}
