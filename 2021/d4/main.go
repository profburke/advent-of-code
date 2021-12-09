package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func copySlice(s []Board) (r []Board) {
	r = make([]Board, len(s))
	copy(r, s)

	return
}

type Entry struct {
	Number int
	Called bool
}

type Row []Entry
type Board struct {
	Rows  []Row
	Bingo bool
}

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
	rows := make([]Row, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			board.Rows = rows
			boards = append(boards, board)
			board = Board{}
			rows = make([]Row, 0)
			continue
		}

		row := Row{}
		parts := strings.Fields(line)
		for _, part := range parts {
			n, _ := strconv.Atoi(part)
			e := Entry{Number: n}
			row = append(row, e)
		}
		rows = append(rows, row)
	}
	board.Rows = rows
	boards = append(boards, board)

	return
}

func mark(n int, boards *[]Board) {
	for _, board := range *boards {
		for _, row := range board.Rows {
			for i, entry := range row {
				if n == entry.Number {
					row[i] = Entry{Number: n, Called: true}
				}
			}
		}
	}
}

func scoreBoard(board Board) (total int) {
	for _, row := range board.Rows {
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
	for _, row := range board.Rows {
		if !row[i].Called {
			return false
		}
	}

	return true
}

func boardWins(board Board) (total int, won bool) {
	for i, row := range board.Rows {
		if checkRow(row) {
			won = true
			total = scoreBoard(board)
		}
		if checkColumn(board, i) {
			won = true
			total = scoreBoard(board)
		}
	}

	return
}

func winner(boards []Board) (total, boardNum int, won bool) {
	for b, board := range boards {
		// this is pointless for part 1
		// and not used in part 2, so ... ???
		if board.Bingo {
			continue
		}

		total, won = boardWins(board)
		if won {
			boardNum = b
			return
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

type BoardRecord struct {
	B     Board
	Score int
}

func part2(numbers []int, boards []Board) {
	boardsInPlay := copySlice(boards)

	finishedBoards := make([]BoardRecord, 0)

	for _, n := range numbers {
		mark(n, &boardsInPlay)

		continueBoards := make([]Board, 0)
		for _, board := range boardsInPlay {

			total, won := boardWins(board)
			if won {
				finishedBoards = append(finishedBoards,
					BoardRecord{B: board, Score: n * total})
			} else {
				continueBoards = append(continueBoards, board)
			}
		}

		boardsInPlay = continueBoards
	}

	fmt.Println(finishedBoards[len(finishedBoards)-1].Score)
}

func main() {
	numbers, boards := readData()
	part1(numbers, copySlice(boards))
	part2(numbers, boards)
}

// Local Variables:
// compile-command: "go build"
// End:
