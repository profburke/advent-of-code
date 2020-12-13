// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func abs(a int) (aa int) {
	if a >= 0 {
		aa = a 
	} else {
		aa = -a
	}
	return
}
	
type Point struct {
	X, Y, Z, T int
}

func (p Point) String() string {
	return fmt.Sprintf("<%d, %d, %d, %d>", p.X, p.Y, p.Z, p.T)
}

func (p Point) distance(q Point) (d int) {
	d = abs(p.X - q.X) + abs(p.Y - q.Y) + abs(p.Z - q.Z) + abs(p.T - q.T)
	return
}

func readPoints(r io.Reader) (points []Point) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var (
			x, y, z, t int
		)
		
		line := scanner.Text()
		matches, err := fmt.Sscanf(line, "%d,%d,%d,%d", &x, &y, &z, &t)
		if err != nil || matches != 4 {
			fmt.Fprintln(os.Stderr, "error reading point: ", err)
			os.Exit(1)
		}
		points = append(points, Point{x, y, z, t})
	}
	return
}

func inConstellation(constellations map[Point]int, p Point) (result bool) {
	_, result = constellations[p]
	return
}

func main() {
	constellations := make(map[Point]int)

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

	points := readPoints(fh)

	for i := 0; i < len(points); i++ {
		candidate := points[i]
		if !inConstellation(constellations, candidate) {
//			fmt.Println("Inserting ", candidate, " at ", i)
			constellations[candidate] = i
		}
		
		for j := i + 1; j < len(points); j++ {
			other := points[j]

			if candidate.distance(other) < 4 {
				//fmt.Println(candidate, " is in range of ", other)
				if !inConstellation(constellations, other) {
					constellations[other] = constellations[candidate]
				} else {
					// change all the candidate's constellation numbers to other's number
					cNum := constellations[candidate]
					oNum := constellations[other]
					for k, v := range constellations {
						if v == cNum {
							constellations[k] = oNum
						}
					}
				}
			}
		}
	}

	usedNums := make(map[int]bool)
	for k, v := range constellations {
		usedNums[v] = true
		fmt.Println(k, " is in ", v)
	}

	fmt.Println("# of constellations: ", len(usedNums))
}
