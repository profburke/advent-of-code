package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var risks [][]int

func readRisks() {
	risks = make([][]int, 0)
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

type RiskFunction func(p Point) int

func dijkstra(getRisks RiskFunction) int {
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
				distances[u]+getRisks(q) < distances[qdex] {
				distances[qdex] = distances[u] + getRisks(q)
			}
		}
	}

	return distances[len(distances)-1]
}

func part1() {
	fmt.Println(dijkstra(func(p Point) int { return risks[p.R][p.C] }))
}

func adjust(r, a int) int {
	for i := 0; i < a; i++ {
		r++
		if r == 10 {
			r = 1
		}
	}

	return r
}

func part2() {
	baseHeight, baseWidth := height, width
	height, width = 5*height, 5*width

	rf := func(p Point) int {
		r := p.R
		c := p.C

		br := r % baseHeight
		bc := c % baseWidth

		ar := r / baseHeight
		ac := c / baseWidth

		baseRisk := risks[br][bc]
		risk := adjust(baseRisk, ar)
		risk = adjust(risk, ac)

		return risk
	}

	fmt.Println(dijkstra(rf))
}

func main() {
	readRisks()

	height = len(risks)
	width = len(risks[0])

	part1()
	part2()
}

// Local Variables:
// compile-command: "go build"
// End:
