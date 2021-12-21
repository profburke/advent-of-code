package main

import "fmt"

func advance(current, steps int) int {
	end := (current + steps) % 10
	if end == 0 {
		end = 10
	}

	return end
}

func part1(positions []int) {
	scores := []int{0, 0}
	current := 0
	dieUp := 1
	rolls := 0

	for scores[0] < 1000 && scores[1] < 1000 {
		steps := 3*dieUp + 3
		dieUp += 3
		positions[current] = advance(positions[current], steps)
		scores[current] += positions[current]
		rolls += 3
		current = 1 - current
	}

	fmt.Println(rolls * scores[current])
}

func part2(depths []int) {
}

func main() {
	// starts := []int{4, 8} -- sample input
	starts := []int{10, 9}
	part1(starts)
	//	part2(depths)
}

// Local Variables:
// compile-command: "go build"
// End:
