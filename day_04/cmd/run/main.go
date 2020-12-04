package main

import (
	"fmt"

	"github.com/iCiaran/AoC20/day_04"
)

func main() {
	day := day_04.New()
	fmt.Println(day.PartA("../../inputs/real_a.txt"))
	fmt.Println(day.PartB("../../inputs/real_b.txt"))
}
