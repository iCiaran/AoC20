package main

import (
	"fmt"
	"time"
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
	challenges := make(map[int]challenge, 25)

	runSingle(challenges, 0, true, true)
	for i := 1; i < 26; i++ {
		runSingle(challenges, i, true, true)
	}
}
