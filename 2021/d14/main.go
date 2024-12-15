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

	fmt.Println(ps)
	l := ps[0].V
	h := ps[len(ps)-1].V
	s = h - l

	return
}

func part1(template string, rules Insertions) {
	orig := template

	for i := 0; i < 20; i++ {
		template = step(template, rules)
	}

	fmt.Println(score(template))

	ol := len(orig)

	fmt.Println("orig", orig)
	for i := 1; i < len(template)-ol; i++ {
		s := template[i : i+ol]
		if s == orig {
			fmt.Println(i, s)
		}
	}
}

func part2(template string, rules Insertions) {
	for i := 0; i < 40; i++ {
		fmt.Println("step", i+1)
		template = step(template, rules)
	}

	fmt.Println(score(template))
}

// too tired to figure out closed form, but this
// should let me calculate the length of the string after k steps
// note: for my input the template is 20 chars long
func a(k int) int {
	if k == 0 {
		return 20
	}
	return 2*a(k-1) - 1
}

// a(40) = 20890720927745

func main() {
	template, i := readData()
	part1(template, i)
	// part2(template, i)
}

// Local Variables:
// compile-command: "go build"
// End:
