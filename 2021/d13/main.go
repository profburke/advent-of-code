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

type Instruction struct {
	Axis  string
	Value int
}

func readDots() (dots []Point, instructions []Instruction) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])

		dots = append(dots, Point{X: x, Y: y})
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimPrefix(line, "fold along")
		parts := strings.Split(line, "=")
		v, _ := strconv.Atoi(parts[1])

		instructions = append(instructions, Instruction{Axis: parts[0], Value: v})
	}

	return
}

func mAbs(v int) int {
	return int(math.Abs(float64(v)))
}

func foldUp(p Point, v int) Point {
	d := mAbs(p.Y - v)
	return Point{X: p.X, Y: v + d}
}

func foldLeft(p Point, v int) Point {
	d := mAbs(p.X - v)
	return Point{X: v - d, Y: p.Y}
}

func part1(dots []Point, instructions []Instruction) {
	translated := make([]Point, 0)

	i := instructions[0]
	for _, p := range dots {
		var newP Point
		if i.Axis == "y" {
			newP = foldUp(p, i.Value)
		} else {
			newP = foldLeft(p, i.Value)
		}
		translated = append(translated, newP)
	}

	allPoints := make(map[Point]bool)
	for _, p := range translated {
		allPoints[p] = true
	}

	translated = make([]Point, 0)
	for p := range allPoints {
		translated = append(translated, p)
	}

	fmt.Println(len(translated))
}

func part2(dots []Point, instructions []Instruction) {
}

func main() {
	dots, instructions := readDots()
	part1(dots, instructions)
	part2(dots, instructions)
}

// Local Variables:
// compile-command: "go build"
// End:
