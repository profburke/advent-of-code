package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func readData() (data []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		data = append(data, n)
	}

	return
}

func found(list []int, item int) bool {
	index := sort.SearchInts(list, item)

	if index == len(list) {
		return false
	}

	if list[index] == item {
		return true
	} else {
		return false
	}
}

func twosum(target int, preamble []int) (x, y int, err error) {
	list := make([]int, len(preamble))
	copy(list, preamble)

	sort.Ints(list)

	for _, val := range list {
		other := target - val

		if found(list, other) {
			x, y = val, other
			return
		}
	}

	err = errors.New("Couldn't find a pair")
	return
}

func isValid(target int, preamble []int) bool {
	_, _, err := twosum(target, preamble)
	return (err == nil)
}

func part1(data []int, size int) (badwolf int, err error) {
	preamble := data[:size]
	data = data[size:]

	for true {
		target := data[0]
		if !isValid(target, preamble) {
			fmt.Println("Part 1 =", target)

			badwolf = target
			return
		}
		preamble = preamble[1:]
		preamble = append(preamble, data[0])
		data = data[1:]

		if len(data) == 0 {
			break
		}
	}

	err = errors.New("every item was valid")
	return
}

func sum(low, high int, data []int) (s, min, max int) {
	min = math.MaxInt64
	max = math.MinInt64

	for index := low; index <= high; index++ {
		if data[index] < min {
			min = data[index]
		}

		if data[index] > max {
			max = data[index]
		}

		s += data[index]
	}

	return
}

func part2(target int, data []int) {
	var low, high int

	for low = 0; low < len(data); low++ {
		for high = low + 1; high < len(data); high++ {
			s, min, max := sum(low, high, data)

			if s == target {
				fmt.Println("Part 2:", min+max)
				return
			} else if s > target {
				break
			}
		}
	}
	fmt.Println("Something's wrong...")
}

func main() {
	data := readData()

	badwolf, _ := part1(data, 25)
	part2(badwolf, data)
}

// Local Variables:
// compile-command: "go build"
// End:
