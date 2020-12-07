package day_07

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day_07 struct{}

type bag struct {
	name   string
	amount int
}

func New() *Day_07 {
	return &Day_07{}
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

func parseBag(input string) *bag {
	split := strings.Split(input, " ")
	amount, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	name := split[2] + split[3]

	return &bag{name, amount}
}

func buildBagsGraph(input []string) map[string][]*bag {
	bags := make(map[string][]*bag)

	for _, line := range input {
		split := strings.Split(line, "contain")
		parentFields := strings.Split(split[0], " ")
		parentName := parentFields[0] + parentFields[1]
		bags[parentName] = make([]*bag, 0)

		if split[1] == " no other bags." {
			continue
		}

		children := strings.Split(split[1], ",")

		for _, child := range children {
			bags[parentName] = append(bags[parentName], parseBag(child))
		}
	}
	return bags
}

func searchBags(start, target string, bags map[string][]*bag, memo map[string]bool) bool {
	// If we've already searched this bag return the previous result
	if v, ok := memo[start]; ok {
		return v
	}

	// If this bag has no children return false
	if len(bags[start]) == 0 {
		memo[start] = false
		return false
	}

	// Search all child bags
	for _, bag := range bags[start] {
		if bag.name == target {
			memo[start] = true
			return true
		} else if searchBags(bag.name, target, bags, memo) {
			memo[start] = true
			return true
		}
	}

	// Target not found in any children
	memo[start] = false
	return false
}

func countBags(start string, bags map[string][]*bag, memo map[string]int) int {
	// If we've already searched this bag return the previous result
	if v, ok := memo[start]; ok {
		return v
	}

	// If this bag has no children return 1
	if len(bags[start]) == 0 {
		memo[start] = 1
		return 1
	}

	total := 1
	// Count all child bags
	for _, bag := range bags[start] {
		total += bag.amount * countBags(bag.name, bags, memo)
	}

	// Return sum of children
	memo[start] = total
	return total
}

func (d *Day_07) PartA(input string) string {

	memo := make(map[string]bool)
	bags := buildBagsGraph(getLines(input))
	count := 0

	for k := range bags {
		if searchBags(k, "shinygold", bags, memo) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func (d *Day_07) PartB(input string) string {
	memo := make(map[string]int)
	bags := buildBagsGraph(getLines(input))

	// -1 because the initial shiny gold bag is included in the count
	return strconv.Itoa(countBags("shinygold", bags, memo) - 1)
}
