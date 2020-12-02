package main

import (
	"fmt"

	"github.com/iCiaran/AoC20/day_02"
)

func main() {
	day := day_02.New()
	fmt.Println(day.PartA("../../inputs/real_a.txt"))
	fmt.Println(day.PartB("../../inputs/real_b.txt"))
}
