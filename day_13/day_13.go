package day_13

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Day_13 struct{}

func New() *Day_13 {
	return &Day_13{}
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

var one = big.NewInt(1)

// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func (d *Day_13) PartA(input string) string {
	lines := getLines(input)

	startTime, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	busIds := make([]int, 0)

	for _, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			continue
		}

		id, err := strconv.Atoi(bus)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		busIds = append(busIds, id)
	}

	minWait := math.MaxInt64
	minId := math.MaxInt64

	for _, id := range busIds {
		wait := id - (startTime % id)
		if wait < minWait {
			minWait = wait
			minId = id
		}
	}

	return strconv.Itoa(minWait * minId)
}

func (d *Day_13) PartB(input string) string {
	lines := getLines(input)

	a := make([]*big.Int, 0)
	n := make([]*big.Int, 0)

	for i, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			continue
		}

		id, err := strconv.Atoi(bus)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		a = append(a, big.NewInt(int64(-1*i)))
		n = append(n, big.NewInt(int64(id)))
	}

	res, _ := crt(a, n)
	return res.String()
}
