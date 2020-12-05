package main

import (
	"fmt"
	"time"

	"github.com/iCiaran/AoC20/day_01"
	"github.com/iCiaran/AoC20/day_02"
	"github.com/iCiaran/AoC20/day_03"
	"github.com/iCiaran/AoC20/day_04"
	"github.com/iCiaran/AoC20/day_05"
)

type challenge interface {
	PartA(input string) string
	PartB(input string) string
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}

func inputPath(day int) string {
	return fmt.Sprintf("day_%02d/inputs/", day)
}

func runSingle(challenges map[int]challenge, day int, a, b bool) {
	fmt.Printf("Day %d\n", day)
	inputPath := inputPath(day)
	if c, ok := challenges[day]; ok {
		if a {
			start := time.Now()
			result := c.PartA(inputPath + "real_a.txt")
			elapsed := time.Since(start)
			fmt.Printf("├ A - %v - %s\n", elapsed, result)
		} else {
			fmt.Println("├ A - Skipped")
		}
		if b {
			start := time.Now()
			result := c.PartB(inputPath + "real_b.txt")
			elapsed := time.Since(start)
			fmt.Printf("└ B - %v - %s\n", elapsed, result)
		} else {
			fmt.Println("└ B - Skipped")
		}
	} else {
		fmt.Println("└ Skipped")
	}
}

func main() {
	challenges := make_challenges()

	for i := 1; i < 26; i++ {
		runSingle(challenges, i, true, true)
	}
}

func make_challenges() map[int]challenge {
	challenges := make(map[int]challenge, 25)

	challenges[1] = day_01.New()
	challenges[2] = day_02.New()
	challenges[3] = day_03.New()
	challenges[4] = day_04.New()
	challenges[5] = day_05.New()

	return challenges
}
