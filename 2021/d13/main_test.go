package main

import "testing"

func TestFoldUp(t *testing.T) {
	cases := []struct {
		p      Point
		v      int
		wanted Point
	}{
		{Point{X: 1, Y: 1}, 7, Point{X: 1, Y: 13}},
		{Point{X: 0, Y: 8}, 7, Point{X: 0, Y: 8}},
		{Point{X: 2, Y: -5}, 0, Point{X: 2, Y: 5}},
	}

	for _, tt := range cases {
		got := foldUp(tt.p, tt.v)
		if got != tt.wanted {
			t.Errorf("for %v: got %v want %v", tt.p, got, tt.wanted)
		}
	}
}

func TestFoldLeft(t *testing.T) {
	cases := []struct {
		p      Point
		v      int
		wanted Point
	}{
		{Point{X: 1, Y: 1}, 7, Point{X: 1, Y: 1}},
		{Point{X: 20, Y: 8}, 7, Point{X: -6, Y: 8}},
		{Point{X: 2, Y: -5}, 0, Point{X: -2, Y: -5}},
	}

	for _, tt := range cases {
		got := foldLeft(tt.p, tt.v)
		if got != tt.wanted {
			t.Errorf("for %v: got %v want %v", tt.p, got, tt.wanted)
		}
	}
}
