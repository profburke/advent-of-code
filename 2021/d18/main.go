package main

import (
	"bufio"
	"fmt"
	"os"
)

type SnailNumber struct {
	Left  int
	Right int
}

func (l SnailNumber) add(r SnailNumber) SnailNumber {
	return SnailNumber{Left: 5, Right: 10}
}

func (n SnailNumber) reduce() SnailNumber {
	return SnailNumber{Left: 2, Right: 4}
}

func (n SnailNumber) magnitude() int {
	return 7
}

func readSnailNumbers() (numbers []SnailNumber) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		//line := scanner.Text()
		numbers = append(numbers, SnailNumber{})
	}

	return
}

func part1(numbers []SnailNumber) {
	sum := numbers[0]

	for _, n := range numbers[1:] {
		sum = sum.add(n)
		sum = sum.reduce()
	}

	fmt.Println(sum.magnitude())
}

func part2() {
}

func main() {
	numbers := readSnailNumbers()
	part1(numbers)
	part2()
}

// Local Variables:
// compile-command: "go build"
// End:
