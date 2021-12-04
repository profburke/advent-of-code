package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Description struct {
	Penultimate int
	Ultimate    int
}

func readData() (numbers map[int]Description, last int) {
	numbers = make(map[int]Description)
	scanner := bufio.NewScanner(os.Stdin)
	turn := 1
	var number int

	scanner.Scan()
	for _, e := range strings.Split(scanner.Text(), ",") {
		number, _ = strconv.Atoi(e)
		numbers[number] = Description{Penultimate: turn, Ultimate: turn}
		turn++
	}
	last = number

	return
}

func playGame(numbers map[int]Description, last, cutoff int) (result int) {
	turn := len(numbers) + 1
	var say int

	for true {
		ld := numbers[last]
		say = ld.Ultimate - ld.Penultimate

		if turn == cutoff {
			break
		}

		sd, found := numbers[say]
		if found {
			numbers[say] = Description{Penultimate: sd.Ultimate, Ultimate: turn}
		} else {
			numbers[say] = Description{Penultimate: turn, Ultimate: turn}
		}

		last = say
		turn++
	}
	result = say

	return
}

func part1(numbers map[int]Description, last int) {
	result := playGame(numbers, last, 2020)
	fmt.Println("Part 1 =", result)
}

func part2(numbers map[int]Description, last int) {
	result := playGame(numbers, last, 30000000)
	fmt.Println("Part 2 =", result)
}

func main() {
	numbers, last := readData()
	numbers2 := make(map[int]Description)
	for key, value := range numbers {
		numbers2[key] = value
	}

	part1(numbers, last)
	part2(numbers2, last)
}
