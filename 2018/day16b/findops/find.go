// -*- compile-command: "go build"; -*-
package main

// Ought to be able to do this programatically, but I did it by hand...so sue me...

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Opcode int

type Instruction struct {
	Opcode Opcode
	A, B, C int
}

const (
	UNKNOWN Opcode = iota
	ADDR
	ADDI
	MULR
	MULI
	BANR
	BANI
	BORR
	BORI
	SETR
	SETI
	GTIR
	GTRI
	GTRR
	EQIR
	EQRI
	EQRR
)

func (o Opcode) String() (result string) {
	switch o {
	case UNKNOWN:
		result = "UNKN"
	case ADDR:
		result = "ADDR"
	case ADDI:
		result = "ADDI"
	case MULR:
		result = "MULR"
	case MULI:
		result = "MULI"
	case BANR:
		result = "BANR"
	case BANI:
		result = "BANI"
	case BORR:
		result = "BORR"
	case BORI:
		result = "BORI"
	case SETR:
		result = "SETR"
	case SETI:
		result = "SETI"
	case GTIR:
		result = "GTIR"
	case GTRI:
		result = "GTRI"
	case GTRR:
		result = "GTRR"
	case EQIR:
		result = "EQIR"
	case EQRI:
		result = "EQRI"
	case EQRR:
		result = "EQRR"
	}

	return
}

var AllOps = []Opcode {
	ADDR, ADDI, MULR, MULI, BANR, BANI, BORR, BORI,
	SETR, SETI, GTIR, GTRI, GTRR, EQIR, EQRI, EQRR,
}

func execute(registers []int, op Instruction) {
	switch op.Opcode {
	case ADDR:
		registers[op.C] = registers[op.A] + registers[op.B]
	case ADDI:
		registers[op.C] = registers[op.A] + op.B
	case MULR:
		registers[op.C] = registers[op.A] * registers[op.B]
	case MULI:
		registers[op.C] = registers[op.A] * op.B
	case BANR:
		registers[op.C] = registers[op.A] & registers[op.B]
	case BANI:
		registers[op.C] = registers[op.A] & op.B
	case BORR:
		registers[op.C] = registers[op.A] | registers[op.B]
	case BORI:
		registers[op.C] = registers[op.A] | op.B
	case SETR:
		registers[op.C] = registers[op.A]
	case SETI:
		registers[op.C] = op.A
	case GTIR:
		if op.A > registers[op.B] {
			registers[op.C] = 1
		} else {
			registers[op.C] = 0
		}
	case GTRI:
		if registers[op.A] > op.B {
			registers[op.C] = 1
		} else {
			registers[op.C] = 0
		}
	case GTRR:
		if registers[op.A] > registers[op.B] {
			registers[op.C] = 1
		} else {
			registers[op.C] = 0
		}
	case EQIR:
		if op.A == registers[op.B] {
			registers[op.C] = 1
		} else {
			registers[op.C] = 0
		}
	case EQRI:
		if registers[op.A] == op.B {
			registers[op.C] = 1
		} else {
			registers[op.C] = 0
		}
	case EQRR:
		if registers[op.A] == registers[op.B] {
			registers[op.C] = 1
		} else {
			registers[op.C] = 0
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid opcode")
		os.Exit(1)
	}
}

type Sample struct {
	Before, After []int
	Op Instruction
}

func readSamples(r io.Reader, ch chan Sample) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var (
			a, b, c, d int
		)
		
		line := scanner.Text()
		matches, err := fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &a, &b, &c, &d)
		if err != nil || matches != 4 {
			fmt.Fprintln(os.Stderr, "error reading samples: ", err)
			os.Exit(1)
		}
		before := []int{a, b, c, d}
		scanner.Scan()
		line = scanner.Text()
		matches, err = fmt.Sscanf(line, "%d %d %d %d", &a, &b, &c, &d)
		if err != nil || matches != 4 {
			fmt.Fprintln(os.Stderr, "error reading samples: ", err)
			os.Exit(1)
		}
		instruction := Instruction{Opcode(a), b, c, d}
		scanner.Scan()
		line = scanner.Text()
		matches, err = fmt.Sscanf(line, "After:  [%d, %d, %d, %d]", &a, &b, &c, &d)
		if err != nil || matches != 4 {
			fmt.Fprintln(os.Stderr, "error reading samples: ", err)
			os.Exit(1)
		}
		after := []int{a, b, c, d}
		scanner.Scan() // blank line, discard

		sample := Sample{}
		sample.Before = before
		sample.Op = instruction
		sample.After = after
		ch <- sample
	}
	close(ch)
}

func match(a []int, b []int) bool {
	if len(a) != len(b) { return false }

	for i := range a {
		if a[i] != b[i] { return false }
	}
	return true
}

func testSample(sample Sample) (candidates []Opcode) {
	for _, op := range AllOps {
		registers := make([]int, len(sample.Before))
		for i, v := range sample.Before {
			registers[i] = v
		}
		i := sample.Op
		i.Opcode = op
		execute(registers, i)
		if match(sample.After, registers) {
			candidates = append(candidates, op)
		}
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <filename>", os.Args[0])
		os.Exit(1)
	}

	fh, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file: ", err)
		os.Exit(1)
	}
	defer fh.Close()
	
	ch := make(chan Sample, 100)
	reader := bufio.NewReader(fh)
	
	go readSamples(reader, ch)

	possibilities := make(map[int][]Opcode)
	for i := 0; i < 16; i++ {
		possibilities[i] = make([]Opcode, 0)
	}
	
	for sample := range ch {
		candidates := testSample(sample)
		i := int(sample.Op.Opcode)
		for _, c := range candidates {
			if !contains(possibilities[i], c) {
				possibilities[i] = append(possibilities[i], c)
			}
		}
	}

	for i := 0; i < 16; i++ {
		fmt.Println(fmt.Sprintf("[%d] %q", i, possibilities[i]))
	}
	
}

func contains(set []Opcode, e Opcode) bool {
	for _, v := range set {
		if v == e { return true }
	}
	return false
}
