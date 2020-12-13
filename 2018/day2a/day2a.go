package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadLines(r io.Reader) (ids []string, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return ids, nil
}

func main() {
	ids, err := ReadLines(os.Stdin)
	numberOfDupes, numberOfTriples := 0, 0
	
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading ids: ", err)
		os.Exit(1)
	}

	for _, id := range ids {
		hasDupe, hasTriple := false, false
		letters := strings.Split(id, "")
		counts := make(map[string]int)

		for _, letter := range letters {
			counts[letter] += 1
		}

		for _, count := range counts {
			if count == 2 {
				hasDupe = true
			} else if count == 3 {
				hasTriple = true
			}
		}

		if hasDupe {
			numberOfDupes++
		}
		if hasTriple {
			numberOfTriples++
		}
	}
	
	fmt.Println("Product of dupes and triples = ", numberOfDupes * numberOfTriples)
}
