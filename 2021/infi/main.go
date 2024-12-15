package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PartSpec struct {
	Quantity int
	Name     string
}

type ToySpec struct {
	Name  string
	Parts []PartSpec
}

func readSpecs() (missing int, toys []ToySpec) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	missing, _ = strconv.Atoi(parts[0])

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		rawspecs := strings.Split(parts[1], ", ")
		specs := make([]PartSpec, 0)
		for _, spec := range rawspecs {
			fields := strings.Fields(spec)
			n, _ := strconv.Atoi(fields[0])
			specs = append(specs, PartSpec{Quantity: n, Name: fields[1]})
		}

		toys = append(toys, ToySpec{Name: parts[0], Parts: specs})
	}

	return
}

func partCount(toy ToySpec, tp map[string]ToySpec) int {
	total := 0

	for _, spec := range toy.Parts {
		t, ok := tp[spec.Name]
		if ok {
			total += spec.Quantity * partCount(t, tp)
		} else {
			total += spec.Quantity
		}
	}

	return total
}

func part1(toys []ToySpec) (partsPerToy map[string]int) {
	partsPerToy = make(map[string]int)
	max := 0
	var toy string

	toysWithParts := make(map[string]ToySpec)
	for _, spec := range toys {
		toysWithParts[spec.Name] = spec
	}

	for _, spec := range toys {
		nparts := partCount(spec, toysWithParts)
		partsPerToy[spec.Name] = nparts
		fmt.Println(spec.Name, nparts)
		if nparts > max {
			max = nparts
			toy = spec.Name
		}
	}

	fmt.Println("MAX:", toy, max)

	return
}

/*
For the second part we need to check the partitions of 20 where
we are partitioning it into at most 8 bins (there
are 8 toys that are not parts of other toys). we check
a partition to see if sum_n(b_n * p_n) = 446383.

b_n is the nth bin and
p_n is the number of parts that make up toy_n

The algorithm generates unique partitions so we than need to
check each permutation of toys (i.e. assign toy1 to bin 1 ...
then assign toy2 to bin 1 ... etc

*/

// func part2(missing int, toys []ToySpec, ppt map[string]int) {

// 	var p [8]int
// 	k := 0 // index of last element in partition
// 	p[k] = 20

// 	for {
// 		// check current partition
// 		fmt.Println(p)

// 		remVal := 0
// 		for k >= 0 && p[k] == 1 {
// 			remVal += p[k]
// 			k--
// 		}

// 		if k < 0 {
// 			break
// 		}

// 		p[k]--
// 		remVal++

// 		for remVal > p[k] {
// 			p[k+1] = p[k]
// 			remVal = remVal - p[k]
// 			k++
// 		}

// 		p[k+1] = remVal
// 		k++
// 	}
// }

/*

I don't think partitions is the right way to go. Let's try
making "change".

*/

type Coin struct {
	Name string // not needed in change algo, but we need the names fo the solution
	// to this problem
	Value int // how many parts make up top; i.e. the "denomination"
}

func findCombos(coins []Coin, int amount, used []int) bool {
	if amount == 0 {
		// print solution
		return true
	} else {
		if findCombos(coins, amount-coins[0].Value, used) {
			return true
		} else {
			return findCombos(coins[1:])
		}
	}
}

func part2(amount int, toys []ToySpec, ppt map[string]int) {
	// missing is the number we want to sum to
	// denoms is a set of the possible "coins"

	// Short enough I'll hard code...
	coins := []Coin{
		{"Lightsaber", 6411},
		{"HandheldComputer", 22044},
		{"ElectrischeRacebaan", 9708},
		{"QuadDrone", 45222},
		{"PikachuPlushy", 2173},
		{"Trampoline", 3367},
		{"BatmobileReplica", 16370},
		{"DanceDanceRevolutionMat", 6649},
	}

	findCombos(coins, amount, []int{0, 0, 0, 0, 0, 0, 0, 0})
}

func main() {
	missing, toys := readSpecs()
	ppt := part1(toys)
	part2(missing, toys, ppt)
}

// Local Variables:
// compile-command: "go build"
// End:
