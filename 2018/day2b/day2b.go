package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
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

func CommonLetters(id1 string, id2 string) (result string, err error) {
	if len(id1) != len(id2) {
		return "", errors.New("id1 and id2 have different lengths")
	}

	result = ""

	for index, char := range id1 {
		otherChar := rune(id2[index])

		if char == otherChar {
			result += string(char)
		}
	}
	
	return result, nil
}

func main() {
	ids, err := ReadLines(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading ids: ", err)
		os.Exit(1)
	}

	for index, id1 := range ids {
		for _, id2 := range ids[index + 1:] {
			common, err := CommonLetters(id1, id2)

			if err != nil {
				fmt.Fprintln(os.Stderr, "problem comparing ids: ", err)
				os.Exit(1)
			}
			
			if len(common) == (len(id1) - 1) {
				fmt.Println(common)
				os.Exit(0)
			}
		}
	}
	fmt.Println("could not find solution")
}
