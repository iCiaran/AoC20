package day_14

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day_14 struct{}

func New() *Day_14 {
	return &Day_14{}
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

func getValueMask(mask string) (int, int) {
	maskSet := 0
	maskClear := 0

	length := len(mask)
	pow2 := 1
	for i := 0; i < len(mask); i++ {
		if mask[length-1-i] == '0' {
			maskClear += pow2
		} else if mask[length-1-i] == '1' {
			maskSet += pow2
		}
		pow2 *= 2
	}

	return maskSet, maskClear
}

func getAddressMask(mask string) (int, []int) {
	maskSet := 0
	maskFloating := make([]int, 0)

	length := len(mask)
	pow2 := 1
	for i := 0; i < len(mask); i++ {
		if mask[length-1-i] == 'X' {
			maskFloating = append(maskFloating, pow2)
		} else if mask[length-1-i] == '1' {
			maskSet += pow2
		}
		pow2 *= 2
	}

	return maskSet, maskFloating
}

func getIndex(line string) (int, int) {
	split := strings.Split(line, " = ")
	start := 0
	end := 0
	for i, c := range split[0] {
		if c == '[' {
			start = i
		} else if c == ']' {
			end = i
		}
	}

	index, err := strconv.Atoi(split[0][start+1 : end])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	val, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return index, val
}

func applyMask(n, maskSet, maskClear int) int {
	n |= maskSet
	n &= ^maskClear
	return n
}

func getFloatingAddresses(maskFloating, base []int) []int {
	if len(maskFloating) == 0 {
		return base
	}

	newBase := make([]int, len(base)*2)

	floating := maskFloating[0]

	for i, index := range base {
		newBase[i*2] = index | floating
		newBase[i*2+1] = index & ^floating
	}

	return getFloatingAddresses(maskFloating[1:], newBase)
}

func (d *Day_14) PartA(input string) string {
	lines := getLines(input)

	memory := make(map[int]int)
	maskSet := 0
	maskClear := 0

	for _, line := range lines {
		if line[:4] == "mask" {
			maskSet, maskClear = getValueMask(line[6:])
		} else {
			index, val := getIndex(line)
			val = applyMask(val, maskSet, maskClear)
			memory[index] = val
		}
	}

	total := 0
	for _, v := range memory {
		total += v
	}

	return strconv.Itoa(total)
}

func (d *Day_14) PartB(input string) string {
	lines := getLines(input)

	memory := make(map[int]int)
	maskSet := 0
	var maskFloating []int

	for _, line := range lines {
		if line[:4] == "mask" {
			maskSet, maskFloating = getAddressMask(line[6:])
		} else {
			index, val := getIndex(line)
			index |= maskSet

			for _, floatingIndex := range getFloatingAddresses(maskFloating, []int{index}) {
				memory[floatingIndex] = val
			}
		}
	}

	total := 0
	for _, v := range memory {
		total += v
	}

	return strconv.Itoa(total)
}
