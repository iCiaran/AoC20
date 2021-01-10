package day_17

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type coord struct {
	x int
	y int
	z int
	w int
}

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

// func getNeighbours(c coord) []coord {
// 	neighbours := make([]coord, 0, 26)

// 	for x := -1; x < 2; x++ {
// 		for y := -1; y < 2; y++ {
// 			for z := -1; z < 2; z++ {
// 				if !(x == 0 && y == 0 && z == 0) {
// 					neighbours = append(neighbours, coord{c.x + x, c.y + y, c.z + z, 0})
// 				}
// 			}
// 		}
// 	}

// 	return neighbours
// }

func getNeighbours(c coord, hyper bool) []coord {
	neighbours := make([]coord, 0, 26)

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			for z := -1; z < 2; z++ {
				if hyper {
					for w := -1; w < 2; w++ {
						if !(x == 0 && y == 0 && z == 0 && w == 0) {
							neighbours = append(neighbours, coord{c.x + x, c.y + y, c.z + z, c.w + w})
						}
					}
				} else {
					if !(x == 0 && y == 0 && z == 0) {
						neighbours = append(neighbours, coord{c.x + x, c.y + y, c.z + z, 0})
					}
				}
			}
		}
	}

	return neighbours
}

func countActive(coords []coord, active map[coord]bool) int {
	count := 0

	for _, c := range coords {
		if active[c] {
			count++
		}
	}

	return count
}

func cycle(active map[coord]bool, hyper bool) map[coord]bool {
	nextActive := make(map[coord]bool)

	for current := range active {
		neighbours := getNeighbours(current, hyper)

		activeNeighbours := countActive(neighbours, active)
		if activeNeighbours == 2 || activeNeighbours == 3 {
			nextActive[current] = true
		}

		for _, neighbour := range neighbours {
			if !active[neighbour] && countActive(getNeighbours(neighbour, hyper), active) == 3 {
				nextActive[neighbour] = true
			}
		}
	}

	return nextActive
}

func runBootProcess(input string, hyper bool) map[coord]bool {
	active := make(map[coord]bool)

	for y, line := range getLines(input) {
		for x, c := range line {
			if c == '#' {
				active[coord{x, y, 0, 0}] = true
			}
		}
	}

	for i := 0; i < 6; i++ {
		active = cycle(active, hyper)
	}

	return active
}

func (d *Day_17) PartA(input string) string {
	return strconv.Itoa(len(runBootProcess(input, false)))
}

func (d *Day_17) PartB(input string) string {
	return strconv.Itoa(len(runBootProcess(input, true)))
}
