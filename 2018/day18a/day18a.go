// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const size = 50 // really ought to set this based on input

type State int

const (
	Open State = iota
	Trees
	Lumberyard
)

func NewGrid() (grid [][]State) {
	grid = make([][]State, size)
	for i := range grid {
		grid[i] = make([]State, size)
	}

	return
}

func readGrid(r io.Reader) (grid [][]State) {
	grid = NewGrid()
	scanner := bufio.NewScanner(r)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		for col, c := range line {
			switch string(c) {
			case ".":
				grid[row][col] = Open
			case "|":
				grid[row][col] = Trees
			case "#":
				grid[row][col] = Lumberyard
			}
		}
		row++
	}
	
	return grid
}

func newState(grid [][]State, x, y int) (s State) {
	nOpen, nTrees, nLumberyard := 0, 0, 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			m := x + i
			n := y + j
			if m < 0 || m >= size || n < 0 || n >= size || (i == 0 && j == 0) { continue }
			
			switch grid[m][n] {
			case Open:
				nOpen++
			case Trees:
				nTrees++
			case Lumberyard:
				nLumberyard++
			}
		}
	}

	switch grid[x][y] {
	case Open:
		if nTrees >= 3 {
			s = Trees
		} else {
			s = Open
		}
	case Trees:
		if nLumberyard >= 3 {
			s = Lumberyard
		} else {
			s = Trees
		}
	case Lumberyard:
		if nLumberyard >= 1 && nTrees >= 1 {
			s = Lumberyard
		} else {
			s = Open
		}
	}
	
	return
}

func total(grid [][]State) (nOpen, nTrees, nLumberyard int) {
	for _, row := range grid {
		for _, cell := range row {
			switch cell {
			case Open:
				nOpen++
			case Trees:
				nTrees++
			case Lumberyard:
				nLumberyard++
			}
		}
	}
	return
}

func printGrid(grid [][]State, generation int) {
	fmt.Println("Generation: ", generation)
	for _, row := range grid {
		for j := range row {
			switch row[j] {
			case Open:
				fmt.Print(".")
			case Trees:
				fmt.Print("|")
			case Lumberyard:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <filename>", os.Args[0])
		os.Exit(1)
	}

	fh, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file: ", err)
		os.Exit(1)
	}
	defer fh.Close()
	
	reader := bufio.NewReader(fh)
	grid := readGrid(reader)

	generation := 0
	// max := 1000000000
	max := 100000
	for generation < max {
		// printGrid(grid, generation)
		newGrid := NewGrid()
		
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				newGrid[i][j] = newState(grid, i, j)
			}
		}
		grid = newGrid
		generation++
		_, nTrees, nLumberyard := total(grid)
		fmt.Println("Generation: ", generation, " resource value: ", nLumberyard * nTrees)
	}
	// printGrid(grid, generation)
}


// repeats every 28 generations for N sufficiently large
// so which step is it on at t = 1000000000

// so 28 * i + 544 is 206702
// so that means generation 999999992 is 206702
// so step 8 past in the cycle...
