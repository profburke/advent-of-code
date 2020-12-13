// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"container/ring"
	"errors"
	"fmt"
	"io"
	"os"
)

func readParmaeters(r io.Reader) (nPlayers int, nMarbles int, err error) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return -1, -1, errors.New("no input line")
	}

	text := scanner.Text()
	matches, err := fmt.Sscanf(text, "%d players; last marble is worth %d", &nPlayers, &nMarbles)
	if err != nil || matches != 2 {
	}
	
	return nPlayers, nMarbles, err
}

func addMarble(points int, r *ring.Ring) (s *ring.Ring) {
	if r == nil {
		s = ring.New(1)
		s.Value = points
		return s
	} else {
		t := ring.New(1)
		t.Value = points
		s = r.Link(t)
		s = s.Prev()
		return s
	}
}

func printRing(r *ring.Ring, c *ring.Ring) {
	start := r
	for start != nil {
		open := ""
		close := ""
		if start == c {
			open = "("
			close = ")"
		}
		fmt.Print(fmt.Sprintf("%s%d%s ", open, start.Value.(int), close))
		start = start.Next()
		if start == r { break }
	}
	fmt.Println()
}

func main() {
	nPlayers, nMarbles, err := readParmaeters(os.Stdin)
	if err != nil {
		fmt.Println("error reading parameters: ", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("Playing with %d players up to marble %d", nPlayers, nMarbles))

	currentPlayer := 0
	nextPlayer := func() (result int) {
		result = (currentPlayer + 1) % nPlayers
		return result
	}

	scores := make(map[int]int)
	
	currentMarble := 1
	var r *ring.Ring
	r = addMarble(0, r)
	// start := r

	for currentMarble <= nMarbles {
		// fmt.Print(fmt.Sprintf("[%d] ", currentPlayer))
		// printRing(start, r)
		
		if currentMarble % 23 == 0 {
			// score
			// fmt.Println("Removing marble: ", currentMarble)
			scores[currentPlayer] += currentMarble
			r = r.Move(-8)
			t := r.Unlink(1)
			scores[currentPlayer] += t.Value.(int)
			r = r.Move(1)
			// fmt.Println("Score (", currentPlayer, "): ", scores[currentPlayer])
		} else {
			// place marble
			// fmt.Println("Placing marble: ", currentMarble)
			r = r.Move(1)
			r = addMarble(currentMarble, r)
		}
		currentPlayer = nextPlayer()
		currentMarble += 1
	}
	// fmt.Print(fmt.Sprintf("[%d] ", currentPlayer))
	// printRing(start, r)

	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Println("Max score is ", maxScore)
}
