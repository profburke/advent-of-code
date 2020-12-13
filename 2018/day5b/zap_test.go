package main

import "testing"

func TestZap(t *testing.T) {
	cases := []struct {
		before, after, l string
	} {
		{"aaaa", "", "a"},
		{"Alpha", "lph", "a"},
		{"Alpha", "lph", "A"},
	}

	for _, c := range cases {
		got := zap(c.before, c.l)
		if got != c.after {
			t.Errorf("zap(%q, %q) == %q, want %q", c.before, c.l, got, c.after)
		}
	}
}
