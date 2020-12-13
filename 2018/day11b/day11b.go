// -*- compile-command: "go build"; -*-
package main

import (
	"fmt"
	"math"
)

const gridSerialNumber = 5791
// const gridSerialNumber = 42
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

func sumSquare(grid [][]int, x, y, dial int) (sum int) {
	for i := 0; i < dial; i++ {
		for j := 0; j < dial; j++ {
			sum += grid[x + i][y + j]
		}
	}
	
	return sum
}

// need to sum _every_ grid; not just the ones where n%3 == 0
func sumGrid(grid [][]int, width, dial int) (sums map[string]int) {
	sums = make(map[string]int)
	
	for i := 1; (i + dial) < width; i++ {
		for j := 1; (j + dial) < width; j++ {
			sum := sumSquare(grid, i, j, dial)
			key := fmt.Sprintf("%d,%d", i, j)
			sums[key] = sum
		}
	}

	return sums
}

func maxForDial(grid [][]int, width, dial int) (index string, max int) {
	sums := sumGrid(grid, width, dial)

	index = ""
	max = math.MinInt64

	for k, v := range sums {
		if v > max {
			max = v
			index = k
		}
	}

	return index, max
}

func main() {
	grid := newGrid(width)
	setLevels(grid, width, gridSerialNumber)
	
	dialMaxes := make(map[string]int)
	index := ""
	max := math.MinInt64
	
	for dial := 1; dial < width; dial++ {
		index, max := maxForDial(grid, width, dial)
		key := fmt.Sprintf("%s - %d", index, dial)
		dialMaxes[key] = max
	}

	for k, v := range dialMaxes {
		if v > max {
			max = v
			index = k
		}
	}
	
	fmt.Println(index)
}
