// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"errors"
	"io"
        "fmt"
	"math"
	"os"
	"strings"
)

func gridParse(initial string) (grid map[int]string, err error) {
	const initialLabel = "initial state: "
	if !strings.HasPrefix(initial, initialLabel) {
		return nil, errors.New("malformed initial state")
	}

	grid = make(map[int]string)
	initial = initial[len(initialLabel):]

	for i, r := range initial {
		if string(r) == "#" {
			grid[i] = "#"
		} else {
			grid[i] = "."
		}
	}
	
	return grid, nil
}

func patternToInt(pattern string) (result int, err error) {
	if len(pattern) != 5 {
		return -1, errors.New("error parsing rule pattern")
	}

	for _, r := range pattern {
		if string(r) == "#" {
			result *= 2
			result += 1
		} else {
			result *= 2
		}
	}

	return result, nil
}

func parseRule(rules map[int]string, line string) (err error) {
	pieces := strings.Split(line, " ")
	if len(pieces) != 3 { return errors.New("malformed rule: '" + line + "'") }

	index, err := patternToInt(pieces[0])
	if err != nil {
		return err
	}

	value := "."
	if pieces[2] == "#" {
		value = "#"
	}

	rules[index] = value
	
	return nil
}

func readSetup(r io.Reader) (grid map[int]string, rules map[int]string, err error) {
	scanner := bufio.NewScanner(r)
	rules = make(map[int]string)
	
	scanner.Scan()
	initial := scanner.Text()
	grid, err = gridParse(initial)
	if err != nil {
		return nil, nil, err
	}

	scanner.Scan() // skip blank line
	
	for scanner.Scan() {
		line := scanner.Text()
		if err := parseRule(rules, line); err != nil {
			return nil, nil, err
		}
	}

	err = scanner.Err()
	return grid, rules, err
}

func gridString(grid map[int]string, left, right int) (result string) {
	for i := left; i <= right; i++ {
		result += grid[i]
	}
	return result
}

func getNeighborhood(grid map[int]string, center, left, right int) (neighborhood string) {
	lower := center - 2
	upper := center + 2
	for i := lower; i <= upper; i++ {
		neighborhood += grid[i]
	}
	if center == left {
		neighborhood = ".." + neighborhood
	} else if center == left + 1 {
		neighborhood = "." + neighborhood
	} else if center == right - 1 {
		neighborhood = neighborhood + "."
	} else if center == right {
		neighborhood = neighborhood + ".."
	}
	return neighborhood
}

func main() {
	grid, rules, err := readSetup(os.Stdin)
	if err != nil {
		fmt.Println("error reading input: ", err)
		os.Exit(1)
	}
	
	generation := 0
	left := 0
	right := len(grid) - 1

	// starting with generation 100, each generation is exactly the same as
	// previous, except shifted one spot to the right
	for generation < 101 {
		// gs := gridString(grid, left, right)
		// firstPot := math.MinInt64
		// for i := left; i <= right; i++ {
		// 	if grid[i] == "#" {
		// 		firstPot = i
		// 		break
		// 	}
		// }
		
		// fmt.Println(fmt.Sprintf("[%d - %d - %d] %s", generation, left, firstPot, gs))
		// sum := 0
		// for i := range grid {
		// 	if grid[i] == "#" {
		// 		sum += i
		// 	}
		// }
		//fmt.Println("sum: ", sum)
		// history = append(history, HistoricalRecord{gs, sum})


		grid[left - 1] = "."
		grid[left - 2] = "."
		grid[right + 1] = "."
		grid[right + 2] = "."
		left -= 2
		right += 2

		newGrid := make(map[int]string)
		for i := left; i <= right; i++ {
			neighborhood := getNeighborhood(grid, i, left, right)
			ri, err :=patternToInt(neighborhood)

			if err != nil {
				fmt.Printf("error while evolving: ", err)
				os.Exit(1)
			}
			newGrid[i] = rules[ri]
		}

		// only trim on the right...
		i := right
		for newGrid[i] == "." {
			right--
			delete(newGrid, i)
			i--
		}
		
		grid = newGrid
		generation++
	}
	sum := 0
	count := 0
	for i := range grid {
	 	if grid[i] == "#" {
			sum += i
			count++
		}
	}

	firstPot := math.MinInt64
	for i := left; i <= right; i++ {
		if grid[i] == "#" {
			firstPot = i
			break
		}
	}
	
	fmt.Println(fmt.Sprintf("sum: %d, first: %d, count: %d", sum, firstPot, count))
	// first adjust the sum as if the first plant was at spot 0
	sum -= (count * firstPot)
	fmt.Println("sum: ", sum)
	// now the first occupied position at generation n is in pot n - 11
	// so we adjust back to generation 50000000000
	sum += ((50000000000 - 11) * count)
	fmt.Println(sum)
}
