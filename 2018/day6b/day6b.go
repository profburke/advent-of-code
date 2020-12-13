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
	X int
	Y int
}

func (a Point) String() string {
	return fmt.Sprintf("(%d, %d)", a.X, a.Y)
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
	
	for scanner.Scan() {
		var (
			x, y int
		)
		
		text := scanner.Text()
		matches, err := fmt.Sscanf(text, "%d, %d", &x, &y)
		if err != nil || matches != 2 {
			return points, err
		}
		points = append(points, Point{x, y})
	}

	err = scanner.Err()
	return points, err
}

func totalDistance(points []Point, x, y int) (d int) {
	d = 0
	p0 := Point{x, y}
	
	for _, p := range points {
		d += p0.manhattanDistance(p)
	}
	
	return d
}

func bounds(points []Point) (xmin, xmax, ymin, ymax int) {
	xmin = math.MaxInt64
	ymin = math.MaxInt64
	xmax = math.MinInt64
	ymax = math.MinInt64

	for _, p := range points {
		if p.X > xmax { xmax = p.X }
		if p.Y > ymax { ymax = p.Y }
		if p.X < xmin { xmin = p.X }
		if p.Y < ymin { ymin = p.Y }
	}
	
	return xmin, xmax, ymin, ymax
}

func main() {
	points, err := readPoints(os.Stdin)
	if err != nil {
		fmt.Println("error reading points: ", err)
		os.Exit(1)
	}

	count := 0
	xmin, xmax, ymin, ymax := bounds(points)
	fudge := 10000/len(points) // distance to a point, can't be all the 10k
	xmin -= fudge
	xmax += fudge
	ymin -= fudge
	ymax += fudge

	fmt.Println("x ranges from ", xmin, " to ", xmax)
	fmt.Println("y ranges from ", ymin, " to ", ymax)

	for x := xmin; x < xmax; x++ {
		for y := ymin; y < ymax; y++ {
			d := totalDistance(points, x, y)
			if d < 10000 {
				count++
			}
		}
	}
	
	fmt.Println("count = ", count)
}
