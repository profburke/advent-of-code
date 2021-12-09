package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readHeights() (heights [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "")

		depthLine := make([]int, 0)
		for _, number := range numbers {
			n, _ := strconv.Atoi(number)
			depthLine = append(depthLine, n)
		}
		heights = append(heights, depthLine)
	}

	return heights
}

func getNeighbors(i, j int, heights [][]int) (neighbors []int) {
	height := len(heights)
	width := len(heights[0])

	if i != 0 {
		neighbors = append(neighbors, heights[i-1][j])
	}

	if i != (height - 1) {
		neighbors = append(neighbors, heights[i+1][j])
	}

	if j != 0 {
		neighbors = append(neighbors, heights[i][j-1])
	}

	if j != (width - 1) {
		neighbors = append(neighbors, heights[i][j+1])
	}

	return
}

const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1

func minimum(ns []int) int {
	minimum := MaxInt

	for _, n := range ns {
		if n < minimum {
			minimum = n
		}
	}

	return minimum
}

func part1(heights [][]int) {
	risk := 0

	for i, row := range heights {
		for j, h := range row {
			neighbors := getNeighbors(i, j, heights)
			m := minimum(neighbors)
			if h < m {
				risk += (1 + h)
			}
		}
	}

	fmt.Println(risk)
}

func part2(heights [][]int) {
}

func main() {
	heights := readHeights()
	part1(heights)
	part2(heights)
}

// Local Variables:
// compile-command: "go build"
// End:
