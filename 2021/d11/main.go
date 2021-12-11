package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLevels() (levels [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "")
		row := make([]int, 0)
		for _, t := range numbers {
			n, _ := strconv.Atoi(t)
			row = append(row, n)
		}
		levels = append(levels, row)
	}

	return levels
}

type Point struct {
	R int
	C int
}

func onboard(q Point, height, width int) bool {
	return q.R >= 0 && q.R < height && q.C >= 0 && q.C < width
}

func part1(levels [][]int) {
	flashes := 0
	width := len(levels[0])
	height := len(levels)

	// for _, r := range levels {
	// 	fmt.Println(r)
	// }

	for step := 0; step < 100; step++ {
		// increase each by 1
		for r, row := range levels {
			for c, o := range row {
				levels[r][c] = o + 1
			}
		}

		// do flashing
		flashers := make(map[Point]bool)
		queue := make([]Point, 0)

		for r, row := range levels {
			for c, o := range row {
				if o > 9 {
					queue = append(queue, Point{R: r, C: c})
				}
			}
		}

		for len(queue) > 0 {
			// grab from queue
			p := queue[0]
			queue = queue[1:]

			_, flashed := flashers[p]
			if flashed {
				continue
			}

			// add it to flashers
			flashers[p] = true
			flashes++

			// flash it...and add flashing neighbors to queue
			for dr := -1; dr < 2; dr++ {
				for dc := -1; dc < 2; dc++ {
					if dc == 0 && dr == 0 {
						continue
					}

					q := Point{R: p.R + dr, C: p.C + dc}
					if onboard(q, height, width) {
						levels[q.R][q.C] += 1
						_, flashed := flashers[q]
						if !flashed && levels[q.R][q.C] > 9 {
							queue = append(queue, q)
						}
					}
				}
			}
		}

		// reset flashed
		for r, row := range levels {
			for c, _ := range row {
				_, flashed := flashers[Point{R: r, C: c}]
				if flashed {
					levels[r][c] = 0
				}
			}
		}

		// fmt.Println("\n\n\n")

		// for _, r := range levels {
		// 	fmt.Println(r)
		// }
	}

	fmt.Println(flashes)
}

func part2(levels [][]int) {
}

func main() {
	levels := readLevels()
	part1(levels)
	part2(levels)
}

// Local Variables:
// compile-command: "go build"
// End:
