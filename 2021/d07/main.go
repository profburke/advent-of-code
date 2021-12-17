package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readPositions() (positions []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	numbers := strings.Split(line, ",")

	for _, n := range numbers {
		p, _ := strconv.Atoi(n)
		positions = append(positions, p)
	}

	return
}

const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1

func intAbs(n int) int {
	return int(math.Abs(float64(n)))
}

type FuelCost func(int) int

func compute(positions []int, f FuelCost) (fuelSpent int) {
	fuelSpent = MaxInt
	maxPosition := 0

	for _, p := range positions {
		if p > maxPosition {
			maxPosition = p
		}
	}

	for newPos := 0; newPos <= maxPosition; newPos++ {
		fuel := 0
		for _, p := range positions {
			moveDistance := intAbs(newPos - p)
			fuel += f(moveDistance)
		}

		if fuel < fuelSpent {
			fuelSpent = fuel
		}
	}

	return
}

func part1(positions []int) {
	f := func(d int) int { return d }
	fmt.Println(compute(positions, f))
}

func part2(positions []int) {
	f := func(d int) int { return (d * (d + 1)) / 2 }
	fmt.Println(compute(positions, f))
}

func main() {
	positions := readPositions()
	part1(positions)
	part2(positions)
}

// Local Variables:
// compile-command: "go build"
// End:
