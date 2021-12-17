package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readDepths() (depths []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		depths = append(depths, i)
	}

	return depths
}

const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1

func part1(depths []int) {
	previous := MaxInt
	result := 0

	for _, depth := range depths {
		if depth > previous {
			result++
		}
		previous = depth
	}

	fmt.Fprintln(os.Stdout, "Part 1: Number of increases = ", result)
}

func part2(depths []int) {
	previous := MaxInt
	result := 0

	for i := 0; i+2 < len(depths); i++ {
		sum := depths[i] + depths[i+1] + depths[i+2]
		if sum > previous {
			result++
		}
		previous = sum
	}

	fmt.Fprintln(os.Stdout, "Part 2: Number of increases = ", result)
}

func main() {
	depths := readDepths()
	part1(depths)
	part2(depths)
}

// Local Variables:
// compile-command: "go build"
// End:
