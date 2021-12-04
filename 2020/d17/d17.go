package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readData() (ca CA3) {
	ca = CA3{X1: 0, Y1: 0, Z1: 0, Z2: 0, W1: 0, W2: 0}
	ca.G = make(Grid)

	scanner := bufio.NewScanner(os.Stdin)

	x, y := -1, 0
	for scanner.Scan() {
		x++
		line := scanner.Text()
		cells := strings.Split(line, "")
		y = -1
		for _, cell := range cells {
			y++
			if cell == "#" {
				c := Coordinates{X: x, Y: y, Z: 0, W: 0}
				ca.G[c] = true
			}
		}
	}
	ca.X2 = x
	ca.Y2 = y

	return
}

func part1(ca CA3) {
	for i := 0; i < 6; i++ {
		ca.Step()
	}

	fmt.Println("Part 1 =", ca.ActiveCount())
}

func main() {
	ca := readData()

	part1(ca)
}

// Local Variables:
// compile-command: "go build"
// End:
