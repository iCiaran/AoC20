package day_21

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day_21 struct{}

func New() *Day_21 {
	return &Day_21{}
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

func setIntersect(a, b map[string]bool) map[string]bool {
	ret := make(map[string]bool)

	for k := range a {
		if b[k] {
			ret[k] = true
		}
	}

	return ret
}

func setMinus(a, b map[string]bool) map[string]bool {
	ret := make(map[string]bool)

	for k := range a {
		if !b[k] {
			ret[k] = true
		}
	}

	return ret
}

func parseInput(input string) ([]map[string]bool, []map[string]bool, map[string]bool, map[string]bool) {
	ingredients := make([]map[string]bool, 0)
	allergens := make([]map[string]bool, 0)

	allAllergens := make(map[string]bool, 0)
	allIngredients := make(map[string]bool, 0)

	for _, line := range getLines(input) {
		i := make(map[string]bool)
		a := make(map[string]bool)

		split := strings.Split(line, " (contains ")

		for _, s := range strings.Split(split[0], " ") {
			i[s] = true
			allIngredients[s] = true
		}

		for _, s := range strings.Split(split[1][:len(split[1])-1], ", ") {
			a[s] = true
			allAllergens[s] = true
		}

		ingredients = append(ingredients, i)
		allergens = append(allergens, a)

	}

	return ingredients, allergens, allIngredients, allAllergens
}

func findPossible(ingredients, allergens []map[string]bool, allIngredients, allAllergens map[string]bool) map[string]bool {
	possible := make(map[string]bool)

	for allergen := range allAllergens {
		temp := allIngredients

		for i := range allergens {
			if allergens[i][allergen] {
				temp = setIntersect(temp, ingredients[i])
			}
		}
		for i := range temp {
			possible[i] = true
		}
	}

	return possible
}

func (d *Day_21) PartA(input string) string {

	ingredients, allergens, allIngredients, allAllergens := parseInput(input)

	possible := findPossible(ingredients, allergens, allIngredients, allAllergens)

	notPossible := setMinus(allIngredients, possible)

	total := 0

	for _, r := range ingredients {
		for i := range r {
			if notPossible[i] {
				total++
			}
		}
	}
	return strconv.Itoa(total)
}

func (d *Day_21) PartB(input string) string {
	ingredients, allergens, allIngredients, allAllergens := parseInput(input)

	possible := make(map[string]map[string]bool)

	for allergen := range allAllergens {
		temp := allIngredients

		for i := range allergens {
			if allergens[i][allergen] {
				temp = setIntersect(temp, ingredients[i])
			}
		}
		for ingredient := range temp {
			if _, ok := possible[allergen]; !ok {
				possible[allergen] = make(map[string]bool)
			}
			possible[allergen][ingredient] = true
		}
	}

	pairs := make(map[string]string)

	for len(pairs) != len(possible) {
		for allergen, ingredients := range possible {
			if len(ingredients) == 1 {
				for ingredient := range ingredients {
					pairs[allergen] = ingredient
					for _, v := range possible {
						delete(v, ingredient)
					}
				}
				break
			}
		}
	}

	keys := make([]string, 0, len(pairs))
	for k := range pairs {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	vals := make([]string, 0, len(pairs))
	for _, k := range keys {
		vals = append(vals, pairs[k])
	}

	return strings.Join(vals[:], ",")
}
