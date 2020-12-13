package main

import "testing"

func TestManhattan(t *testing.T) {
	cases := []struct {
		a, b Point
		want int
	} {
		{ Point{1, 10, 10}, Point{2, 10, 10}, 0},
		{ Point{1, 10, 10}, Point{2, 10, 20}, 10},
		{ Point{1, 10, 10}, Point{2, 20, 10}, 10},
		{ Point{1, 10, 10}, Point{2, 20, 20}, 20},
		{ Point{1, -10, -20}, Point{2, 20, 30}, 80},
	}

	for _, c := range cases {
		got := c.a.manhattanDistance(c.b)
		if got != c.want {
			t.Errorf("%q.distance(%q) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
