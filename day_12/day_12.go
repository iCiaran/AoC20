package day_12

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Day_12 struct{}

type instruction struct {
	action byte
	amount int
}

type position struct {
	x int
	y int
}

func New() *Day_12 {
	return &Day_12{}
}

func getInstructions(filepath string) []instruction {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var instructions []instruction

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		action := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		instructions = append(instructions, instruction{action, amount})
	}

	return instructions
}

func nextDirection(current, change int) int {
	next := (current + change/90) % 4
	if next < 0 {
		next += 4
	}
	return next
}

func getDelta(direction, amount int) position {
	delta := position{0, 0}
	switch direction {
	case 0: // North
		delta.y += amount
	case 1: // East
		delta.x += amount
	case 2: // South
		delta.y -= amount
	case 3: // West
		delta.x -= amount
	}
	return delta
}

func manhattanDistance(start, end position) int {
	x := start.x - end.x
	if x < 0 {
		x *= -1
	}
	y := start.y - end.y
	if y < 0 {
		y *= -1
	}

	return x + y
}

func rotateWaypoint(waypointPos position, amount int) position {
	if amount < 0 {
		amount += 360
	}

	for i := 0; i < amount/90; i++ {
		ty := -1 * waypointPos.x
		tx := waypointPos.y

		waypointPos.x = tx
		waypointPos.y = ty
	}

	return position{waypointPos.x, waypointPos.y}
}

func (d *Day_12) PartA(input string) string {
	instructions := getInstructions(input)
	dir := 1
	pos := position{0, 0}

	for _, i := range instructions {
		switch i.action {
		case 'N':
			pos.y += i.amount
		case 'S':
			pos.y -= i.amount
		case 'E':
			pos.x += i.amount
		case 'W':
			pos.y -= i.amount
		case 'L':
			dir = nextDirection(dir, i.amount*-1)
		case 'R':
			dir = nextDirection(dir, i.amount)
		case 'F':
			delta := getDelta(dir, i.amount)
			pos.x += delta.x
			pos.y += delta.y
		}
	}

	return strconv.Itoa(manhattanDistance(position{0, 0}, pos))
}

func (d *Day_12) PartB(input string) string {
	instructions := getInstructions(input)
	shipPos := position{0, 0}
	waypointPos := position{10, 1}

	for _, i := range instructions {
		switch i.action {
		case 'N':
			waypointPos.y += i.amount
		case 'S':
			waypointPos.y -= i.amount
		case 'E':
			waypointPos.x += i.amount
		case 'W':
			waypointPos.x -= i.amount
		case 'L':
			waypointPos = rotateWaypoint(waypointPos, -1*i.amount)
		case 'R':
			waypointPos = rotateWaypoint(waypointPos, i.amount)
		case 'F':
			shipPos.x += waypointPos.x * i.amount
			shipPos.y += waypointPos.y * i.amount
		}
	}

	return strconv.Itoa(manhattanDistance(position{0, 0}, shipPos))
}
