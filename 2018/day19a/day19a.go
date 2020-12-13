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

func parseOpcode(m string) (result Opcode) {
	switch m {
	case "addr":
		result = ADDR
	case "addi":
		result = ADDI
	case "mulr":
		result = MULR
	case "muli":
		result = MULI
	case "banr":
		result = BANR
	case "bani":
		result = BANI
	case "borr":
		result = BORR
	case "bori":
		result = BORI
	case "setr":
		result = SETR
	case "seti":
		result = SETI
	case "gtir":
		result = GTIR
	case "gtri":
		result = GTRI
	case "gtrr":
		result = GTRR
	case "eqir":
		result = EQIR
	case "eqri":
		result = EQRI
	case "eqrr":
		result = EQRR
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

func readProgram(r io.Reader) (code []Instruction, boundRegister int) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	line := scanner.Text()
	matches, err := fmt.Sscanf(line, "#ip %d", &boundRegister)
	if err != nil || matches != 1 {
		fmt.Fprintln(os.Stderr, "error reading ip binding: ", err)
		os.Exit(1)
	}
	
	for scanner.Scan() {
		var (
			a string
			b, c, d int
		)
		
		line := scanner.Text()
		matches, err := fmt.Sscanf(line, "%s %d %d %d", &a, &b, &c, &d)
		if err != nil || matches != 4 {
			fmt.Fprintln(os.Stderr, "error reading instruction: ", err)
			os.Exit(1)
		}
		i := Instruction{parseOpcode(a), b, c, d}
		code = append(code, i)
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
	
	reader := bufio.NewReader(fh)
	code, boundRegister := readProgram(reader)

	registers := make([]int, 6)
	ip := 0
	
	for ip < len(code) && ip > -1 {
		i := code[ip]
		registers[boundRegister] = ip
		execute(registers, i)
		ip = registers[boundRegister]
		ip++
	}
	fmt.Println("registers: ", registers)
}
