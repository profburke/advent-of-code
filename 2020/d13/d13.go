package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readData() (timestamp int, buses []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	timestamp, _ = strconv.Atoi(line)

	scanner.Scan()
	line = scanner.Text()

	for _, s := range strings.Split(line, ",") {
		var bus int
		if s == "x" {
			bus = -1
		} else {
			bus, _ = strconv.Atoi(s)
		}
		buses = append(buses, bus)
	}

	return
}

func part1(timestamp int, buses []int) {
	minWait := math.MaxInt64
	busId := 0

	for _, bus := range buses {
		r := timestamp % bus
		w := bus - r

		if w < minWait {
			minWait = w
			busId = bus
		}
	}

	fmt.Println(minWait, busId, minWait*busId)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmm(ns []int) (result int) {
	if len(ns) == 1 {
		return ns[0]
	}

	result = lcm(ns[0], ns[1])
	for _, n := range ns[2:] {
		result = lcm(result, n)
	}

	return
}

// this isn't working on the test files...
// peeked at a coworker's solution, as best I can tell
// he takes the opposite approach: i.e. take the lcm of
// all the buses that do arrive at the proper time and
// add that to timestamp
func part2(timestamp int, buses []int) {
	// ignore the passed in timestamp
	//	timestamp = 100000000000000
	timestamp = 0

	for true {
		offsets := make([]int, 0)
		satisfies := true

		for i, bus := range buses {
			if bus == -1 {
				continue
			}

			r := (timestamp + i) % bus
			if r != 0 {
				satisfies = false
				s := bus - r
				offsets = append(offsets, s)
				//				break
			}
		}

		if satisfies {
			break
		}

		l := lcmm(offsets)
		fmt.Println("timestamp", timestamp, "l", l)
		timestamp += l
	}

	fmt.Println(timestamp)
}

func main() {
	timestamp, buses := readData()

	//	part1(timestamp, buses)
	part2(timestamp, buses)
}

// Local Variables:
// compile-command: "go build"
// End:
