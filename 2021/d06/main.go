package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFish() (fish []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	numbers := strings.Split(line, ",")

	for _, n := range numbers {
		f, _ := strconv.Atoi(n)
		fish = append(fish, f)
	}

	return
}

func compute(days int, fish []int) {
	for day := 0; day < days; day++ {
		newfish := make([]int, 0)
		for _, f := range fish {
			switch f {
			case 0:
				newfish = append(newfish, 6)
				newfish = append(newfish, 8)
			default:
				newfish = append(newfish, f-1)
			}
		}

		fish = newfish
	}

	fmt.Println(len(fish))
}

func part1(fish []int) {
	compute(80, fish)
}

func part2(fish []int) {
	population := make([]int, 9)
	for _, f := range fish {
		population[f]++
	}

	for day := 0; day < 256; day++ {
		p0 := population[0]
		for s := 0; s < 8; s++ {
			population[s] = population[s+1]
		}
		population[8] = p0
		population[6] += p0
	}

	total := 0
	for _, p := range population {
		total += p
	}
	fmt.Println(total)
}

func main() {
	fish := readFish()
	part1(fish)
	part2(fish)
}

// Local Variables:
// compile-command: "go build"
// End:
