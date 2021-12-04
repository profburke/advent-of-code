package main

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 7, b: 5, want: 1},
		{a: 2, b: 8, want: 2},
		{a: 99, b: 55, want: 11},
	}

	for _, tt := range tests {
		got := gcd(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("gcd(%d, %d) = %d -- got: %d", tt.a, tt.b, tt.want, got)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 7, b: 5, want: 35},
		{a: 2, b: 8, want: 8},
		{a: 99, b: 55, want: 495},
	}

	for _, tt := range tests {
		got := lcm(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("lcm(%d, %d) = %d -- got: %d", tt.a, tt.b, tt.want, got)
		}
	}
}

func TestLCMM(t *testing.T) {
	tests := []struct {
		as   []int
		want int
	}{
		{as: []int{3, 5, 7}, want: 105},
	}

	for _, tt := range tests {
		got := lcmm(tt.as)
		if got != tt.want {
			t.Errorf("lcmm(%v) = %d -- got: %d", tt.as, tt.want, got)
		}
	}
}
