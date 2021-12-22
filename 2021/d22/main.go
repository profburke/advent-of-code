package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Triple struct {
	X int
	Y int
	Z int
}

type State int

const (
	Off State = iota
	On
)

type Instruction struct {
	S    State
	Xmin int
	Xmax int
	Ymin int
	Ymax int
	Zmin int
	Zmax int
}

func parseBounds(s string) (min, max int) {
	// x=-20..26
	i := strings.Index(s, "=")
	parts := strings.Split(s[i+1:], "..")
	min, _ = strconv.Atoi(parts[0])
	max, _ = strconv.Atoi(parts[1])

	return
}

func readInstructions() (instructions []Instruction) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		var s State
		if fields[0] == "on" {
			s = On
		} else {
			s = Off
		}

		parts := strings.Split(fields[1], ",")
		xmin, xmax := parseBounds(parts[0])
		ymin, ymax := parseBounds(parts[1])
		zmin, zmax := parseBounds(parts[2])
		instructions = append(instructions, Instruction{s, xmin, xmax, ymin, ymax, zmin, zmax})
	}

	return
}

const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1

func outOfInitBounds(i Instruction) bool {
	return i.Xmin < -50 || i.Xmax > 50 ||
		i.Ymin < -50 || i.Zmax > 50 ||
		i.Zmin < -50 || i.Zmax > 50
}

func part1(instructions []Instruction) {
	cubes := make(map[Triple]bool)

	for _, i := range instructions {
		if outOfInitBounds(i) {
			continue
		}
		for x := i.Xmin; x <= i.Xmax; x++ {
			for y := i.Ymin; y <= i.Ymax; y++ {
				for z := i.Zmin; z <= i.Zmax; z++ {
					t := Triple{x, y, z}
					if i.S == Off {
						delete(cubes, t)
					} else {
						cubes[t] = true
					}
				}
			}
		}
	}

	fmt.Println(len(cubes))
}

func part2(instructions []Instruction) {
}

func main() {
	instructions := readInstructions()
	part1(instructions)
	part2(instructions)
}

// Local Variables:
// compile-command: "go build"
// End:
