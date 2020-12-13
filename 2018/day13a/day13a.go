// -*- compile-command: "go build"; -*-
package main

import (
	"bufio"
	"errors"
	// "io"
        "fmt"
	"os"
	"sort"
)

func parseRow(track [][]TrackType, carts []Cart, row int, line string) (cells []TrackType, newCarts []Cart) {
	for i, r := range line {
		c := string(r)
		cells = append(cells, trackType(c))
		if isCart(c) {
			cart := newCart(row, i, c)
			newCarts = append(newCarts, cart)
		}
	}
	return cells, newCarts
}

// func readTrack(r io.Reader) (track [][]TrackType, carts []Cart, err error) {
func readTrack(fileName string) (track [][]TrackType, carts []Cart, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("error opening file: ", err))
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	carts = make([]Cart, 0)
	row := 0
	
	for scanner.Scan() {
		line := scanner.Text()
		cells, newCarts := parseRow(track, carts, row, line)
		track = append(track, cells)
		for _, c := range newCarts {
			carts = append(carts, c)
		}
		row++
	}

	err = scanner.Err()
	return track, carts, err
}

func printTrack(track [][]TrackType, carts []Cart) {
	for i := 0; i < len(track); i++ {
		row := track[i]
		for j := 0; j < len(row); j++ {
			var char string
			switch row[j] {
			case Intersection:
				char = "+"
			case Vertical:
				char = "|"
			case Horizontal:
				char = "-"
			case SlashCurve:
				char = "/"
			case BackslashCurve:
				char = "\\"
			case Space:
				char = " "
			}
			cart := cartAt(carts, i, j)
			if cart != "-" {
				char = cart
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func main() {
	fileName := os.Args[1]
	// track, carts, err := readTrack(os.Stdin)
	track, carts, err := readTrack(fileName)
	if err != nil {
		fmt.Println("error reading track: ", err)
		os.Exit(1)
	}

	generation := 0
	for {

		//printTrack(track, carts)
		fmt.Println("gen: ", generation)


		fmt.Println("\n\n\n")

		sort.Slice(carts, func(i, j int) bool {
			ipos := carts[i].Position
			jpos := carts[j].Position
			if ipos.X < jpos.X {
				return true
			} else if ipos.X == jpos.X && ipos.Y < jpos.Y {
				return true
			} else {
				return false
			}
		})

		newCarts := make([]Cart, 0)
		for _, cart := range carts {
			newCart, _ := cart.Move(track)
			newCarts = append(newCarts, newCart)
		}
		carts = newCarts
		crash, where := collisionOccurred(carts)
		if crash {
			fmt.Println(where)
			os.Exit(0)
		}
		generation++
	}
}
