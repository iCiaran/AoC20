package day_22

import (
	"bufio"
	"container/list"
	"log"
	"os"
	"strconv"
)

type Day_22 struct{}

func New() *Day_22 {
	return &Day_22{}
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

func (d *Day_22) PartA(input string) string {
	a := list.New()
	b := list.New()

	readingA := true
	for _, line := range getLines(input) {
		if line == "" || line == "Player 1:" {
			continue
		}

		if line == "Player 2:" {
			readingA = false
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if readingA {
			a.PushBack(val)
		} else {
			b.PushBack(val)
		}
	}

	for a.Len() > 0 && b.Len() > 0 {
		cardA := a.Remove(a.Front()).(int)
		cardB := b.Remove(b.Front()).(int)

		if cardA > cardB {
			a.PushBack(cardA)
			a.PushBack(cardB)
		} else {
			b.PushBack(cardB)
			b.PushBack(cardA)
		}
	}

	var winningDeck *list.List
	if a.Len() > 0 {
		winningDeck = a
	} else {
		winningDeck = b
	}

	total := 0
	for i := winningDeck.Len(); i > 0; i-- {
		total += winningDeck.Remove(winningDeck.Front()).(int) * i
	}

	return strconv.Itoa(total)
}

func (d *Day_22) PartB(input string) string {
	return "B"
}
