package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func readData() (data [][]string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		data = append(data, row)
	}

	return
}

func step(data [][]string) (result [][]string) {
	width := len(data[0])
	height := len(data)

	eresult := make([][]string, 0)
	for r := 0; r < height; r++ {
		row := make([]string, 0)
		for c := 0; c < width; c++ {
			row = append(row, data[r][c])
		}
		eresult = append(eresult, row)
	}

	// move right cukes
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			p := data[r][c]
			if p == ">" {
				c1 := (c + 1) % width
				if data[r][c1] == "." {
					eresult[r][c1] = ">"
					eresult[r][c] = "."
				}
			}
		}
	}

	result = make([][]string, 0)
	for r := 0; r < height; r++ {
		row := make([]string, 0)
		for c := 0; c < width; c++ {
			row = append(row, eresult[r][c])
		}
		result = append(result, row)
	}

	// move down cukes
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			p := eresult[r][c]
			if p == "v" {
				r1 := (r + 1) % height
				if eresult[r1][c] == "." {
					result[r1][c] = "v"
					result[r][c] = "."
				}
			}
		}
	}

	return
}

func part1(data [][]string) {
	count := 0

	for {
		newdata := step(data)
		count++
		fmt.Println(count)
		if reflect.DeepEqual(newdata, data) {
			for _, r := range data {
				fmt.Println(r)
			}
			break
		}
		data = newdata
	}

}

func part2(data [][]string) {
}

func main() {
	data := readData()
	part1(data)
	part2(data)
}

// Local Variables:
// compile-command: "go build"
// End:
