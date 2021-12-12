package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Link struct {
	Ends [2]string
}

type CaveMap map[string][]string

func readLinks() (links []Link) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		caves := strings.Split(line, "-")
		links = append(links, Link{Ends: [2]string{caves[0], caves[1]}})
	}

	return links
}

func buildMap(links []Link) (caveMap CaveMap) {
	caveMap = make(map[string][]string)

	for _, link := range links {
		c0 := link.Ends[0]
		c1 := link.Ends[1]

		ls, ok := caveMap[c0]
		if !ok {
			ls = make([]string, 0)
		}
		ls = append(ls, c1)
		caveMap[c0] = ls

		ls, ok = caveMap[c1]
		if !ok {
			ls = make([]string, 0)
		}
		ls = append(ls, c0)
		caveMap[c1] = ls
	}

	return
}

func isSmall(c string) bool {
	return c > "Z"
}

func allPaths(s string, cm CaveMap, visited map[string]bool, path []string, depth int) (count int) {
	adjacencies := cm[s]

	if s == "end" {
		fmt.Println(path)
		return 1
	}

	if isSmall(s) {
		visited[s] = true
	}

	for _, c := range adjacencies {
		_, seen := visited[c]
		if !seen {
			path = append(path, c)
			count += allPaths(c, cm, visited, path, depth+1)
		}
	}

	delete(visited, s)

	return
}

func part1(caveMap CaveMap) {
	visited := make(map[string]bool)
	fmt.Println(allPaths("start", caveMap, visited, []string{}, 0))
}

func part2(caveMap CaveMap) {
}

func main() {
	links := readLinks()
	caveMap := buildMap(links)
	part1(caveMap)
	part2(caveMap)
}

// Local Variables:
// compile-command: "go build"
// End:
