package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var memo map[int]int

func init() {
	memo = make(map[int]int)
}

func readData() (adaptors []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		adaptor, _ := strconv.Atoi(line)
		adaptors = append(adaptors, adaptor)
	}

	return
}

func removeAt(i int, list []int) (newList []int) {
	newList = make([]int, len(list))
	copy(newList, list)
	copy(newList[i:], newList[i+1:])
	newList = newList[:len(newList)-1]

	return
}

func part1(adaptors []int) {
	currentJoltage := 0
	oneJoltDifferences := 0
	threeJoltDifferences := 0

	for true {
		if len(adaptors) == 0 {
			break
		}

		sort.Ints(adaptors)

		joltage := adaptors[0]
		adaptors = removeAt(0, adaptors)

		difference := joltage - currentJoltage
		if difference == 1 {
			oneJoltDifferences++
		} else if difference == 3 {
			threeJoltDifferences++
		} // difference could be 2 but we ignore it

		currentJoltage = joltage
	}

	threeJoltDifferences++ // for your device
	fmt.Println("Part 1 =", oneJoltDifferences*threeJoltDifferences)
}

func inRange(target, n int) bool {
	difference := target - n

	return (difference >= 1 && difference <= 3)
}

// assuming each adaptor is unique
func dp(current int, adaptors []int) (count int) {
	var found bool
	if count, found = memo[current]; found {
		return
	}

	// if len(adaptors) == 0 {
	// 	// since there aren't any more adaptors then these
	// 	// all plug directly into the socket
	// 	if current == 1 || current == 2 || current == 3 {
	// 		fmt.Println("base case")
	// 		count = 1
	// 	} else { // shouldn't happen
	// 		// TODO: throw an error
	// 		count = 0
	// 	}

	// 	memo[current] = 1
	// 	return
	// }

	count = 0

	if current == 3 || current == 2 || current == 1 {
		count++
	}

	if len(adaptors) > 0 && inRange(current, adaptors[0]) {
		count += dp(adaptors[0], adaptors[1:])
	}

	if len(adaptors) > 1 && inRange(current, adaptors[1]) {
		count += dp(adaptors[1], adaptors[2:])
	}

	if len(adaptors) > 2 && inRange(current, adaptors[2]) {
		count += dp(adaptors[2], adaptors[3:])
	}

	memo[current] = count
	return
}

func part2(adaptors []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(adaptors)))
	device := adaptors[0] + 3

	count := dp(device, adaptors)
	fmt.Println("Part 2 =", count)
}

func main() {
	adaptors := readData()

	part1(adaptors)
	part2(adaptors)
}

// Local Variables:
// compile-command: "go build"
// End:
