package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readData() (data []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		data = append(data, i)
	}

	return data
}

func found(list []int, item int) bool {
	index := sort.SearchInts(list, item)
	if list[index] == item {
		return true
	} else {
		return false
	}
}

func part1(expenses []int) {
	for _, val := range expenses {
		other := 2020 - val

		if found(expenses, other) {
			fmt.Fprintln(os.Stdout, val, "*", other, "=", val*other)
			return
		}
	}

	fmt.Fprintln(os.Stdout, "Couldn't find a match")
}

func copyWithout(s []int, i int) []int {
	result := make([]int, len(s))
	copy(result, s)
	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]

	return s
}

func part2(expenses []int) {
	for i, val1 := range expenses {
		target := 2020 - val1
		var reducedExpenses = copyWithout(expenses, i)

		for _, val := range reducedExpenses {
			other := target - val

			if found(reducedExpenses, other) {
				fmt.Fprintln(os.Stdout, val1, "*", val, "*", other, "=",
					val1*val*other)
				return
			}
		}
	}

	fmt.Fprintln(os.Stdout, "Couldn't find a match")
}

func main() {
	expenses := readData()
	sort.Ints(expenses)

	part1(expenses)
	part2(expenses)
}

// Local Variables:
// compile-command: "go build"
// End:
