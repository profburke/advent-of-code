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
		line = strings.TrimPrefix(line, "fold along ")
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
	return Point{X: p.X, Y: v - d}
}

func foldLeft(p Point, v int) Point {
	d := mAbs(p.X - v)
	return Point{X: v - d, Y: p.Y}
}

func fold(dots []Point, i Instruction) (newDots []Point) {
	translated := make([]Point, 0)

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

	newDots = translated
	return
}

func part1(dots []Point, instructions []Instruction) {
	dots = fold(dots, instructions[0])
	fmt.Println(len(dots))
}

// apparently these are all defined in go 1.17; too bad I'm using 1.16 and lazy...
const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1
const MinInt = -MaxInt // this isn't quite right, but it's close enough

func findMin(ns []int) int {
	min := MaxInt

	for _, n := range ns {
		if n < min {
			min = n
		}
	}

	return min
}

func findMax(ns []int) int {
	max := MinInt

	for _, n := range ns {
		if n > max {
			max = n
		}
	}

	return max
}

func display(dots []Point) {
	xs := make([]int, 0)
	ys := make([]int, 0)

	for _, p := range dots {
		xs = append(xs, p.X)
		ys = append(ys, p.Y)
	}

	minX := findMin(xs)
	minY := findMin(ys)
	maxX := findMax(xs)
	maxY := findMax(ys)

	allPoints := make(map[Point]bool)
	for _, p := range dots {
		allPoints[p] = true
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := allPoints[Point{X: x, Y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func part2(dots []Point, instructions []Instruction) {
	for _, i := range instructions {
		dots = fold(dots, i)
	}

	display(dots)
}

func main() {
	dots, instructions := readDots()
	part1(dots, instructions)
	part2(dots, instructions)
}

// Local Variables:
// compile-command: "go build"
// End:
