package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readRisks() (risks [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "")
		row := make([]int, 0)
		for _, n := range numbers {
			r, _ := strconv.Atoi(n)
			row = append(row, r)
		}
		risks = append(risks, row)
	}

	return risks
}

const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1

type Point struct {
	R int
	C int
}

var height, width int

func (p Point) onboard() bool {
	return p.R >= 0 && p.R < height && p.C >= 0 && p.C < width
}

func (p Point) neighbors() (ns []Point) {
	ns = make([]Point, 0)

	q := Point{R: p.R - 1, C: p.C}
	if q.onboard() {
		ns = append(ns, q)
	}

	q = Point{R: p.R + 1, C: p.C}
	if q.onboard() {
		ns = append(ns, q)
	}

	q = Point{R: p.R, C: p.C - 1}
	if q.onboard() {
		ns = append(ns, q)
	}

	q = Point{R: p.R, C: p.C + 1}
	if q.onboard() {
		ns = append(ns, q)
	}

	return
}

func minOf(xs []int) (min int) {
	min = MaxInt

	for _, x := range xs {
		if x < min {
			min = x
		}
	}

	return
}

func toCoordinates(i, width int) (r, c int) {
	r = i / width
	c = i % width

	return
}

func toIndex(p Point, width int) int {
	return p.R*width + p.C
}

func minDistance(distances []int, sptSet map[Point]bool, width int) int {
	min := MaxInt
	mindex := -1

	for i := 0; i < height*width; i++ {
		r, c := toCoordinates(i, width)
		_, ok := sptSet[Point{R: r, C: c}]
		if !ok && distances[i] < min {
			min = distances[i]
			mindex = i
		}
	}

	return mindex
}

func dijkstra(risks [][]int) int {
	height := len(risks)
	width := len(risks[0])
	nv := height * width

	distances := make([]int, nv)
	for i := range distances {
		distances[i] = MaxInt
	}
	distances[0] = 0

	sptSet := make(map[Point]bool)

	for count := 0; count < nv-1; count++ {

		u := minDistance(distances, sptSet, width)
		r, c := toCoordinates(u, width)
		p := Point{R: r, C: c}
		sptSet[p] = true

		for _, q := range p.neighbors() {
			_, ok := sptSet[q]
			qdex := toIndex(q, width)
			if !ok && distances[u] != MaxInt &&
				distances[u]+risks[q.R][q.C] < distances[qdex] {
				distances[qdex] = distances[u] + risks[q.R][q.C]
			}
		}
	}

	return distances[len(distances)-1]
}

func part1(risks [][]int) {
	fmt.Println(dijkstra(risks))
}

func buildMap(m [][]int) (f [][]int) {
	f = make([][]int, 0)

	for _, row := range m {
		newrow := make([]int, 0)
		f = append(f, newrow)
	}

	return
}

func part2(risks [][]int) {
	fullrisks := buildMap(risks)
	fmt.Println(dijkstra(fullrisks))
}

func main() {
	risks := readRisks()

	height = len(risks)
	width = len(risks[0])

	part1(risks)
	part2(risks)
}

// Local Variables:
// compile-command: "go build"
// End:
