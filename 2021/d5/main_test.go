package main

import "testing"

func TestIsHorizontal(t *testing.T) {
	cases := []struct {
		line   Line
		wanted bool
	}{
		{Line{E1: Point{X: 2, Y: 10}, E2: Point{X: 4, Y: 10}}, true},
		{Line{E1: Point{X: 14, Y: 0}, E2: Point{X: -1, Y: 0}}, true},
		{Line{E1: Point{X: 5, Y: 1}, E2: Point{X: 5, Y: 10}}, false},
		{Line{E1: Point{X: 1, Y: 1}, E2: Point{X: -7, Y: -7}}, false},
		{Line{E1: Point{X: 2, Y: 3}, E2: Point{X: 4, Y: 5}}, false},
	}

	for _, tt := range cases {
		got := isHorizontal(tt.line)
		if got != tt.wanted {
			t.Errorf("for %v: got %v want %v", tt.line, got, tt.wanted)
		}
	}
}

func TestIsVertical(t *testing.T) {
	cases := []struct {
		line   Line
		wanted bool
	}{
		{Line{E1: Point{X: 10, Y: 2}, E2: Point{X: 10, Y: 4}}, true},
		{Line{E1: Point{X: 0, Y: 15}, E2: Point{X: 0, Y: -1}}, true},
		{Line{E1: Point{X: 1, Y: 5}, E2: Point{X: 10, Y: 5}}, false},
		{Line{E1: Point{X: 1, Y: 1}, E2: Point{X: -7, Y: -7}}, false},
		{Line{E1: Point{X: 2, Y: 3}, E2: Point{X: 4, Y: 5}}, false},
	}

	for _, tt := range cases {
		got := isVertical(tt.line)
		if got != tt.wanted {
			t.Errorf("for %v: got %v want %v", tt.line, got, tt.wanted)
		}
	}
}
