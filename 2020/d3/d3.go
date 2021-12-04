package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const tree string = "#"

type Position struct {
	X, Y int
}

// NOTE: because after decades of programming rows/columns vs x/y
// always gives me headaches... treat the coordinates opposite from
// what they should be.
func (p Position) addMod(other Position, bound int) (sum Position) {
	sum.Y = (p.Y + other.Y) % bound
	sum.X = p.X + other.X

	return
}

func readData() (trees [][]string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, strings.Split(line, ""))
	}

	return
}

func countCollisions(trees [][]string, slope Position) int {
	rows := len(trees)
	columns := len(trees[0])
	treeCount := 0
	position := Position{X: 0, Y: 0}

	for position.X < rows {
		if trees[position.X][position.Y] == tree {
			treeCount++
		}

		position = position.addMod(slope, columns)
	}

	return treeCount
}

func collisionProduct(trees [][]string, slopes []Position) (result int) {
	result = 1

	for _, slope := range slopes {
		result *= countCollisions(trees, slope)
	}

	return
}

func part1(trees [][]string) {
	slopes := []Position{
		Position{X: 1, Y: 3},
	}

	fmt.Println("Part 1 =", collisionProduct(trees, slopes))
}

func part2(trees [][]string) {
	slopes := []Position{
		Position{X: 1, Y: 1},
		Position{X: 1, Y: 3},
		Position{X: 1, Y: 5},
		Position{X: 1, Y: 7},
		Position{X: 2, Y: 1},
	}

	fmt.Println("Part 2 =", collisionProduct(trees, slopes))
}

func main() {
	trees := readData()

	part1(trees)
	part2(trees)
}

// Local Variables:
// compile-command: "go build"
// End:
