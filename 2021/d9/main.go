package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Point struct {
	Row int
	Col int
}

func part1(heights [][]int) (lps []Point) {
	risk := 0

	for i, row := range heights {
		for j, h := range row {
			neighbors := getNeighbors(i, j, heights)
			m := minimum(neighbors)
			if h < m {
				lps = append(lps, Point{Row: i, Col: j})
				risk += (1 + h)
			}
		}
	}

	fmt.Println(risk)

	return
}

func threeLargestProduct(ns []int) int {
	sort.Ints(ns)
	l := len(ns)

	return (ns[l-1] * ns[l-2] * ns[l-3])
}

func getBasinSize(heights [][]int, lp Point) int {
	visited := make(map[Point]bool)
	queue := make([]Point, 0)

	height := len(heights)
	width := len(heights[0])

	queue = append(queue, lp)
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		// Could simplify this a lot if we used points in part 1 and
		// rewrote getNeighbors accordingly

		// for each neighbor, if not seen, and height < 9, add to queue

		i := c.Row
		j := c.Col

		if i != 0 {
			x, y := i-1, j
			_, seen := visited[Point{Row: x, Col: y}]
			if !seen && heights[i-1][j] < 9 {
				queue = append(queue, Point{Row: i - 1, Col: j})
			}
		}

		if i != (height - 1) {
			x, y := i+1, j
			_, seen := visited[Point{Row: x, Col: y}]
			if !seen && heights[i+1][j] < 9 {
				queue = append(queue, Point{Row: i + 1, Col: j})
			}
		}

		if j != 0 {
			x, y := i, j-1
			_, seen := visited[Point{Row: x, Col: y}]
			if !seen && heights[i][j-1] < 9 {
				queue = append(queue, Point{Row: i, Col: j - 1})
			}
		}

		if j != (width - 1) {
			x, y := i, j+1
			_, seen := visited[Point{Row: x, Col: y}]
			if !seen && heights[i][j+1] < 9 {
				queue = append(queue, Point{Row: i, Col: j + 1})
			}
		}

		// Finally, mark the current point as in the basin
		visited[c] = true
	}

	return len(visited)
}

func getBasinSizes(heights [][]int, lps []Point) (sizes []int) {
	for _, lp := range lps {
		fmt.Println("calculating basin for", lp)
		size := getBasinSize(heights, lp)
		sizes = append(sizes, size)
	}
	return
}

func part2(heights [][]int, lps []Point) {
	basinSizes := getBasinSizes(heights, lps)
	basinProduct := threeLargestProduct(basinSizes)
	fmt.Println(basinProduct)
}

func main() {
	heights := readHeights()
	lps := part1(heights)
	part2(heights, lps)
}

// Local Variables:
// compile-command: "go build"
// End:
