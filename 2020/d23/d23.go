package main

import (
	"container/ring"
	"fmt"
	"math"
)

func ToString(r *ring.Ring) string {
	result := ""

	r.Do(func(p interface{}) {
		result += fmt.Sprintf("%d ", p.(int))
	})

	return result
}

func findDestination(r *ring.Ring, min, max int) (d *ring.Ring) {
	label := r.Value.(int) - 1
	if label < min {
		label = max
	}

	for d == nil {
		f := r
		for true {
			v := f.Value.(int)
			if v == label {
				d = f
				return
			}
			f = f.Next()
			if f == r {
				break
			}
		}

		label--
		if label < min {
			label = max
		}
	}

	return
}

func play(cups []int, n int) (r *ring.Ring) {
	min := math.MaxInt64
	max := math.MinInt64
	l := len(cups)
	r = ring.New(l)

	for i := 0; i < l; i++ {
		r.Value = cups[i]
		r = r.Next()
	}

	r.Do(func(p interface{}) {
		v := p.(int)
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	})

	for i := 0; i < n; i++ {
		// remove three cups to right of current
		removed := r.Unlink(3)

		// calculate destination
		destination := findDestination(r, min, max)

		// insert cups to right of destination
		destination.Link(removed)

		// adjust current
		r = r.Next()
	}

	for true {
		if r.Value.(int) == 1 {
			break
		}

		r = r.Next()
	}

	return
}

func part1(cups []int) {
	result := ""

	r := play(cups, 100)
	r.Do(func(p interface{}) {
		result += fmt.Sprintf("%d", p.(int))
	})

	fmt.Println(result)
}

func part2(cups []int) {
}

func main() {
	cups := []int{1, 5, 6, 7, 9, 4, 8, 2, 3}

	part1(cups)
	part2(cups)
}

// Local Variables:
// compile-command: "go build"
// End:
