// -*- compile-command: "go build"; -*-
package main

import (
	"fmt"
	"os"
	"strconv"
)

func printState(scores []int, a, b int) {
	for i, v := range scores {
		opener, closer := "", ""
		if i == a {
			opener, closer = "(", ")"
		} else if i == b {
			opener, closer = "[", "]"
		}
			
		fmt.Print(fmt.Sprintf("%s%d%s ", opener, v, closer))
	}
	fmt.Println()
}

// rely on fact that (given our starting conditions) sum is always a 1 or 2 digit number
func bake(a, b int) (results []int) {
	// ugh! this special case was not covered in the samples
	// and 実は I didn't think of it; had to skim the reddit discussions
	if a == 0 && b == 0 {
		results = append(results, 0)
		return results
	}
	
	sum := a + b
	for sum > 0 {
		digit := sum % 10
		sum = sum / 10
		results = append(results, digit)
	}
	for left, right := 0, len(results)-1; left < right; left, right = left+1, right-1 {
		results[left], results[right] = results[right], results[left]
	}
	
	return results
}

func min(a, b int) int { if a < b { return a } else { return b } }

func main() {
	var targetNumber int
	if len(os.Args) > 1 {
		targetNumber, _ = strconv.Atoi(os.Args[1])
	} else {
		targetNumber = 147061
	}
	scores := []int{3, 7}
	currentA := 0
	currentB := 1

	for len(scores) < targetNumber + 10 {
		// printState(scores, currentA, currentB)
		for _, digit := range bake(scores[currentA], scores[currentB]) {
			scores = append(scores, digit)
		}
		
		currentA = (currentA + (1 + scores[currentA])) % len(scores)
		currentB = (currentB + (1 + scores[currentB])) % len(scores)
	}
	// printState(scores, currentA, currentB)
	lower := targetNumber
	upper := min(len(scores), targetNumber + 10)

	targetPattern := scores[lower:upper]
	fmt.Println("target pattern: ", targetPattern)
}
