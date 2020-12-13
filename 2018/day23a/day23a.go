// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Position struct {
	X, Y, Z int
}

type Nanobot struct {
	Position Position
	SignalStrength int
}

func readNanobots(r io.Reader) (nanobots []Nanobot, strongest int, err error) {
	scanner := bufio.NewScanner(r)
	maxSignal := -1
	strongest = -1

	var (
		x, y, z, s int
	)

	n := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &s)
		if err != nil || matches != 4 {
			return nil, -1, errors.New(fmt.Sprintf("error reading nanobot %d", n))
		}
		if s > maxSignal {
			maxSignal = s
			strongest = n
		}
		nano := Nanobot{Position{x, y, z}, s}
		nanobots = append(nanobots, nano)
		n++
	}

	err = scanner.Err()
	return
}

func abs(a int) (aa int) {
	if a >= 0 { return a }
	return -a
}

func mDistance(a, b Position) (d int) {
	return abs(a.X - b.X) + abs(a.Y - b.Y) + abs(a.Z - b.Z)
}

func countInRange(nanobots []Nanobot, strongest int) (count int) {
	strongestNano := nanobots[strongest]
	snp := strongestNano.Position
	snr := strongestNano.SignalStrength
	
	for _, nano := range nanobots {
		np := nano.Position
		if mDistance(np, snp) <= snr {
			count++
		}
	}
	return
}

func main() {
	nanobots, strongest, err := readNanobots(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading nanobot info: ", err)
		os.Exit(1)
	}

	fmt.Println("strongest: ", strongest)
	count := countInRange(nanobots, strongest)
	fmt.Println(count)
}
