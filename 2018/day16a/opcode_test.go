package main

import "testing"

func TestExecute(t *testing.T) {
	cases := []struct {
		start, want []int
		op Instruction
	} {
		{ []int{4, 6, 0, 0}, []int{4, 6, 0, 10}, Instruction{ADDR, 0, 1, 3} },
		{ []int{4, 6, 0, 0}, []int{4, 6, 0, 10}, Instruction{ADDI, 0, 17, 3} },
		
	}

	for _, c := range cases {
		registers := c.start
		execute(registers, c.op)
		if !match(registers, c.want) {
			t.Errorf("got %q, want %q", registers, c.want)
		}
	}
}
