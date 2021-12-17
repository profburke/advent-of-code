package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Value     int
}

func readCourse() (course []Command) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Fields(line)
		direction := pieces[0]
		value, _ := strconv.Atoi(pieces[1])
		command := Command{direction, value}
		course = append(course, command)
	}

	return
}

func part1(course []Command) {
	x := 0
	d := 0

	for _, c := range course {
		switch c.Direction {
		case "forward":
			x += c.Value
		case "down":
			d += c.Value
		case "up":
			d -= c.Value
		default:
			fmt.Fprintln(os.Stderr, "whoops")
		}
	}

	fmt.Fprintln(os.Stdout, "Part 1: ", x*d)
}

func part2(course []Command) {
	x := 0
	d := 0
	a := 0

	for _, c := range course {
		switch c.Direction {
		case "forward":
			x += c.Value
			d += (a * c.Value)
		case "down":
			a += c.Value
		case "up":
			a -= c.Value
		default:
			fmt.Fprintln(os.Stderr, "whoops")
		}
	}

	fmt.Fprintln(os.Stdout, "Part 2: ", x*d)
}

func main() {
	course := readCourse()

	part1(course)
	part2(course)
}

// Local Variables:
// compile-command: "go build"
// End:
