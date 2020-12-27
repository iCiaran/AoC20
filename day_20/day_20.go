package day_20

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type tile struct {
	id      int
	cells   [][]byte
	borders []int
}

type tileOrientation struct {
	id       int
	rotation int
	flipped  bool
}

type tileBorders struct {
	id       int
	borders  []int
	rotation int
	flipped  bool
}

type offset struct {
	x int
	y int
}

const (
	TOP    = 0
	RIGHT  = 1
	BOTTOM = 2
	LEFT   = 3

	TILE_SIZE = 10
)

type Day_20 struct{}

func New() *Day_20 {
	return &Day_20{}
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

func parseTiles(lines []string) map[int]*tile {
	tiles := make(map[int]*tile)

	currentTile := &tile{}

	for _, line := range lines {
		if line == "" {
			tiles[currentTile.id] = currentTile
			currentTile = &tile{}
		} else if line[:4] == "Tile" {
			idString := line[5 : len(line)-1]
			id, err := strconv.Atoi(idString)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			currentTile.id = id
		} else {
			currentTile.cells = append(currentTile.cells, []byte(line))
		}
	}

	for _, v := range tiles {
		v.borders = make([]int, 4)
		for n := 0; n < TILE_SIZE; n++ {
			v.borders[TOP] *= 2
			v.borders[BOTTOM] *= 2
			v.borders[LEFT] *= 2
			v.borders[RIGHT] *= 2

			if v.cells[0][n] == '#' {
				v.borders[TOP] += 1
			}

			if v.cells[TILE_SIZE-1][n] == '#' {
				v.borders[BOTTOM] += 1
			}

			if v.cells[n][0] == '#' {
				v.borders[LEFT] += 1
			}

			if v.cells[n][TILE_SIZE-1] == '#' {
				v.borders[RIGHT] += 1
			}

		}
	}

	return tiles
}

func reverseBinary(in, bits int) int {
	res := 0
	for bits > 0 {
		res *= 2
		res |= in & 1
		in /= 2
		bits--
	}
	return res
}

// Flip cells vertically
func flipCells(cells [][]byte) [][]byte {
	temp := make([][]byte, len(cells))
	for i := 0; i < len(cells); i++ {
		temp[i] = make([]byte, len(cells))
	}

	for y := 0; y < len(cells); y++ {
		for x := 0; x < len(cells); x++ {
			temp[len(cells)-1-y][x] = cells[y][x]
		}
	}

	return temp
}

// Rotate cells 90 degrees clockwise
func rotateCells(cells [][]byte) [][]byte {
	temp := make([][]byte, len(cells))
	for i := 0; i < len(cells); i++ {
		temp[i] = make([]byte, len(cells))
	}

	n := len(cells)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			temp[j][n-1-i] = cells[i][j]
		}
	}

	return temp
}

func rotateBordersClockwise(borders []int) []int {
	rotatedBorders := make([]int, 4)

	rotatedBorders[TOP] = reverseBinary(borders[LEFT], TILE_SIZE)
	rotatedBorders[BOTTOM] = reverseBinary(borders[RIGHT], TILE_SIZE)
	rotatedBorders[RIGHT] = borders[TOP]
	rotatedBorders[LEFT] = borders[BOTTOM]

	return rotatedBorders
}

func flipBordersVertical(borders []int) []int {
	flippedBorders := make([]int, 4)

	flippedBorders[TOP] = borders[BOTTOM]
	flippedBorders[BOTTOM] = borders[TOP]
	flippedBorders[RIGHT] = reverseBinary(borders[RIGHT], TILE_SIZE)
	flippedBorders[LEFT] = reverseBinary(borders[LEFT], TILE_SIZE)

	return flippedBorders
}

func getSingleOrientation(borders []int, rotation int, flipped bool) []int {
	temp := borders

	if flipped {
		temp = flipBordersVertical(temp)
	}

	for i := 0; i < rotation; i++ {
		temp = rotateBordersClockwise(temp)
	}

	return temp
}

func getBorderOrientations(borders []int, id int) []tileBorders {
	allOrientations := make([]tileBorders, 8)

	// non flipped 4 rotations
	allOrientations[0].borders = borders
	allOrientations[0].id = id
	for i := 0; i < 3; i++ {
		allOrientations[i+1].borders = rotateBordersClockwise(allOrientations[i].borders)
		allOrientations[i+1].rotation = i + 1
		allOrientations[i+1].id = id
	}

	// flipped 4 rotations
	allOrientations[4].borders = flipBordersVertical(borders)
	allOrientations[4].id = id
	allOrientations[4].flipped = true
	for i := 0; i < 3; i++ {
		allOrientations[i+5].borders = rotateBordersClockwise(allOrientations[i+4].borders)
		allOrientations[i+5].rotation = i + 1
		allOrientations[i+5].id = id
		allOrientations[i+5].flipped = true
	}
	return allOrientations
}

func fillEdges(tileLayout [][]tileOrientation, edgeSet map[int]int, sideSet map[int]map[int]bool, tiles map[int]*tile, topLeft int, unused map[int]bool, cornerSet map[int]bool) {

	delete(unused, topLeft)
	tileLayout[0][0].id = topLeft

	next := make([]tileOrientation, 0)

	for _, b := range getBorderOrientations(tiles[topLeft].borders, topLeft) {
		if v, ok := sideSet[b.borders[RIGHT]]; !b.flipped && ok {
			for side := range v {
				if _, ok := edgeSet[side]; ok {
					if _, ok := unused[side]; ok {
						next = append(next, tileOrientation{side, b.rotation, b.flipped})
					}
				}
			}
		}
	}

	if next[0].rotation == 0 && next[1].rotation == 3 {
		tileLayout[0][0].rotation = 0
	} else if next[0].rotation == 0 && next[1].rotation == 1 {
		tileLayout[0][0].rotation = 1
	} else if next[0].rotation == 1 && next[1].rotation == 2 {
		tileLayout[0][0].rotation = 2
	} else if next[0].rotation == 2 && next[1].rotation == 3 {
		tileLayout[0][0].rotation = 3
	}

	for x := 1; x < len(tileLayout[0]); x++ {
		target := getSingleOrientation(tiles[tileLayout[0][x-1].id].borders, tileLayout[0][x-1].rotation, tileLayout[0][x-1].flipped)[RIGHT]
		possible := sideSet[target]
		for p := range possible {
			if _, ok := unused[p]; ok {
				tileLayout[0][x].id = p
			}
		}

		delete(unused, tileLayout[0][x].id)

		for _, b := range getBorderOrientations(tiles[tileLayout[0][x].id].borders, tileLayout[0][x].id) {
			if b.borders[LEFT] == target {
				tileLayout[0][x].flipped = b.flipped
				tileLayout[0][x].rotation = b.rotation
				break
			}
		}
	}

	for y := 1; y < len(tileLayout); y++ {
		target := getSingleOrientation(tiles[tileLayout[y-1][0].id].borders, tileLayout[y-1][0].rotation, tileLayout[y-1][0].flipped)[BOTTOM]
		possible := sideSet[target]
		for p := range possible {
			if _, ok := unused[p]; ok {
				tileLayout[y][0].id = p
			}
		}

		delete(unused, tileLayout[y][0].id)

		for _, b := range getBorderOrientations(tiles[tileLayout[y][0].id].borders, tileLayout[y][0].id) {
			if b.borders[TOP] == target {
				tileLayout[y][0].flipped = b.flipped
				tileLayout[y][0].rotation = b.rotation
				break
			}
		}
	}
}

func fillOthers(tileLayout [][]tileOrientation, edgeSet map[int]int, sideSet map[int]map[int]bool, tiles map[int]*tile, unused map[int]bool, cornerSet map[int]bool) {
	for y := 1; y < len(tileLayout); y++ {
		for x := 1; x < len(tileLayout[0]); x++ {
			top := getSingleOrientation(tiles[tileLayout[y-1][x].id].borders, tileLayout[y-1][x].rotation, tileLayout[y-1][x].flipped)[BOTTOM]
			left := getSingleOrientation(tiles[tileLayout[y][x-1].id].borders, tileLayout[y][x-1].rotation, tileLayout[y][x-1].flipped)[RIGHT]

			found := false
			for t := range sideSet[top] {
				for l := range sideSet[left] {
					if t == l {
						for _, b := range getBorderOrientations(tiles[t].borders, t) {
							if b.borders[TOP] == top && b.borders[LEFT] == left {
								found = true
								tileLayout[y][x].id = b.id
								tileLayout[y][x].rotation = b.rotation
								tileLayout[y][x].flipped = b.flipped
							}
						}
					}
				}
			}
			if !found {
				panic("No matching tile found")
			}
		}
	}
}

func fillGrid(tileLayout [][]tileOrientation, tiles map[int]*tile, tileWidth int) [][]byte {
	grid := make([][]byte, len(tileLayout)*tileWidth)
	for i := 0; i < len(tileLayout)*tileWidth; i++ {
		grid[i] = make([]byte, len(tileLayout[0])*tileWidth)
	}

	for y := 0; y < len(tileLayout[0]); y++ {
		for x := 0; x < len(tileLayout[y]); x++ {
			temp := tiles[tileLayout[y][x].id].cells

			if tileLayout[y][x].flipped {
				temp = flipCells(temp)
			}

			for r := 0; r < tileLayout[y][x].rotation; r++ {
				temp = rotateCells(temp)
			}

			for yy := 0; yy < tileWidth; yy++ {
				for xx := 0; xx < tileWidth; xx++ {
					grid[y*tileWidth+yy][x*tileWidth+xx] = temp[yy][xx]
				}
			}
		}
	}
	return grid
}

func removeGridOverlaps(grid [][]byte, gridDimensions, tileWidth int) [][]byte {
	newGrid := make([][]byte, (tileWidth-2)*gridDimensions)
	for i := 0; i < len(newGrid); i++ {
		newGrid[i] = make([]byte, (tileWidth-2)*gridDimensions)
	}

	nx, ny := 0, 0

	for y := 0; y < len(grid); y++ {
		if !(y%tileWidth == 0 || y%tileWidth == (tileWidth-1)) {
			nx = 0
			for x := 0; x < len(grid[0]); x++ {
				if !(x%tileWidth == 0 || x%tileWidth == (tileWidth-1)) {
					newGrid[ny][nx] = grid[y][x]
					nx++
				}
			}
			ny++
		}
	}

	return newGrid
}

func printGrid(grid [][]byte, gaps bool, tileWidth int) {
	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%c", grid[y][x])
			if gaps && (x+1)%tileWidth == 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		if gaps && (y+1)%tileWidth == 0 {
			fmt.Println()
		}
	}
}

func findMonsters(grid [][]byte) int {
	//0-                  #
	//1-#    ##    ##    ###
	//2- #  #  #  #  #  #
	//  ||||||||||||||||||||
	//  01234567890123456789
	//  00000000001111111111

	monsterOffsets := []offset{{0, 1}, {1, 2}, {4, 2}, {5, 1}, {6, 1}, {7, 2}, {10, 2}, {11, 1}, {12, 1}, {13, 2}, {16, 2}, {17, 1}, {18, 0}, {18, 1}, {19, 1}}
	maxYOffset := 2
	maxXOffset := 19

	monsterCount := 0
	i := 0
	for monsterCount == 0 && i < 8 {
		for y := 0; y+maxYOffset < len(grid); y++ {
			for x := 0; x+maxXOffset < len(grid[0]); x++ {
				found := true
				for _, o := range monsterOffsets {
					if grid[y+o.y][x+o.x] != '#' {
						found = false
						break
					}
				}

				if found {
					monsterCount++
				}
			}
		}
		i++

		// After trying all rotations switch to flipped orientation
		if i == 4 {
			grid = flipCells(grid)
		} else {
			grid = rotateCells(grid)
		}
	}

	return monsterCount
}

func findSidesEdgesCorners(tiles map[int]*tile) (map[int]map[int]bool, map[int]int, map[int]bool) {
	sideSet := make(map[int]map[int]bool)

	for _, tile := range tiles {
		borderList := getBorderOrientations(tile.borders, tile.id)

		for _, border := range borderList {

			for _, side := range border.borders {
				if _, ok := sideSet[side]; !ok {
					sideSet[side] = make(map[int]bool)
				}
				sideSet[side][border.id] = true
			}
		}
	}

	edgeSet := make(map[int]int)

	for _, v := range sideSet {
		if len(v) == 1 {
			for k := range v {
				edgeSet[k] += 1
			}
		}
	}

	cornerSet := make(map[int]bool)

	for k, v := range edgeSet {
		if v == 4 {
			cornerSet[k] = true
			delete(edgeSet, k)
		} else if v == 8 {
			delete(edgeSet, k)
		}
	}

	return sideSet, edgeSet, cornerSet
}

func (d *Day_20) PartA(input string) string {
	tiles := parseTiles(getLines(input))

	_, _, cornerSet := findSidesEdgesCorners(tiles)

	total := 1

	for corner := range cornerSet {
		total *= corner
	}

	return strconv.Itoa(total)
}

func (d *Day_20) PartB(input string) string {
	tiles := parseTiles(getLines(input))

	sideSet, edgeSet, cornerSet := findSidesEdgesCorners(tiles)

	gridDimensions := int(math.Sqrt(float64(len(tiles))))

	tileLayout := make([][]tileOrientation, gridDimensions)
	for i := 0; i < gridDimensions; i++ {
		tileLayout[i] = make([]tileOrientation, gridDimensions)
	}

	unused := make(map[int]bool)
	for _, tile := range tiles {
		unused[tile.id] = true
	}

	var firstCorner int
	for corner := range cornerSet {
		firstCorner = corner
		break
	}

	fillEdges(tileLayout, edgeSet, sideSet, tiles, firstCorner, unused, cornerSet)
	fillOthers(tileLayout, edgeSet, sideSet, tiles, unused, cornerSet)

	grid := fillGrid(tileLayout, tiles, TILE_SIZE)
	grid = removeGridOverlaps(grid, gridDimensions, TILE_SIZE)

	monsterCount := findMonsters(grid)
	filledCells := 0
	monsterSize := 15

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				filledCells++
			}
		}
	}

	roughWaters := filledCells - (monsterCount * monsterSize)

	return strconv.Itoa(roughWaters)
}
