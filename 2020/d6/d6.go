package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readData() (data [][]map[string]bool) {
	scanner := bufio.NewScanner(os.Stdin)
	group := make([]map[string]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()
		person := make(map[string]bool)

		if line == "" {
			data = append(data, group)
			group = make([]map[string]bool, 0)
			continue
		}

		for _, c := range strings.Split(line, "") {
			person[c] = true
		}

		group = append(group, person)
	}
	data = append(data, group)

	return
}

func part1(data [][]map[string]bool) {
	sum := 0

	for _, group := range data {
		yesQuestions := make(map[string]bool)

		for _, person := range group {
			for key, _ := range person {
				yesQuestions[key] = true
			}
		}

		sum += len(yesQuestions)
	}

	fmt.Println("Part 1 =", sum)
}

func part2(data [][]map[string]bool) {
	sum := 0

	for _, group := range data {
		groupSize := len(group)
		yesQuestions := make(map[string]int)

		for _, person := range group {
			for key, _ := range person {
				yesQuestions[key]++
			}
		}

		for key, _ := range yesQuestions {
			if yesQuestions[key] == groupSize {
				sum++
			}
		}
	}

	fmt.Println("Part 2 =", sum)
}

func main() {
	data := readData()

	part1(data)
	part2(data)
}

// Local Variables:
// compile-command: "go build"
// End:
