package main

const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1

type Pair struct {
	L int
	M int
}

type Target struct {
	Xrange Pair
	Yrange Pair
}

/*

x(n+1) = x0 + v_x(n)
y(n+1) = y0 + v_y(n)

Assume v0_x is positive so drag is -1. But really, it has to be
since the target area is in positive x.

v_x(n+1) = v0_x - n  -- this should really be max(0, ...) but we'll just keep that in mind

So

x(n+1) = x0 + v0_x - n

Similarly,

y(n+1) = y0 + v0_y - n -- and there is no max, and we don't have to assume w/ delta is +/- 1

Now, we want to solve these simultaneously:

xmin <= x0 + x0_x - n <= xmax
ymin <= y0 + v0_y - n <= ymax

xmin - x0 <= v0_x - n <= xmax - x0
ymin - y0 <= v0_y - n <= ymax - y0

let m_x = xmin - x0, M_x = xmax - x0, etc. (just to simplify)


m_x <= v0_x - n <= Mx
m_y <= v0_y - n <= My

Except this only has to hold true for _some_ n, not for _all_ n.

... need graph paper (and I can't go downstairs until cleaners are done :( ...


.....

hmm..... max height will be reached on step v0_y and will be a height of (v0_y(v0_y + 1))/2


*/

func part1(t Target) {
	// from chicken scratches this brackets the possible initial x speeds
	for vx0 := 10; vx0 < 148; vx0++ {
	}
}

func part2() {
}

func main() {
	// sampleTarget := Target{Xrange: Pair{L: 20, M: 30}, Yrange: Pair{L: -10, M: -5}}
	puzzleTarget := Target{Xrange: Pair{L: 85, M: 145}, Yrange: Pair{L: -163, M: -108}}

	part1(puzzleTarget)
	part2()
}

// Local Variables:
// compile-command: "go build"
// End:
