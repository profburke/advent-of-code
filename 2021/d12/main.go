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

func readLinks() (links []Link, smalls []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		caves := strings.Split(line, "-")
		name := caves[0]
		links = append(links, Link{Ends: [2]string{name, caves[1]}})
		if isSmall(name) && name != "start" && name != "end" {
			smalls = append(smalls, name)
		}
	}

	return
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

var visited2 map[string]bool
var revisitable string
var visitCount int

func canVisit(c string) bool {
	if c == "start" || c == "end" {
		return false
	}

	if isSmall(c) {
		_, seen := visited2[c]

		if !seen {
			if c == revisitable {
				visitCount++
			}
			return true
		} else {
			if c == revisitable && visitCount < 2 {
				visitCount++
				return true
			} else {
				return false
			}
		}
	}

	return true
}

func allPaths2(s string, cm CaveMap, path []string, depth int) (count int) {
	adjacencies := cm[s]

	if s == "end" {
		fmt.Println(path)
		return 1
	}

	for _, c := range adjacencies {
		if canVisit(c) {
			path = append(path, c)
			count += allPaths2(c, cm, path, depth+1)
		}
	}

	// delete(visited, s)

	return
}

func part2(caveMap CaveMap, smalls []string) {
	total := 0

	for _, small := range smalls {
		revisitable = small
		visitCount = 0
		visited2 = make(map[string]bool)
		total += allPaths2("start", caveMap, []string{}, 0)
	}

	fmt.Println(total)
}

func main() {
	links, smalls := readLinks()
	caveMap := buildMap(links)
	part1(caveMap)
	part2(caveMap, smalls)
}

// Local Variables:
// compile-command: "go build"
// End:
