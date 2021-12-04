package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Opcode  string
	Operand int
}

func readData() (program []Instruction) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		i := Instruction{}
		i.Opcode = parts[0]
		operand, _ := strconv.Atoi(parts[1])
		i.Operand = operand

		program = append(program, i)
	}

	return
}

func run(program []Instruction) (acc int, reason string) {
	pc := 0
	acc = 0
	seen := make(map[int]bool)

	for true {
		_, found := seen[pc]
		if found {
			reason = "loop"
			break
		}

		seen[pc] = true

		if pc >= len(program) {
			reason = "terminate"
			break
		}

		i := program[pc]

		switch i.Opcode {
		case "acc":
			acc += i.Operand
		case "jmp":
			pc += (i.Operand - 1)
		case "nop":
		}

		pc++
	}

	return
}

func part1(program []Instruction) {
	acc, _ := run(program)
	fmt.Println("part 1 =", acc)
}

func part2(program []Instruction) {
	var acc int
	var reason string

	for index, instruction := range program {
		alteredProgram := make([]Instruction, len(program))
		copy(alteredProgram, program)

		if instruction.Opcode == "jmp" {
			instruction.Opcode = "nop"
			alteredProgram[index] = instruction
		} else if instruction.Opcode == "nop" {
			instruction.Opcode = "jmp"
			alteredProgram[index] = instruction
		} else {
			continue
		}

		acc, reason = run(alteredProgram)
		if reason == "terminate" {
			break
		}
	}

	fmt.Println("part 2 =", acc)
}

func main() {
	program := readData()

	part1(program)
	part2(program)
}

// Local Variables:
// compile-command: "go build"
// End:
