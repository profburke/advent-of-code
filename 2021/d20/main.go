package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func readData() (algo string, data [][]string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	algo = scanner.Text()

	scanner.Scan() // skip blank line

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")

		data = append(data, row)
	}

	return
}

var ul, lr Point

func buildImage(data [][]string) (image map[Point]bool) {
	ul = Point{0, 0}
	lr = Point{len(data[0]), len(data)}
	image = make(map[Point]bool)

	for x := ul.X; x < lr.X; x++ {
		for y := ul.Y; y < lr.Y; y++ {
			if data[x][y] == "#" {
				image[Point{x, y}] = true
			}
		}
	}

	ul = Point{-1, -1}
	lr = Point{lr.X + 1, lr.Y + 1}
	// ul = Point{-10, -10}
	// lr = Point{lr.X + 10, lr.Y + 10}

	return
}

func addOffBoard(p Point, step int) bool {
	if step == 2 && (p.X < ul.X || p.X > lr.X || p.Y < ul.Y || p.Y > lr.Y) {
		return true
	}

	return false
}

func getNeighbors(image map[Point]bool, coords Point, sn int) (neighbors string) {
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			p := Point{coords.X + dx, coords.Y + dy}
			_, ok := image[p]
			if ok || addOffBoard(p, sn) {
				neighbors += "1" // found a "#"
			} else {
				neighbors += "0" // found a "."
			}
		}
	}

	return
}

func getIndex(neighbors string) int {
	i, _ := strconv.ParseInt(neighbors, 2, 64)
	return int(i)
}

func step(algo string, image map[Point]bool, sn int) (ni map[Point]bool) {
	ni = make(map[Point]bool)

	for x := ul.X; x < lr.X; x++ {
		for y := ul.Y; y < lr.Y; y++ {
			coords := Point{x, y}
			neighbors := getNeighbors(image, coords, sn)
			index := getIndex(neighbors)
			if string(algo[index]) == "#" {
				ni[coords] = true
			}
		}
	}

	ul = Point{ul.X - 1, ul.Y - 1}
	lr = Point{lr.X + 1, lr.Y + 1}

	return
}

func print(image map[Point]bool) {
	for x := ul.X; x < lr.X; x++ {
		for y := ul.Y; y < lr.Y; y++ {
			_, ok := image[Point{x, y}]
			var c string
			if ok {
				c = "#"
			} else {
				c = "."
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func part1(algo string, data [][]string) {
	image := buildImage(data)
	// print(image)
	// fmt.Println(len(image))

	image = step(algo, image, 1)
	// print(image)
	// fmt.Println(len(image))

	image = step(algo, image, 2)
	// print(image)
	fmt.Println(len(image))
}

func part2(algo string, data [][]string) {
}

func main() {
	algo, data := readData()
	part1(algo, data)
	part2(algo, data)
}

// Local Variables:
// compile-command: "go build"
// End:
