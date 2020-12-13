package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadPolymer(r io.Reader) (s string, err error) {
	var b []byte
	b, err = ioutil.ReadAll(r)
	s = string(b)

	return s, err
}

func reacts(a string, b string) bool {
	if a != b && strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}

func main() {
	polymer, err := ReadPolymer(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading polymer: ", err)
		os.Exit(1)
	}

	// remove linefeed at end
	polymer = string(polymer[:len(polymer) - 1])

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
			fmt.Println(len(polymer))
			break
		}
	}
}
