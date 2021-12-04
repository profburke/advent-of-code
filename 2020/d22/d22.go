package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deck []int

func readData() (decks []Deck) {
	scanner := bufio.NewScanner(os.Stdin)
	decks = make([]Deck, 0)
	deck := make(Deck, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Player") {
			continue
		}

		if line == "" {
			decks = append(decks, deck)
			deck = make(Deck, 0)
			continue
		}

		card, _ := strconv.Atoi(line)
		deck = append(deck, card)
	}
	decks = append(decks, deck)

	return
}

func gameOver(decks []Deck) (winner int) {
	for i, deck := range decks {
		if len(deck) == 0 {
			return 1 - i
		}
	}

	return -1
}

func score(deck Deck) (result int) {
	multiplier := len(deck)

	for _, card := range deck {
		result += (card * multiplier)
		multiplier--
	}

	return
}

func part1(decks []Deck) {
	var winner int

	for true {
		if winner = gameOver(decks); winner > -1 {
			break
		}

		aCard := decks[0][0]
		decks[0] = decks[0][1:]

		bCard := decks[1][0]
		decks[1] = decks[1][1:]

		if aCard > bCard {
			decks[0] = append(decks[0], aCard, bCard)
		} else {
			decks[1] = append(decks[1], bCard, aCard)
		}
	}

	score := score(decks[winner])
	fmt.Println("Part 1 =", score)
}

func keyFrom(decks []Deck) (result string) {
	ds := make([]string, 0)

	for _, deck := range decks {
		s := make([]string, 0)
		for _, card := range deck {
			s = append(s, fmt.Sprintf("%d", card))
		}
		ds = append(ds, strings.Join(s, "-"))
	}
	result = strings.Join(ds, "::")

	return
}

func recursivePlay(decks []Deck) (winner int) {
	ko := make(map[string]bool)

	for true {
		key := keyFrom(decks)

		if _, found := ko[keyFrom(decks)]; found {
			winner = 0
			break
		}

		ko[key] = true

		if winner = gameOver(decks); winner > -1 {
			break
		}

		aCard := decks[0][0]
		decks[0] = decks[0][1:]
		aLen := len(decks[0])

		bCard := decks[1][0]
		decks[1] = decks[1][1:]
		bLen := len(decks[1])

		roundWinner := 0
		if aLen < aCard || bLen < bCard {
			// if one player doesn't have enough, winner is high card

			if aCard > bCard {
				roundWinner = 0
			} else {
				roundWinner = 1
			}
		} else {
			// otherwise winner is winner of subgame

			newDecks := make([]Deck, 0)
			var nd0, nd1 Deck

			for _, v := range decks[0][0:aCard] {
				nd0 = append(nd0, v)
			}
			for _, v := range decks[1][0:bCard] {
				nd1 = append(nd1, v)
			}
			newDecks = append(newDecks, nd0, nd1)

			roundWinner = recursivePlay(newDecks)
		}

		// add cards to winner's deck

		if roundWinner == 0 {
			decks[0] = append(decks[0], aCard, bCard)
		} else {
			decks[1] = append(decks[1], bCard, aCard)
		}
	}

	return
}

func part2(decks []Deck) {
	winner := recursivePlay(decks)
	score := score(decks[winner])

	fmt.Println("Part 2 =", score)
}

func main() {
	decks := readData()

	decks2 := make([]Deck, 0)
	var deck20, deck21 Deck

	for _, v := range decks[0] {
		deck20 = append(deck20, v)
	}
	for _, v := range decks[1] {
		deck21 = append(deck21, v)
	}
	decks2 = append(decks2, deck20, deck21)

	part1(decks)
	part2(decks2)
}

// Local Variables:
// compile-command: "go build"
// End:
