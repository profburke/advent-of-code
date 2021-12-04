package main

import (
	"errors"
	"testing"
)

func TestFound(t *testing.T) {
	tests := []struct {
		want   bool
		list   []int
		target int
	}{
		{true, []int{1, 4, 9, 10}, 4},
		{false, []int{1, 4, 9, 10}, 5},
	}

	for _, tt := range tests {
		got := found(tt.list, tt.target)
		if got != tt.want {
			t.Errorf("looking for %d: got %v, want %v", tt.target, got, tt.want)
		}
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		want   error
		list   []int
		target int
	}{
		{nil, []int{4, 1, 9, 10, 3}, 7},
		{nil, []int{4, 1, 9, 10, 3}, 13},
		{errors.New(""), []int{4, 1, 9, 10, 3}, 23},
	}

	for _, tt := range tests {
		_, _, got := twosum(tt.target, tt.list)
		if (got == nil) != (tt.want == nil) {
			t.Errorf("looking for %d: got %v, want %v", tt.target, got, tt.want)
		}
	}
}
