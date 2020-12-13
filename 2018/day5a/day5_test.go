package main

import "testing"

func TestReacts(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	} {
		{"a", "a", false},
		{"a", "A", true},
		{"A", "A", false},
		{"A", "a", true},
		{"a", "b", false},
		{"A", "b", false},
		{"A", "B", false},
		{"a", "B", false},
	}

	for _, c := range cases {
		got := reacts(c.a, c.b)
		if got != c.want {
			t.Errorf("react(%q, %q) == %t, want %t", c.a, c.b, got, c.want)
		}
	}
}
