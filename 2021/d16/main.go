package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var versionSum int

func readInput() (input string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	return
}

func adjust(c int32) int32 {
	if c <= '9' {
		return c - '0'
	} else {
		return c - 'A' + 10
	}
}

func tobits(input string) (bits string) {
	for _, c := range input {
		ac := adjust(c)
		cbits := fmt.Sprintf("%04b", ac)
		bits += cbits
	}

	return
}

func parseLiteral(bits string) (l int, s string) {
	contbit := "1"
	ls := ""

	for contbit == "1" {
		contbit = string(bits[0])
		ls += bits[1:5]
		bits = bits[5:]
	}

	l64, _ := strconv.ParseInt(ls, 2, 64)
	l = int(l64)
	s = bits

	return
}

func parseOperator(bits string) string {
	ltb := string(bits[0])
	bits = bits[1:]

	if ltb == "0" {
		tl64, _ := strconv.ParseInt(bits[:15], 2, 64)
		tl := int(tl64)
		bits = bits[15:]
		parseBitstring(bits[:tl])
		bits = bits[tl:]
	} else {
		nsub64, _ := strconv.ParseInt(bits[:11], 2, 64)
		nsub := int(nsub64)
		bits = bits[11:]
		for j := 0; j < nsub; j++ {
			bits = parsePacket(bits)
		}
	}

	return bits
}

func parsePacket(bits string) string {
	vbits := bits[0:3]
	tbits := bits[3:6]
	bits = bits[6:]

	i64, _ := strconv.ParseInt(vbits, 2, 64)
	v := int(i64)
	i64, _ = strconv.ParseInt(tbits, 2, 64)
	t := int(i64)

	versionSum += v

	switch t {
	case 4:
		// for some reason if I try
		// l, i := ....
		// it thinks I'm declaring a new i
		var l int
		l, bits = parseLiteral(bits)
		fmt.Printf("version: %d type: %d literal %d\n", v, t, l)
	default:
		bits = parseOperator(bits)
		fmt.Printf("version: %d type: %d operator\n", v, t)
	}

	return bits
}

func parseBitstring(bits string) {
	for len(bits) > 7 { // must be at least the 6 header bits remaining...
		bits = parsePacket(bits)
	}
}

func part1(input string) {
	bits := tobits(input)
	parseBitstring(bits)

	fmt.Println(versionSum)
}

func part2(input string) {
}

func main() {
	input := readInput()
	part1(input)
	part2(input)
}

// Local Variables:
// compile-command: "go build"
// End:
