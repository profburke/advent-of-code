package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Insertions map[string]string

func readData() (template string, i Insertions) {
	i = make(Insertions)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	template = strings.TrimSpace(scanner.Text())

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		i[key] = value
	}

	return
}

func step(s string, rules Insertions) (ns string) {
	for i := 0; i < len(s)-1; i++ {
		key := s[i : i+2]
		v, ok := rules[key]
		if !ok {
			fmt.Println("key", key)
			fmt.Println(rules)
			log.Fatal("unexpected pair")
		}
		ns += (string(key[0]) + v)
	}
	ns += string(s[len(s)-1:])

	return
}

type Pair struct {
	E string
	V int
}

func score(template string) (s int) {
	counts := make(map[string]int)

	for _, e := range template {
		counts[string(e)]++
	}

	ps := make([]Pair, 0)

	for k, v := range counts {
		ps = append(ps, Pair{E: k, V: v})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].V < ps[j].V
	})

	l := ps[0].V
	h := ps[len(ps)-1].V
	s = h - l

	return
}

func part1(template string, rules Insertions) {
	for i := 0; i < 10; i++ {
		template = step(template, rules)
	}

	fmt.Println(score(template))
}

func part2(template string, i Insertions) {
}

func main() {
	template, i := readData()
	part1(template, i)
	part2(template, i)
}

// Local Variables:
// compile-command: "go build"
// End:
