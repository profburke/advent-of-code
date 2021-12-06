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

func part1(fish []int) {
	for day := 0; day < 80; day++ {
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

func part2(fish []int) {
}

func main() {
	fish := readFish()
	fmt.Println(fish)
	part1(fish)
	part2(fish)
}

// Local Variables:
// compile-command: "go build"
// End:
