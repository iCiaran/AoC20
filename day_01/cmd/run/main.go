package main

import (
	"fmt"

	"github.com/iCiaran/AoC20/day_01"
)

func main() {
	day := day_01.New()
	fmt.Println(day.PartA("../../inputs/real_a.txt"))
	fmt.Println(day.PartB("../../inputs/real_b.txt"))
}
