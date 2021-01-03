package day_24

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type hexCoord struct {
	q int
	r int
}

// Using the Axial coordinate system for hexagonal grids from:
// https://www.redblobgames.com/grids/hexagons/#coordinates-axial
var offsets = [...]hexCoord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, -1}, {-1, 1}}

type Day_24 struct{}

func New() *Day_24 {
	return &Day_24{}
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

func findInitialBlackTiles(input string) map[hexCoord]bool {
	flippedTiles := make(map[hexCoord]bool)

	for _, line := range getLines(input) {
		c := hexCoord{}

		for i := 0; i < len(line); i++ {
			if line[i] == 'e' {
				c.q++
			} else if line[i] == 'w' {
				c.q--
			} else if line[i] == 'n' {
				i++
				if line[i] == 'e' {
					c.q++
					c.r--
				} else if line[i] == 'w' {
					c.r--
				}
			} else if line[i] == 's' {
				i++
				if line[i] == 'e' {
					c.r++
				} else if line[i] == 'w' {
					c.q--
					c.r++
				}
			}

		}

		if flippedTiles[c] {
			delete(flippedTiles, c)
		} else {
			flippedTiles[c] = true
		}
	}

	return flippedTiles
}

func getAdjacent(c hexCoord) []hexCoord {
	adjacent := make([]hexCoord, 6)

	for i := range offsets {
		adjacent[i] = hexCoord{c.q + offsets[i].q, c.r + offsets[i].r}
	}

	return adjacent
}

func countAdjacent(c hexCoord, blackTiles map[hexCoord]bool) int {
	count := 0

	for _, a := range getAdjacent(c) {
		if blackTiles[a] {
			count++
		}
	}

	return count
}

func (d *Day_24) PartA(input string) string {
	return strconv.Itoa(len(findInitialBlackTiles(input)))
}

func (d *Day_24) PartB(input string) string {
	blackTiles := findInitialBlackTiles(input)

	for iteration := 0; iteration < 100; iteration++ {
		nextBlackTiles := make(map[hexCoord]bool)
		seen := make(map[hexCoord]bool)

		for black := range blackTiles {
			// Check which black tiles will stay alive
			count := countAdjacent(black, blackTiles)
			if count == 1 || count == 2 {
				nextBlackTiles[black] = true
			}

			// Check the tiles around each black tile
			for _, surrounding := range getAdjacent(black) {
				// If the tile is white, and we haven't already checked it
				if !blackTiles[surrounding] && !seen[surrounding] {
					if countAdjacent(surrounding, blackTiles) == 2 {
						nextBlackTiles[surrounding] = true
					}
					seen[surrounding] = true
				}
			}
		}

		blackTiles = nextBlackTiles
	}

	return strconv.Itoa(len(blackTiles))
}
