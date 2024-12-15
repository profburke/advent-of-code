package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func canonicalize(i string) string {
	chars := strings.Split(i, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func puzzle5(u string, rd map[int][]string) int {
	return 0
}

func puzzle6(u string, rd map[int][]string) int {
	splitU := strings.Split(u, "")
	p8 := rd[8]
	cs := p8[3] // this should be the cross segment in the 8
	rs := p8[2] // this should be the upper right vertical segment

	is0 := true
	for _, s := range splitU {
		if s == cs {
			is0 = false
		}
	}

	if is0 {
		return 0
	}

	is6 := true
	for _, s := range splitU {
		if s == rs {
			is6 = false
		}
	}

	if is6 {
		return 6
	} else {
		return 9
	}
}

func part2(iomaps []IoMap) {
	display := make(map[string]rune) // this maps wire inputs to digits
	reverseDisplay := make(map[int][]string)
	unknowns := make([]string, 0)

	for _, m := range iomaps {
		for _, i := range m.Inputs {
			// check if it's a 1, 4, 7, 8
			switch len(i) {
			case 2:
				k := canonicalize(i)
				display[k] = 1
				reverseDisplay[1] = strings.Split(k, "")
			case 3:
				k := canonicalize(i)
				display[k] = 7
				reverseDisplay[7] = strings.Split(k, "")
			case 4:
				k := canonicalize(i)
				display[k] = 4
				reverseDisplay[4] = strings.Split(k, "")
			case 7:
				k := canonicalize(i)
				display[k] = 8
				reverseDisplay[8] = strings.Split(k, "")
			default:
				unknowns = append(unknowns, canonicalize(i))
			}
		}

		// now try and puzzle out the rest of the wires
		// 0, 2, 3, 4, 6, 9
		for _, u := range unknowns {
			if len(u) == 5 {
				// d := puzzle5(u, reverseDisplay)
				// display[u] = d
			} else if len(u) == 6 {
				d := puzzle6(u, reverseDisplay)
				display[u] = rune(d) + rune(0)
			} else {
				log.Fatal("oops!")
			}
		}

		for _, o := range m.Outputs {
			d, ok := display[canonicalize(o)]
			if ok {
				fmt.Print(d)
			} else {
				fmt.Print("#")
			}
		}

		fmt.Println()
	}
}

func main() {
	iomaps := readIomaps()
	part1(iomaps)
	part2(iomaps)
}

// Local Variables:
// compile-command: "go build"
// End:
