package main

import (
	"fmt"

	"github.com/iCiaran/AoC20/day_05"
)

func main() {
	day := day_05.New()
	fmt.Println(day.PartA("../../inputs/real_a.txt"))
	fmt.Println(day.PartB("../../inputs/real_b.txt"))
}
