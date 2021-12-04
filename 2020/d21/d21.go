package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readData() (ingredients []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "(contains")
		ingredients = append(ingredients, strings.Split(strings.TrimSpace(parts[0]), " ")...)
	}

	return
}

func part1(ingredients []string, allergens map[string]bool) {
	count := 0

	for _, ingredient := range ingredients {
		if _, found := allergens[ingredient]; !found {
			count++
		}
	}

	fmt.Println("Part 1 =", count)
}

func main() {
	ingredients := readData()
	allergens := make(map[string]bool)
	allergens["bcdgf"] = true
	allergens["xcgtv"] = true
	allergens["dhbxtb"] = true
	allergens["scxxn"] = true
	allergens["xhrdsl"] = true
	allergens["bvcrrfbr"] = true
	allergens["vndrb"] = true
	allergens["lbnmsr"] = true

	part1(ingredients, allergens)
}

// Local Variables:
// compile-command: "go build"
// End:
