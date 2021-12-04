package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

func readData() (foods []Food) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ")", "")
		parts := strings.Split(line, "(contains")

		ingredients := strings.Split(strings.TrimSpace(parts[0]), " ")
		allergens := strings.Split(strings.TrimSpace(parts[1]), ", ")

		food := Food{Ingredients: ingredients, Allergens: allergens}
		foods = append(foods, food)
	}

	return
}

func Contains(l []string, s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}

	return false
}

func part1(foods []Food) {
	allergens := make(map[string]bool)

	for _, food := range foods {
		for _, allergen := range food.Allergens {
			allergens[allergen] = true
		}
	}

	for allergen, _ := range allergens {
		counts := make(map[string]int)
		nFoods := 0

		for _, food := range foods {
			if Contains(food.Allergens, allergen) {
				nFoods++
				for _, ingredient := range food.Ingredients {
					counts[ingredient]++
				}
			}
		}

		candidates := make([]string, 0)

		for ingredient, count := range counts {
			if count == nFoods {
				candidates = append(candidates, ingredient)
			}
		}

		fmt.Println(allergen, candidates)
	}
}

func main() {
	foods := readData()

	part1(foods)
}

// Local Variables:
// compile-command: "go build"
// End:
