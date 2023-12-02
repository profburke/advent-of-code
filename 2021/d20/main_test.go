package main

import "testing"

func TestGetIndex(t *testing.T) {
	cases := []struct {
		pattern string
		want    int
	}{
		{"111111111", 511},
		{"000001010", 10},
	}

	for _, tt := range cases {
		got := getIndex(tt.pattern)
		if got != tt.want {
			t.Errorf("for %s: got %d want %d", tt.pattern, got, tt.want)
		}
	}
}
