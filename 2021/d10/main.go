package main

import (
	"bufio"
	"fmt"
	"os"
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

func part1(lines [][]string) {
	score := 0

	for _, line := range lines {
		score += scoreLine(line)
	}

	fmt.Println(score)
}

func part2(lines [][]string) {
}

func main() {
	lines := readLines()
	part1(lines)
	part2(lines)
}

// Local Variables:
// compile-command: "go build"
// End:
