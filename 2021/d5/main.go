package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) sum(a Point) Point {
	return Point{X: a.X + p.X, Y: a.Y + p.Y}
}

type Line struct {
	E1 Point
	E2 Point
}

func readLines() (lines []Line) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inline := scanner.Text()
		pieces := strings.Split(inline, " -> ")
		coordinates := strings.Split(pieces[0], ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		p1 := Point{X: x, Y: y}

		coordinates = strings.Split(pieces[1], ",")
		x, _ = strconv.Atoi(coordinates[0])
		y, _ = strconv.Atoi(coordinates[1])
		p2 := Point{X: x, Y: y}

		line := Line{E1: p1, E2: p2}
		lines = append(lines, line)
	}

	return
}

func isHorizontal(line Line) bool {
	return line.E1.Y == line.E2.Y
}

func isVertical(line Line) bool {
	return line.E1.X == line.E2.X
}

func isDiagonal(line Line) bool {
	return math.Abs(float64(line.E1.X-line.E2.X)) == math.Abs(float64(line.E1.Y-line.E2.Y))
}

// will only be called for horizontal, vertical
// and diagonal lines. So dx, dy elements of { -1, 0, 1 }
func walk(line Line) (points []Point) {
	var s, e, delta Point

	if isHorizontal(line) {
		delta = Point{X: 1, Y: 0}

		if line.E1.X < line.E2.X {
			s = line.E1
			e = line.E2
		} else {
			s = line.E2
			e = line.E1
		}
	} else if isVertical(line) {
		delta = Point{X: 0, Y: 1}

		if line.E1.Y < line.E2.Y {
			s = line.E1
			e = line.E2
		} else {
			s = line.E2
			e = line.E1
		}
	} else {
		if line.E1.X < line.E2.X {
			s = line.E1
			e = line.E2
		} else {
			s = line.E2
			e = line.E1
		}

		if s.Y > e.Y {
			delta = Point{X: 1, Y: -1}
		} else {
			delta = Point{X: 1, Y: 1}
		}
	}

	for {
		points = append(points, s)
		if s == e {
			break
		} else {
			s = s.sum(delta)
		}
	}

	return
}

func part1(lines []Line) {
	dangerousPoints := 0
	vents := make(map[Point]int)

	for _, line := range lines {
		if isHorizontal(line) || isVertical(line) {
			points := walk(line)
			for _, point := range points {
				vents[point]++
			}
		}
	}

	for _, count := range vents {
		if count > 1 {
			dangerousPoints++
		}
	}

	fmt.Println(dangerousPoints)
}

func part2(lines []Line) {
	dangerousPoints := 0
	vents := make(map[Point]int)

	for _, line := range lines {
		if isHorizontal(line) || isVertical(line) || isDiagonal(line) {
			points := walk(line)
			for _, point := range points {
				vents[point]++
			}
		}
	}

	for _, count := range vents {
		if count > 1 {
			dangerousPoints++
		}
	}

	fmt.Println(dangerousPoints)
}

func main() {
	lines := readLines()
	part1(lines)
	part2(lines)
}

// Local Variables:
// compile-command: "go build"
// End:
