package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Opcode    string
	Variable1 string
	Variable2 string // second param may be a variable
	Value     int    // or may be a constant
	UseV2     bool
}

func readInstructions() (instructions []Instruction) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var i Instruction
		if fields[0] == "inp" {
			i = Instruction{fields[0], fields[1], "", 0, false}
		} else {
			if v, err := strconv.Atoi(fields[2]); err == nil {
				i = Instruction{fields[0], fields[1], "", v, false}
			} else {
				i = Instruction{fields[0], fields[1], fields[2], 0, true}
			}

		}
		instructions = append(instructions, i)
	}

	return
}

const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func run(inputs []int, instructions []Instruction) (state map[string]int) {
	state = make(map[string]int)
	state["w"] = 0
	state["x"] = 0
	state["y"] = 0
	state["z"] = 0
	idx := 0

	for _, i := range instructions {
		var op1, op2 int
		op1 = state[i.Variable1]
		if i.UseV2 {
			op2 = state[i.Variable2]
		} else {
			op2 = i.Value
		}

		switch i.Opcode {
		case "inp":
			state[i.Variable1] = inputs[idx]
			idx++
		case "add":
			state[i.Variable1] = op1 + op2
		case "mul":
			state[i.Variable1] = op1 * op2
		case "div":
			state[i.Variable1] = op1 / op2
		case "mod":
			state[i.Variable1] = op1 % op2
		case "eql":
			if op1 == op2 {
				state[i.Variable1] = 1
			} else {
				state[i.Variable1] = 0
			}
		}
	}

	return
}

func makeInputs(code int) []int {
	inputs := make([]int, 0)

	for i := 14; i > 0; i-- {
		inputs = append(inputs, digit(code, i))
	}

	return inputs
}

func isValid(code int, instructions []Instruction) bool {
	inputs := makeInputs(code)
	state := run(inputs, instructions)
	fmt.Println(code, state)
	return state["z"] == 0
}

func part1(instructions []Instruction) {
	//low := 11111111111111
	low := 88888888888888
	// high := 99999999999999

	for {
		r := isValid(low, instructions)
		if r == true {
			fmt.Println(low, true)
			break
		}
		low++
	}

	// var mid int
	// for low <= high {
	// 	mid = (low + high) / 2

	// 	if isValid(mid, instructions) {
	// 		low = mid + 1
	// 	} else {
	// 		high = mid - 1
	// 	}
	// }

	// fmt.Println(mid) // correct?
}

func part2(instructions []Instruction) {
}

func main() {
	instructions := readInstructions()
	part1(instructions)
	part2(instructions)
}

// Local Variables:
// compile-command: "go build"
// End:
