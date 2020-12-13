// -*- compile-command: "go build"; -*-
package main

import "fmt"

type Position struct {
	X, Y int
}

func main() {
	const DEPTH = 11109
	target := Position{9, 731}

	geoIndices := make(map[Position]int)
	erosionLevels := make(map[Position]int)

	for x := 0; x <= target.X; x++ {
		v := 16807 * x
		geoIndices[Position{x, 0}] = v
		erosionLevels[Position{x, 0}] = (v + DEPTH) % 20183
	}

	for y := 0; y <= target.Y; y++ {
		v := 48271 * y
		geoIndices[Position{0, y}] = v
		erosionLevels[Position{0, y}] = (v + DEPTH) % 20183
	}

	for x := 1; x <= target.X; x++ {
		for y := 1; y <= target.Y; y++ {
			v := erosionLevels[Position{x - 1, y}] * erosionLevels[Position{x, y - 1}]
			geoIndices[Position{x, y}] = v
			erosionLevels[Position{x, y}] = (v + DEPTH) % 20183
		}
	}

	geoIndices[target] = 0
	erosionLevels[target] = DEPTH * 20183
	
	riskLevel := 0
	for _, v := range erosionLevels {
		riskLevel += (v % 3)
	}

	fmt.Println("Risk Level: ", riskLevel)
}
