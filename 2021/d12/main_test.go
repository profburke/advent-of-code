package main

import "testing"

func TestIsSmall(t *testing.T) {
	cases := []struct {
		cave   string
		wanted bool
	}{
		{"a", true},
		{"z", true},
		{"B", false},
		{"HX", false},
		{"Z", false},
		{"alpha", true},
	}

	for _, tt := range cases {
		got := isSmall(tt.cave)
		if got != tt.wanted {
			t.Errorf("for %v: got %v wanted %v", tt.cave, got, tt.wanted)
		}
	}
}
