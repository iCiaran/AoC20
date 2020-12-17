package day_16

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day_16 struct{}

func New() *Day_16 {
	return &Day_16{}
}

type valueRange struct {
	start int
	end   int
}

type ticketField struct {
	name     string
	ranges   []valueRange
	possible []int
	index    int
}

type ticket struct {
	fields []int
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

func parseInput(input string) ([]*ticketField, ticket, []ticket) {
	ticketFields := make([]*ticketField, 0)
	myTicket := ticket{make([]int, 0)}
	nearbyTickets := make([]ticket, 0)

	lineNumber := 0
	lines := getLines(input)
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		split := strings.Split(line, ":")
		fields := strings.FieldsFunc(split[1], func(r rune) bool {
			return r == ' ' || r == ':' || r == '-'
		})

		ticketField := ticketField{}
		ticketField.name = split[0]
		ticketField.index = -1

		start, end := 0, 0
		var err error
		for i, v := range fields {
			if i%3 == 0 {
				start, err = strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			} else if i%3 == 1 {
				end, err = strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
				ticketField.ranges = append(ticketField.ranges, valueRange{start, end})
			}

		}
		ticketFields = append(ticketFields, &ticketField)
		lineNumber++
	}

	for _, val := range strings.Split(lines[lineNumber+2], ",") {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		myTicket.fields = append(myTicket.fields, num)
	}

	for i := lineNumber + 5; i < len(lines); i++ {
		splitLine := strings.Split(lines[i], ",")
		ticket := ticket{make([]int, len(splitLine))}
		for j, val := range splitLine {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			ticket.fields[j] = num
		}
		nearbyTickets = append(nearbyTickets, ticket)
	}

	return ticketFields, myTicket, nearbyTickets
}

func valueInAllRanges(ticketFields []*ticketField, value int) bool {
	for _, field := range ticketFields {
		if valueInFieldRanges(field.ranges, value) {
			return true
		}
	}
	return false
}

func valueInFieldRanges(ranges []valueRange, value int) bool {
	for _, r := range ranges {
		if value >= r.start && value <= r.end {
			return true
		}
	}
	return false
}

func ticketIsValid(ticketFields []*ticketField, t ticket) bool {
	for _, v := range t.fields {
		if !valueInAllRanges(ticketFields, v) {
			return false
		}
	}
	return true
}

func allFieldsFound(ticketFields []*ticketField) bool {
	for _, f := range ticketFields {
		if f.index == -1 {
			return false
		}
	}
	return true
}

func findValidTickets(nearbyTickets []ticket, ticketFields []*ticketField) []ticket {
	validTickets := make([]ticket, 0)

	for _, t := range nearbyTickets {
		if ticketIsValid(ticketFields, t) {
			validTickets = append(validTickets, t)
		}
	}

	return validTickets
}

func findIndices(ticketFields []*ticketField) {
	used := make(map[int]bool)

	for !allFieldsFound(ticketFields) {
		for i, f := range ticketFields {
			if f.index == -1 {
				possible := make([]int, 0)
				for _, p := range f.possible {
					if !used[p] {
						possible = append(possible, p)
					}
				}

				if len(possible) == 0 {
					panic("No possible choices")
				}

				if len(possible) == 1 {
					used[possible[0]] = true
					ticketFields[i].index = possible[0]
				}
			}
		}
	}
}

func findPossible(ticketFields []*ticketField, validTickets []ticket) {
	for _, field := range ticketFields {
		for index := range ticketFields {
			possible := true
			for _, ticket := range validTickets {
				if !valueInFieldRanges(field.ranges, ticket.fields[index]) {
					possible = false
				}
			}

			if possible {
				field.possible = append(field.possible, index)
			}
		}
	}
}

func (d *Day_16) PartA(input string) string {
	ticketFields, _, nearbyTickets := parseInput(input)

	total := 0

	for _, t := range nearbyTickets {
		for _, value := range t.fields {
			if !valueInAllRanges(ticketFields, value) {
				total += value
			}
		}
	}

	return strconv.Itoa(total)
}

func (d *Day_16) PartB(input string) string {
	ticketFields, myTicket, nearbyTickets := parseInput(input)

	validTickets := findValidTickets(nearbyTickets, ticketFields)

	findPossible(ticketFields, validTickets)

	findIndices(ticketFields)

	total := 1

	for _, f := range ticketFields {
		if strings.HasPrefix(f.name, "departure") {
			total *= myTicket.fields[f.index]
		}
	}

	return strconv.Itoa(total)
}
