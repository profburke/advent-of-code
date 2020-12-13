package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Vertex struct {
	X int
	Y int
}

type Claim struct {
	Id int
	X int
	Y int
	W int
	H int
}

var claimRE *regexp.Regexp
var cells map[Vertex]int

func ReadClaims(r io.Reader) (claims []Claim, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		matches := claimRE.FindAllStringSubmatch(line, -1)
		id, err := strconv.Atoi(string(matches[0][1]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error converting input: ", err)
			break
		}
		x, err := strconv.Atoi(string(matches[0][2]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error converting input: ", err)
			break
		}
		y, err := strconv.Atoi(string(matches[0][3]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error converting input: ", err)
			break
		}
		w, err := strconv.Atoi(string(matches[0][4]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error converting input: ", err)
			break
		}
		h, err := strconv.Atoi(string(matches[0][5]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error converting input: ", err)
			break
		}
		fmt.Println(id, x, y, w, h)
		claim := Claim{id, x, y, w, h}
		claims = append(claims, claim)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return claims, nil
}

func processClaim(claim Claim) {
	for x := claim.X + 1; x <= claim.X + claim.W; x++ {
		for y := claim.Y + 1; y <= claim.Y + claim.H; y++ {
			vertex := Vertex{x, y}
			cells[vertex] += 1
		}
	}
}

func main() {
	claimRE = regexp.MustCompile(`^#(\d+)\s+@\s+(\d+),(\d+):\s+(\d+)+x(\d+)$`)
	claims, err := ReadClaims(os.Stdin)
	cells = make(map[Vertex]int)
	
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading claims: ", err)
		os.Exit(1)
	}

	fmt.Println("number of claims: ", len(claims))
	for _, claim := range claims {
		processClaim(claim)
	}

	count := 0
	for _, v := range cells {
		if v > 1 {
			count++
		}
	}
	fmt.Println("There are ", count, "square inches of overlap.")
}
