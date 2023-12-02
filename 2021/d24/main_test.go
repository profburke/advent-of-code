package main

import (
	"reflect"
	"testing"
)

func TestMakeInputs(t *testing.T) {
	got := makeInputs(12345678901234)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

type TestCase struct {
	input    []int
	variable string
	wanted   int
}

func doTests(t *testing.T, cases []TestCase, instructions []Instruction) {
	for _, tt := range cases {
		state := run(tt.input, instructions)
		got := state[tt.variable]

		if got != tt.wanted {
			t.Errorf("got %d, wanted %d", got, tt.wanted)
		}
	}
}

func TestInp(t *testing.T) {
	instructions := []Instruction{
		Instruction{"inp", "z", "", 0, false},
		Instruction{"inp", "y", "", 0, false},
		Instruction{"inp", "x", "", 0, false},
		Instruction{"inp", "w", "", 0, false},
	}

	cases := []TestCase{
		{[]int{1, 2, 3, 4}, "w", 4},
		{[]int{1, 2, 3, 4}, "x", 3},
		{[]int{1, 2, 3, 4}, "y", 2},
		{[]int{1, 2, 3, 4}, "z", 1},
	}

	doTests(t, cases, instructions)
}

func TestAdd(t *testing.T) {
	instructions := []Instruction{
		Instruction{"inp", "x", "", 0, false},
		Instruction{"inp", "y", "", 0, false},
		Instruction{"add", "y", "x", 0, true},
		Instruction{"add", "x", "", 17, false},
	}

	cases := []TestCase{
		{[]int{1, 2, 3, 4}, "x", 18},
		{[]int{1, 2, 3, 4}, "y", 3},
	}

	doTests(t, cases, instructions)
}

func TestMul(t *testing.T) {
	instructions := []Instruction{
		Instruction{"inp", "x", "", 0, false},
		Instruction{"inp", "y", "", 0, false},
		Instruction{"mul", "y", "x", 0, true},
		Instruction{"mul", "x", "", 17, false},
	}

	cases := []TestCase{
		{[]int{2, 3, 4}, "x", 34},
		{[]int{2, 3, 4}, "y", 6},
	}

	doTests(t, cases, instructions)
}

func TestDiv(t *testing.T) {
	instructions := []Instruction{
		Instruction{"inp", "x", "", 0, false},
		Instruction{"inp", "y", "", 0, false},
		Instruction{"inp", "z", "", 0, false},
		Instruction{"inp", "w", "", 0, false},
		Instruction{"div", "x", "y", 0, true},
		Instruction{"div", "z", "w", 0, true},
	}

	cases := []TestCase{
		{[]int{15, 5, 7, 3}, "x", 3},
		{[]int{15, 5, 7, 3}, "z", 2},
	}

	doTests(t, cases, instructions)
}

func TestMod(t *testing.T) {
	instructions := []Instruction{
		Instruction{"inp", "x", "", 0, false},
		Instruction{"inp", "y", "", 0, false},
		Instruction{"inp", "z", "", 0, false},
		Instruction{"mod", "y", "x", 0, true},
		Instruction{"mod", "z", "", 3, false},
	}

	cases := []TestCase{
		{[]int{2, 7, 10}, "y", 1},
		{[]int{2, 7, 10}, "z", 1},
	}

	doTests(t, cases, instructions)
}

func TestEql(t *testing.T) {
	instructions := []Instruction{
		Instruction{"inp", "x", "", 0, false},
		Instruction{"inp", "y", "", 0, false},
		Instruction{"inp", "z", "", 0, false},
		Instruction{"eql", "y", "x", 0, true},
		Instruction{"eql", "z", "x", 0, true},
		Instruction{"eql", "w", "", 11, false},
	}

	cases := []TestCase{
		{[]int{10, 17, 10}, "y", 0},
		{[]int{10, 17, 10}, "z", 1},
		{[]int{10, 17, 10}, "w", 0},
	}

	doTests(t, cases, instructions)
}

func TestP2(t *testing.T) {
	instructions := []Instruction{
		Instruction{"inp", "z", "", 0, false},
		Instruction{"inp", "x", "", 0, false},
		Instruction{"mul", "z", "", 3, false},
		Instruction{"eql", "z", "x", 0, true},
	}

	cases := []TestCase{
		{[]int{1, 3}, "z", 1},
		{[]int{1, 4}, "z", 0},
		{[]int{2, 5}, "z", 0},
	}

	doTests(t, cases, instructions)
}

func TestP1(t *testing.T) {
	instructions := []Instruction{
		Instruction{"inp", "x", "", 0, false},
		Instruction{"mul", "x", "", -1, false},
	}

	cases := []TestCase{
		{[]int{12}, "x", -12},
		{[]int{100}, "x", -100},
		{[]int{-3}, "x", 3},
	}

	doTests(t, cases, instructions)
}
