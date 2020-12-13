package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func ReadPolymer(r io.Reader) (s string, err error) {
	var b []byte
	b, err = ioutil.ReadAll(r)
	s = string(b)
	// remove linefeed at end
	s = string(s[:len(s) - 1])

	return s, err
}

func reacts(a string, b string) bool {
	if a != b && strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}

func zap(str string, l string) (result string) {
	zapper := func (r rune) rune {
		if strings.ToLower(string(r)) == strings.ToLower(l) {
			return -1
		} else {
			return r
		}
	}
	result = strings.Map(zapper, str)
	return result
}

func reduce(polymer string) (reduction string) {
	for {
		deletions := false
		prev := 0
		current := 1
		
		for current < len(polymer) {
			leftChar := string(polymer[prev])
			rightChar := string(polymer[current])
			if reacts(leftChar, rightChar) {
				if prev == 0 { // two reacting chars at beginning of string
					polymer = polymer[current + 1:]
				} else if current == len(polymer) - 1 { // at end
					polymer = polymer[:prev]
				} else { // in middle
					polymer = polymer[:prev] + polymer[current + 1:]
				}
				deletions = true
			} else {
				prev++
				current++
			}
		}
		if !deletions {
			break
		}
	}
	reduction = polymer
	return reduction
}

func main() {
	polymer, err := ReadPolymer(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading polymer: ", err)
		os.Exit(1)
	}

	lengths := make(map[string]int)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, rune := range alphabet {
		letter := string(rune)
		newPoly := zap(polymer, letter)
		reduction := reduce(newPoly)
		lengths[letter] = len(reduction)
	}

	minl := ""
	minimum := math.MaxInt64
	for l, v := range lengths {
		if v < minimum {
			minimum = v
			minl = l
		}
	}

	fmt.Println("min is ", minimum, " where l is ", minl)
}
