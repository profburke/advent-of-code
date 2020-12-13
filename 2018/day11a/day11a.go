// -*- compile-command: "go build"; -*-
package main

import (
	"fmt"
	"math"
)

const gridSerialNumber = 5791
//const gridSerialNumber = 42
const width = 301

func powerLevel(rackId, y, gridSerialNumber int) (powerLevel int) {
	powerLevel = (rackId * y + gridSerialNumber) * rackId
	powerLevel = (powerLevel / 100) % 10
	powerLevel -= 5

	return powerLevel
}

func newGrid(width int) (grid [][]int) {
	grid = make([][]int, width)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	return grid
}

func setLevels(grid [][]int, width int, gridSerialNumber int) {
	for x := 1; x < width; x++ {
		rackId := x + 10
		for y := 1; y < width; y++ {
			grid[x][y] = powerLevel(rackId, y, gridSerialNumber)
		}
	}
}

func sumSquare(grid [][]int, x, y int) (sum int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += grid[x + i][y + j]
		}
	}
	
	return sum
}

// need to sum _every_ grid; not just the ones where n%3 == 0
func sumGrid(grid [][]int, width int) (sums map[string]int) {
	sums = make(map[string]int)
	
	for i := 1; (i + 3) < width; i++ {
		for j := 1; (j + 3) < width; j++ {
			sum := sumSquare(grid, i, j)
			key := fmt.Sprintf("%d,%d", i, j)
			sums[key] = sum
		}
	}

	return sums
}

func main() {
	grid := newGrid(width)
	setLevels(grid, width, gridSerialNumber)
	
	sums := sumGrid(grid, width)

	index := ""
	max := math.MinInt64

	for k, v := range sums {
		if v > max {
			max = v
			index = k
		}
	}
	fmt.Println(index)
}
