package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Number int
	Called bool
}

type Row []Entry
type Board []Row

func readData() (numbers []int, boards []Board) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ",")
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		numbers = append(numbers, n)
	}

	scanner.Scan() // skip blank line

	board := Board{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, board)
			board = Board{}
			continue
		}

		row := Row{}
		parts := strings.Fields(line)
		for _, part := range parts {
			n, _ := strconv.Atoi(part)
			e := Entry{Number: n}
			row = append(row, e)
		}
		board = append(board, row)
	}
	boards = append(boards, board)

	return
}

func mark(n int, boards *[]Board) {
	for _, board := range *boards {
		for _, row := range board {
			for i, entry := range row {
				if n == entry.Number {
					row[i] = Entry{Number: n, Called: true}
				}
			}
		}
	}
}

func scoreBoard(board Board) (total int) {
	for _, row := range board {
		for _, e := range row {
			if !e.Called {
				total += e.Number
			}
		}
	}

	return
}

func checkRow(row Row) bool {
	for _, e := range row {
		if !e.Called {
			return false
		}
	}

	return true
}

func checkColumn(board Board, i int) bool {
	for _, row := range board {
		if !row[i].Called {
			return false
		}
	}

	return true
}

func winner(boards []Board) (total, boardNum int, won bool) {
	for b, board := range boards {
		for i, row := range board {
			if checkRow(row) {
				won = true
				total = scoreBoard(board)
				boardNum = b
				return
			}

			if checkColumn(board, i) {
				won = true
				total = scoreBoard(board)
				boardNum = b
				return
			}
		}
	}

	return
}

func part1(numbers []int, boards []Board) {
	score := 0

	for _, n := range numbers {
		mark(n, &boards)
		total, _, won := winner(boards)
		if won {
			fmt.Println(n, total)
			score = n * total
			break
		}
	}

	fmt.Println(score)
}

func part2(numbers []int, boards []Board) {
}

func main() {
	numbers, boards := readData()
	part1(numbers, boards)
	part2(numbers, boards)
}

// Local Variables:
// compile-command: "go build"
// End:
