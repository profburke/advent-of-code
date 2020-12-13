// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

type Point struct {
	Name int
	X int
	Y int
}

func (a Point) String() string {
	return fmt.Sprintf("(%d: %d, %d)", a.Name, a.X, a.Y)
}

func (a Point) manhattanDistance(b Point) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		} else {
			return x
		}
	}
	
	return abs(a.X - b.X) + abs(a.Y - b.Y)
}

func readPoints(r io.Reader) (points []Point, err error) {
	scanner := bufio.NewScanner(r)
	id := 0
	
	for scanner.Scan() {
		var (
			x, y int
		)
		
		text := scanner.Text()
		matches, err := fmt.Sscanf(text, "%d, %d", &x, &y)
		if err != nil || matches != 2 {
			return points, err
		}
		points = append(points, Point{id, x, y})
		id++
	}

	err = scanner.Err()
	return points, err
}

func findBasin(points []Point, i int, j int) int {
	ties := make(map[int]int)
	p0 := Point{-3, i, j}
	id := -1
	minDist := math.MaxInt64
	
	for _, p := range points {
		d := p0.manhattanDistance(p)
		ties[d]++
		if d < minDist {
			minDist = d
			id = p.Name
		}
	}

	if ties[minDist] > 1 {
		return -1
	} else {
		return id
	}
}

func main() {
	points, err := readPoints(os.Stdin)
	if err != nil {
		fmt.Println("error reading points: ", err)
		os.Exit(1)
	}

	grid := make([][]int, 360)
	for i := range grid {
		grid[i] = make([]int, 360)
	}

	counts := make(map[int]int)

	for i := range grid {
		row := grid[i]
		for j := range row {
			cp := findBasin(points, i, j)
			grid[i][j] = cp
			counts[cp]++
		}
	}

	// Now remove the basins that are infinite.
	// If a basin contains a point on the boundary, than it is infinite.

	for i := 0; i < 360; i++ {
		cp := grid[i][0]
		delete(counts, cp)
		cp = grid[i][359]
		delete(counts, cp)
		cp = grid[0][i]
		delete(counts, cp)
		cp = grid[359][i]
		delete(counts, cp)
	}

	// find max of remainder
	max := 0

	for _, v := range counts {
		if v > max {
			max = v
		}
	}
	
	fmt.Println("max = ", max)
}
