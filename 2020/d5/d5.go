package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Seating struct {
	Upper    int
	Occupied map[int]bool
}

func readData() (passes []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		pass := scanner.Text()
		passes = append(passes, pass)
	}

	return
}

func encodingToInt(encoding string) int {
	upper := int(math.Pow(2, float64(len(encoding))))
	lower := 0

	for _, c := range encoding {
		if c == 'B' {
			lower += (upper - lower) / 2
		} else {
			upper = lower + (upper-lower)/2
		}
	}

	return lower
}

func findSeatId(pass string) int {
	row := 0
	seat := 0

	rowEncoding := pass[0:7]
	row = encodingToInt(rowEncoding)

	seatEncoding := pass[7:10]
	seatEncoding = strings.ReplaceAll(seatEncoding, "R", "B")
	seatEncoding = strings.ReplaceAll(seatEncoding, "L", "F")
	seat = encodingToInt(seatEncoding)

	return row*8 + seat
}

func part1(passes []string) (seating Seating) {
	var seatIds []int
	seating.Occupied = make(map[int]bool)

	for _, pass := range passes {
		seatId := findSeatId(pass)
		seating.Occupied[seatId] = true
		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)

	upper := seatIds[len(seatIds)-1]
	seating.Upper = upper
	fmt.Println("Part 1 =", upper)

	return
}

func part2(seating Seating) {
	prev := -1
	current := 0
	next := 1

	for current = 0; current < seating.Upper; prev, current, next = prev+1, current+1, next+1 {
		if seating.Occupied[prev] && !seating.Occupied[current] && seating.Occupied[next] {
			break
		}
	}

	fmt.Println("Part 2 =", current)
}

func main() {
	passes := readData()

	seating := part1(passes)
	part2(seating)
}

// Local Variables:
// compile-command: "go build"
// End:
