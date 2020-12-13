// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Opcode int

const (
	ADDR = 10
	ADDI = 6
	MULR = 9
	MULI = 0
	BANR = 14
	BANI = 2
	BORR = 11
	BORI = 12
	SETR = 15
	SETI = 1
	GTIR = 7
	GTRI = 3
	GTRR = 4
	EQIR = 8
	EQRI = 13
	EQRR = 5
)

func (o Opcode) String() (result string) {
	switch o {
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

type Instruction struct {
	Opcode Opcode
	A, B, C int
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

func readProgram(r io.Reader, ch chan Instruction) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var (
			a, b, c, d int
		)
		
		line := scanner.Text()
		matches, err := fmt.Sscanf(line, "%d %d %d %d", &a, &b, &c, &d)
		if err != nil || matches != 4 {
			fmt.Fprintln(os.Stderr, "error reading instruction: ", err)
			os.Exit(1)
		}
		i := Instruction{Opcode(a), b, c, d}
		ch <- i
	}
	close(ch)
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
	
	ch := make(chan Instruction, 100)
	reader := bufio.NewReader(fh)
	
	go readProgram(reader, ch)

	registers := make([]int, 4)
	count := 0
	for instruction := range ch {
		fmt.Println(fmt.Sprintf("[%d] reg: %d %d %d %d - %q", count,
			registers[0], registers[1], registers[2], registers[3],
			instruction))
		execute(registers, instruction)
		count++
	}
	fmt.Println("registers: ", registers)
}
