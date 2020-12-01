package day_01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Day_01 struct{}

func New() *Day_01 {
	return &Day_01{}
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
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		lines = append(lines, n)
	}

	return lines
}

func (d *Day_01) PartA(input string) string {
	expenses := getLines(input)
	values := make(map[int]int)
	target := 2020

	for i, n := range expenses {
		if v, ok := values[target-n]; ok {
			return strconv.Itoa(n * expenses[v])
		}
		values[n] = i
	}

	return "-1"
}

func (d *Day_01) PartB(input string) string {
	expenses := getLines(input)
	target := 2020

	sort.Ints(expenses)

	for ia := 0; ia < len(expenses)-2; ia++ {
		a := expenses[ia]

		ib := ia + 1
		ic := len(expenses) - 1

		for ib < ic {
			b := expenses[ib]
			c := expenses[ic]

			if a+b+c == target {
				return strconv.Itoa(a * b * c)
			}

			if a+b+c >= target {
				ic--
			}

			if a+b+c <= target {
				ib++
			}
		}
	}

	return "-1"
}
