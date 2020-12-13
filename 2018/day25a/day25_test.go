package main

import "testing"

func TestD(t *testing.T) {
	cases := []struct {
		a, b Point
		want int
	} {
		{Point{-1, 2, 2, 0}, Point{0, 0, 2, -2}, 5},
	}

	for _, c := range cases {
		got := c.a.distance(c.b)
		if got != c.want {
			t.Errorf("%q distance to %q == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
