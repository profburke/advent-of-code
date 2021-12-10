package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func readLines() (lines [][]string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "")
		lines = append(lines, tokens)
	}

	return
}

func isOpener(token string) bool {
	return token == "(" || token == "[" || token == "{" || token == "<"
}

func match(a, b string) bool {
	switch a {
	case "(":
		return b == ")"
	case "[":
		return b == "]"
	case "{":
		return b == "}"
	case "<":
		return b == ">"
	default:
		return false
	}
}

func scoreFor(token string) int {
	switch token {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		return 0
	}
}

func scoreLine(line []string) int {
	score := 0
	stack := make([]string, 0)

	for _, token := range line {
		if isOpener(token) {
			stack = append(stack, token)
		} else { // is closer
			if len(stack) == 0 {
				break
			}

			opener := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !match(opener, token) {
				score += scoreFor(token)
				break
			}
		}
	}

	return score
}

func part1(lines [][]string) (incompleteLines [][]string) {
	score := 0

	for _, line := range lines {
		lineScore := scoreLine(line)
		score += lineScore

		if lineScore == 0 {
			incompleteLines = append(incompleteLines, line)
		}
	}

	fmt.Println(score)

	return
}

func makeCompletion(stack []string) (completer string) {
	for i := len(stack) - 1; i >= 0; i-- {
		token := stack[i]
		switch token {
		case "(":
			completer += ")"
		case "[":
			completer += "]"
		case "{":
			completer += "}"
		case "<":
			completer += ">"
		}
	}
	return
}

func scoreCompleter(completer string) (score int) {
	for _, token := range completer {
		score *= 5

		switch string(token) {
		case ")":
			score += 1
		case "]":
			score += 2
		case "}":
			score += 3
		case ">":
			score += 4
		}
	}

	return
}

func part2(lines [][]string) {
	completers := make([]string, 0)
	stack := make([]string, 0)

	for _, line := range lines {
		for _, token := range line {
			if isOpener(token) {
				stack = append(stack, token)
			} else { // is closer
				if len(stack) == 0 {
					log.Fatal(errors.New("tried to pop from empty stack"))
				}

				opener := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if !match(opener, token) {
					log.Fatal(errors.New("corrupted line"))
				}
			}
		}

		completionString := makeCompletion(stack)
		completers = append(completers, completionString)
		stack = make([]string, 0)
	}

	scores := make([]int, 0)
	for _, completer := range completers {
		cScore := scoreCompleter(completer)
		scores = append(scores, cScore)
	}

	sort.Ints(scores)
	n := len(scores) / 2
	fmt.Println(scores[n])
}

func main() {
	lines := readLines()
	incompleteLines := part1(lines)
	part2(incompleteLines)
}

// Local Variables:
// compile-command: "go build"
// End:
