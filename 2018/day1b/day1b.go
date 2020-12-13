package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadInts(r io.Reader) (numbers []int, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		n, err := strconv.Atoi(text)

		if err == nil {
			numbers = append(numbers, n)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return numbers, nil
}

var seenFreqs map[int]bool
func haveBeenAt(freq int) bool {
	_, seen := seenFreqs[freq]
	return seen
}

func main() {
	currentFreq := 0
	seenFreqs = make(map[int]bool)
	seenFreqs[currentFreq] = true

	freqShifts, err := ReadInts(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading ints: ", err)
		os.Exit(1)
	}

	for {
		for _, shift := range freqShifts {
			currentFreq += shift

			if haveBeenAt(currentFreq) {
				fmt.Println("First duplicated frequency is ", currentFreq)
				os.Exit(0)
			}

			seenFreqs[currentFreq] = true
		}
	}
}
