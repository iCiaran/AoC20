package day_11

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Day_11 struct{}

type status int

const (
	FLOOR status = iota
	EMPTY
	OCCUPIED
)

func New() *Day_11 {
	return &Day_11{}
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

func to1D(x, y, width int) int {
	return y*width + x
}

func isInBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func occupiedNeighbours(grid []status, x, y, width, height int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				newX := x + i
				newY := y + j

				if isInBounds(newX, newY, width, height) {
					if grid[to1D(newX, newY, width)] == OCCUPIED {
						count++
					}
				}
			}
		}
	}
	return count
}

func checkSeats(grid []status, x, y, width, height, dx, dy int) int {
	newX := x + dx
	newY := y + dy
	for isInBounds(newX, newY, width, height) {
		if grid[to1D(newX, newY, width)] == OCCUPIED {
			return 1
		} else if grid[to1D(newX, newY, width)] == EMPTY {
			return 0
		}

		newX += dx
		newY += dy
	}
	return 0
}

func visibleOccupiedNeighbours(grid []status, x, y, width, height int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				count += checkSeats(grid, x, y, width, height, i, j)
			}
		}
	}
	return count
}

func readGrid(lines []string, width, height int) [][]status {
	grid := make([][]status, 2)

	grid[0] = make([]status, height*width)
	grid[1] = make([]status, height*width)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] == '.' {
				grid[0][to1D(x, y, width)] = FLOOR
			} else if lines[y][x] == 'L' {
				grid[0][to1D(x, y, width)] = EMPTY
			}
		}
	}
	return grid
}

func countOccupied(grid [][]status, height int, width int, iteration int) int {
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[iteration%2][to1D(x, y, width)] == OCCUPIED {
				count++
			}
		}
	}
	return count
}

func (d *Day_11) PartA(input string) string {
	lines := getLines(input)
	width := len(lines[0])
	height := len(lines)

	grid := readGrid(lines, width, height)
	iteration := 0
	hasChange := true

	for hasChange {
		hasChange = false

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				coord := to1D(x, y, width)
				if grid[iteration%2][coord] == EMPTY {
					if occupiedNeighbours(grid[iteration%2], x, y, width, height) == 0 {
						hasChange = true
						grid[(iteration+1)%2][coord] = OCCUPIED
					} else {
						grid[(iteration+1)%2][coord] = grid[iteration%2][coord]
					}
				} else if grid[iteration%2][coord] == OCCUPIED {
					if occupiedNeighbours(grid[iteration%2], x, y, width, height) >= 4 {
						hasChange = true
						grid[(iteration+1)%2][coord] = EMPTY
					} else {
						grid[(iteration+1)%2][coord] = grid[iteration%2][coord]
					}
				}
			}
		}

		iteration++
	}

	count := countOccupied(grid, height, width, iteration)

	return strconv.Itoa(count)
}

func (d *Day_11) PartB(input string) string {
	lines := getLines(input)
	width := len(lines[0])
	height := len(lines)

	grid := readGrid(lines, width, height)
	iteration := 0
	hasChange := true

	for hasChange {
		hasChange = false

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				coord := to1D(x, y, width)
				if grid[iteration%2][coord] == EMPTY {
					if visibleOccupiedNeighbours(grid[iteration%2], x, y, width, height) == 0 {
						hasChange = true
						grid[(iteration+1)%2][coord] = OCCUPIED
					} else {
						grid[(iteration+1)%2][coord] = grid[iteration%2][coord]
					}
				} else if grid[iteration%2][coord] == OCCUPIED {
					if visibleOccupiedNeighbours(grid[iteration%2], x, y, width, height) >= 5 {
						hasChange = true
						grid[(iteration+1)%2][coord] = EMPTY
					} else {
						grid[(iteration+1)%2][coord] = grid[iteration%2][coord]
					}
				}
			}
		}

		iteration++
	}

	count := countOccupied(grid, height, width, iteration)

	return strconv.Itoa(count)
}
