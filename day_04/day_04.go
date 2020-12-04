package day_04

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day_04 struct{}
type passport map[string]string

func New() *Day_04 {
	return &Day_04{}
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

func getPassportList(input []string) []passport {
	passports := make([]passport, 0)
	p := make(passport)

	for _, line := range input {
		if len(line) > 0 {
			split := strings.Split(line, " ")
			for _, pair := range split {
				keyValue := strings.Split(pair, ":")
				p[keyValue[0]] = keyValue[1]
			}
		} else {
			passports = append(passports, p)
			p = make(passport)
		}
	}

	passports = append(passports, p)

	return passports
}

func validatePassportValues(p passport) bool {
	return validateYear(p["byr"], 1920, 2002) &&
		validateYear(p["iyr"], 2010, 2020) &&
		validateYear(p["eyr"], 2020, 2030) &&
		validateHeight(p["hgt"]) &&
		validateHairColour(p["hcl"]) &&
		validateEyeColour(p["ecl"]) &&
		validatePassportId(p["pid"])
}

func validateYear(input string, min, max int) bool {
	if len(input) != 4 {
		return false
	}

	year, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return year >= min && year <= max
}

func validateHeight(input string) bool {
	if len(input) < 2 {
		return false
	}

	unit := input[len(input)-2:]
	if unit != "cm" && unit != "in" {
		return false
	}

	height, err := strconv.Atoi(input[:len(input)-2])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if unit == "cm" && height >= 150 && height <= 193 {
		return true
	}

	if unit == "in" && height >= 59 && height <= 76 {
		return true
	}

	return false
}

func validateHairColour(input string) bool {
	if len(input) != 7 {
		return false
	}

	if input[0] != '#' {
		return false
	}

	for _, char := range []byte(input[1:]) {
		if !(char >= 48 && char <= 57 || char >= 97 && char <= 102) {
			return false
		}
	}

	return true
}

func validateEyeColour(input string) bool {
	return input == "amb" ||
		input == "blu" ||
		input == "brn" ||
		input == "gry" ||
		input == "grn" ||
		input == "hzl" ||
		input == "oth"
}

func validatePassportId(input string) bool {
	if len(input) != 9 {
		return false
	}

	for _, char := range []byte(input) {
		if !(char >= 48 && char <= 57) {
			return false
		}
	}

	return true
}

func (d *Day_04) PartA(input string) string {
	validCount := 0
	passports := getPassportList(getLines(input))

	for _, p := range passports {
		if len(p) == 8 {
			validCount++
		} else if len(p) == 7 {
			if _, ok := p["cid"]; !ok {
				validCount++
			}
		}
	}

	return strconv.Itoa(validCount)
}

func (d *Day_04) PartB(input string) string {
	validCount := 0
	passports := getPassportList(getLines(input))

	for _, p := range passports {
		if validatePassportValues(p) {
			validCount++
		}
	}

	return strconv.Itoa(validCount)
}
