package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coordinates struct {
	X int
	Y int
}

type Direction int

const (
	West      Direction = iota
	Northwest           = iota
	Northeast           = iota
	East                = iota
	Southeast           = iota
	Southwest           = iota
)

var AllDirections = []Direction{East, West, Northeast, Northwest, Southeast, Southwest}

func (d Direction) String() string {
	return [...]string{"West", "Northwest", "Northeast", "East",
		"Southeast", "Southwest"}[d]
}

func (d Direction) From(location Coordinates) (newLocation Coordinates) {
	deltaE, deltaW := 0, -1
	if location.Y%2 != 0 {
		deltaE, deltaW = 1, 0
	}

	switch d {
	case East:
		newLocation = Coordinates{X: location.X + 1, Y: location.Y}
	case West:
		newLocation = Coordinates{X: location.X - 1, Y: location.Y}
	case Northeast:
		newLocation = Coordinates{X: location.X + deltaE, Y: location.Y - 1}
	case Northwest:
		newLocation = Coordinates{X: location.X + deltaW, Y: location.Y - 1}
	case Southeast:
		newLocation = Coordinates{X: location.X + deltaE, Y: location.Y + 1}
	case Southwest:
		newLocation = Coordinates{X: location.X + deltaW, Y: location.Y + 1}
	}

	return
}

func determineSteps(line string) (steps []Direction) {
	steps = make([]Direction, 0)

	complete := false
	var dir Direction
	var prev rune
	for _, d := range line {
		switch d {
		case 'n':
			prev = 'n'
		case 's':
			prev = 's'
		case 'e':
			if prev == 0 {
				dir = East
				complete = true
				prev = 0
			} else {
				if prev == 'n' {
					dir = Northeast
				} else {
					dir = Southeast
				}
				complete = true
				prev = 0
			}
		case 'w':
			if prev == 0 {
				dir = West
				complete = true
				prev = 0
			} else {
				if prev == 'n' {
					dir = Northwest
				} else {
					dir = Southwest
				}
				complete = true
				prev = 0
			}
		}

		if complete {
			steps = append(steps, dir)
			complete = false
		}
	}

	return
}

func readData() (data [][]Direction) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		steps := determineSteps(line)
		data = append(data, steps)
	}

	return
}

func walk(steps []Direction) (location Coordinates) {
	location = Coordinates{X: 0, Y: 0}
	for _, step := range steps {
		location = step.From(location)
	}

	return
}

func part1(data [][]Direction) (blackTiles map[Coordinates]bool) {
	blackTiles = make(map[Coordinates]bool)

	for _, steps := range data {
		location := walk(steps)

		if _, found := blackTiles[location]; found {
			delete(blackTiles, location)
		} else {
			blackTiles[location] = true
		}
	}

	fmt.Println("Part 1 =", len(blackTiles))

	return
}

func getBounds(blackTiles map[Coordinates]bool) (minX, minY, maxX, maxY int) {
	minX = math.MaxInt64
	minY = math.MaxInt64
	maxX = math.MinInt64
	maxY = math.MinInt64

	for c, _ := range blackTiles {
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	return
}

func neighbors(cell Coordinates, blackTiles map[Coordinates]bool) (count int) {

	for _, d := range AllDirections {
		l := d.From(cell)
		if _, found := blackTiles[l]; found {
			count++
		}
	}

	return
}

type CellState int

const (
	Black CellState = 0
	White           = 1
)

func getState(cell Coordinates, blackTiles map[Coordinates]bool) CellState {
	if _, found := blackTiles[cell]; found {
		return Black
	} else {
		return White
	}
}

func step(blackTiles map[Coordinates]bool) (newBlackTiles map[Coordinates]bool) {
	minX, minY, maxX, maxY := getBounds(blackTiles)
	newBlackTiles = make(map[Coordinates]bool)

	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			cell := Coordinates{X: x, Y: y}
			count := neighbors(cell, blackTiles)
			state := getState(cell, blackTiles)
			switch state {
			case Black:
				if count == 1 || count == 2 {
					newBlackTiles[cell] = true
				}
			case White:
				if count == 2 {
					newBlackTiles[cell] = true
				}
			}
		}
	}

	return
}

func part2(blackTiles map[Coordinates]bool) {

	for i := 0; i < 100; i++ {
		blackTiles = step(blackTiles)
	}

	fmt.Println("Part 2 =", len(blackTiles))
}

func main() {
	data := readData()

	blackTiles := part1(data)
	part2(blackTiles)
}

// Local Variables:
// compile-command: "go build"
// End:
