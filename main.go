package main

import (
	"fmt"
	"time"

	"github.com/iCiaran/AoC20/day_01"
	"github.com/iCiaran/AoC20/day_02"
	"github.com/iCiaran/AoC20/day_03"
	"github.com/iCiaran/AoC20/day_04"
	"github.com/iCiaran/AoC20/day_05"
	"github.com/iCiaran/AoC20/day_06"
	"github.com/iCiaran/AoC20/day_07"
	"github.com/iCiaran/AoC20/day_08"
	"github.com/iCiaran/AoC20/day_09"
	"github.com/iCiaran/AoC20/day_10"
	"github.com/iCiaran/AoC20/day_11"
	"github.com/iCiaran/AoC20/day_12"
	"github.com/iCiaran/AoC20/day_13"
	"github.com/iCiaran/AoC20/day_14"
	"github.com/iCiaran/AoC20/day_15"
	"github.com/iCiaran/AoC20/day_16"
	"github.com/iCiaran/AoC20/day_17"
	"github.com/iCiaran/AoC20/day_18"
	"github.com/iCiaran/AoC20/day_20"
	"github.com/iCiaran/AoC20/day_21"
	"github.com/iCiaran/AoC20/day_23"
	"github.com/iCiaran/AoC20/day_24"
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
	challenges[6] = day_06.New()
	challenges[7] = day_07.New()
	challenges[8] = day_08.New()
	challenges[9] = day_09.New()
	challenges[10] = day_10.New()
	challenges[11] = day_11.New()
	challenges[12] = day_12.New()
	challenges[13] = day_13.New()
	challenges[14] = day_14.New()
	challenges[15] = day_15.New()
	challenges[16] = day_16.New()
	challenges[17] = day_17.New()
	challenges[18] = day_18.New()
	challenges[20] = day_20.New()
	challenges[21] = day_21.New()
	challenges[23] = day_23.New()
	challenges[24] = day_24.New()

	return challenges
}
