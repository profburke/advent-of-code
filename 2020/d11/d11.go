package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readData() (grid Grid) {
	scanner := bufio.NewScanner(os.Stdin)

	now := make([][]string, 0)
	future := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		now = append(now, row)
		future = append(future, row)
	}

	grid.Grid = make([][][]string, 2)
	grid.Current = 0
	grid.Next = 1
	grid.Grid[grid.Current] = now
	grid.Grid[grid.Next] = future

	return
}

func part1(grid Grid) {
	occupiedSeats := 0

	for true {
		fmt.Println("step")
		grid.Step()
		if grid.IsStable() {
			occupiedSeats = grid.Count(OccupiedSeat)
			break
		}
	}

	fmt.Println("Part 1 =", occupiedSeats)
}

func part2(grid Grid) {
}

func main() {
	grid := readData()

	part1(grid)
	part2(grid)
}

// Local Variables:
// compile-command: "go build"
// End:
