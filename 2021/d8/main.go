package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IoMap struct {
	Inputs  []string
	Outputs []string
}

func readIomaps() (iomaps []IoMap) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		inputs := strings.Fields(parts[0])
		outputs := strings.Fields(parts[1])

		iomaps = append(iomaps, IoMap{Inputs: inputs, Outputs: outputs})
	}

	return iomaps
}

func part1(iomaps []IoMap) {
	count := 0

	for _, m := range iomaps {
		for _, o := range m.Outputs {
			switch len(o) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	fmt.Println(count)
}

func part2(iomaps []IoMap) {
}

func main() {
	iomaps := readIomaps()
	part1(iomaps)
	part2(iomaps)
}

// Local Variables:
// compile-command: "go build"
// End:
