package main

import (
	"bufio"
	"fmt"
	"log"
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

const UintSize = 64 << (^uint(0) >> 64 & 1)
const MaxInt = 1<<(UintSize-1) - 1
const MinInt = -MaxInt // this isn't quite right, but it's close enough

func calc(vals []int, t int) (value int) {
	switch t {
	case 0:
		for _, v := range vals {
			value += v
		}
	case 1:
		value = vals[0]
		for _, v := range vals[1:] {
			value *= v
		}
	case 2:
		value = MaxInt
		for _, v := range vals {
			if v < value {
				value = v
			}
		}
	case 3:
		value = MinInt
		for _, v := range vals {
			if v > value {
				value = v
			}
		}
	case 5:
		if vals[0] > vals[1] {
			value = 1
		} else {
			value = 0
		}
	case 6:
		if vals[0] < vals[1] {
			value = 1
		} else {
			value = 0
		}
	case 7:
		if vals[0] == vals[1] {
			value = 1
		} else {
			value = 0
		}
	default:
		log.Fatal("oops")
	}

	return
}

func parseOperator(bits string, t int) (value int, nbits string) {
	ltb := string(bits[0])
	bits = bits[1:]

	if ltb == "0" {
		tl64, _ := strconv.ParseInt(bits[:15], 2, 64)
		tl := int(tl64)
		bits = bits[15:]
		value = parseBitstring(bits[:tl], t)
		bits = bits[tl:]
	} else {
		nsub64, _ := strconv.ParseInt(bits[:11], 2, 64)
		nsub := int(nsub64)
		bits = bits[11:]
		var v int
		vals := make([]int, 0)
		for j := 0; j < nsub; j++ {
			v, bits = parsePacket(bits)
			vals = append(vals, v)
		}

		value = calc(vals, t)
	}

	nbits = bits
	return
}

func parsePacket(bits string) (value int, nbits string) {
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
		value, bits = parseLiteral(bits)
	default:
		value, bits = parseOperator(bits, t)
	}

	nbits = bits
	return
}

func parseBitstring(bits string, t int) (value int) {
	vals := make([]int, 0)

	for len(bits) > 7 { // must be at least the 6 header bits remaining...
		var v int
		v, bits = parsePacket(bits)
		vals = append(vals, v)
	}

	value = calc(vals, t)
	return
}

func parts1And2(input string) {
	bits := tobits(input)
	fmt.Println("value", parseBitstring(bits, 0))
	fmt.Println(versionSum)
}

func main() {
	input := readInput()
	parts1And2(input)
}

// Local Variables:
// compile-command: "go build"
// End:
